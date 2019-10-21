package controller

import (
	"Pirates/api/request"
	"Pirates/events"
	"Pirates/game"
	"Pirates/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ApiController struct {
	handler *ShipConnectionHandler
	game    *game.Game
	secret  string
}

func NewApiController(router *gin.RouterGroup, game *game.Game) *ApiController {
	controller := ApiController{}
	controller.handler = NewShipConnectionHandler(game)
	controller.game = game
	controller.generateSecret()

	router.POST("registerPlayer", controller.RegisterPlayer)
	router.POST("buyShip", controller.BuyShip)

	router.GET("player/:player/:secret", controller.GetPlayer)
	router.GET("ships/:player/:secret", controller.GetShips)
	router.GET("shipControl/:shipId/:player/:secret", controller.ShipController)

	router.GET("status/"+controller.secret, controller.GetStatus)

	return &controller
}

func (a *ApiController) generateSecret() {
	a.secret = util.RandSeq(16)
	log.Println("Status Secret: ", a.secret)
}

func (a *ApiController) GetStatus(gc *gin.Context) {
	gc.JSON(http.StatusOK, gin.H{
		"ocean":   a.game.GetOcean(),
		"players": a.game.GetPlayers(),
		"events":  events.GetInstance().GetEvents(),
	})
}

func (a *ApiController) RegisterPlayer(gc *gin.Context) {
	var playerRequest request.PlayerRequest
	gc.BindJSON(&playerRequest)

	player := a.game.AddPlayer(playerRequest.Name, playerRequest.Secret)
	if player == nil {
		gc.JSON(http.StatusConflict, "[PLAYER-2] Could not register Player, Name already taken")
	} else {
		gc.JSON(http.StatusCreated, player)
	}
}

func (a *ApiController) GetPlayer(gc *gin.Context) {
	player := a.game.GetPlayerWithRequest(request.PlayerRequest{
		gc.Param("player"),
		gc.Param("secret"),
	})
	if player == nil {
		gc.JSON(http.StatusConflict, "[PLAYER-3] Player not found")
	} else {
		gc.JSON(http.StatusOK, player)
	}
}

func (a *ApiController) GetShips(gc *gin.Context) {
	player := a.game.GetPlayerWithRequest(request.PlayerRequest{
		gc.Param("player"),
		gc.Param("secret"),
	})

	if player == nil {
		gc.JSON(http.StatusConflict, "[PLAYER-3] Player not found")
		return
	}

	ships, serr := a.game.GetShipsForPlayer(player)
	if serr != nil {
		gc.JSON(http.StatusInternalServerError, "[SHIPS-1] Server error")
	} else {
		gc.JSON(http.StatusOK, ships)
	}
}

func (a *ApiController) BuyShip(gc *gin.Context) {
	var buyRequest request.Buy
	if err := gc.BindJSON(&buyRequest); err != nil {
		gc.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	if ship, err := a.game.NewShip(&buyRequest); err != nil {
		gc.JSON(http.StatusConflict, err.Error())
		return
	} else {
		gc.JSON(http.StatusCreated, ship)
	}
}

func (a *ApiController) ShipController(gc *gin.Context) {
	player := a.game.GetPlayerWithRequest(request.PlayerRequest{
		gc.Param("player"),
		gc.Param("secret"),
	})
	if player == nil {
		gc.JSON(http.StatusConflict, "[PLAYER-4] Player for ship not found")
		return
	}

	if a.game.PlayerHasShip(player, gc.Param("shipId")) {
		a.handler.AddConnection(gc)
	} else {
		gc.JSON(http.StatusInternalServerError, "[SHIPS-2] Not your ship!")
	}
}
