package api

import (
	"Pirates/api/controller"
	"Pirates/game"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

type HttpApi struct {
}

func (a *HttpApi) Init(game *game.Game) {
	r := gin.Default()

	r.Use(cors.Default())

	g := r.Group("/")
	controller.NewApiController(g, game)

	r.Run(":1337")
}
