package controller

import (
	"Pirates/api/request"
	"Pirates/api/response"
	"Pirates/game"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const ANSWER_TIME = 100

type Client struct {
	ShipId     string
	ToSlow     bool
	Connection *websocket.Conn
}

type ShipConnectionHandler struct {
	clients []*Client
	game    *game.Game
	mutex   sync.Mutex
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewShipConnectionHandler(game *game.Game) *ShipConnectionHandler {
	handler := new(ShipConnectionHandler)
	handler.game = game
	handler.clients = make([]*Client, 0, 0)

	game.RegisterUpdateListener(handler.Tick)

	return handler
}

func (s *ShipConnectionHandler) AddConnection(gc *gin.Context) {
	shipId := gc.Param("shipId")

	conn, err := upgrader.Upgrade(gc.Writer, gc.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := Client{shipId, false, conn}
	conn.SetCloseHandler(s.removeClientHandler(&client))

	s.mutex.Lock()
	s.clients = append(s.clients, &client)
	s.mutex.Unlock()
}

func (s *ShipConnectionHandler) removeClientHandler(client *Client) func(code int, text string) error {
	return func(code int, text string) error {
		s.mutex.Lock()
		s.closeAndRemoveClient(client)
		s.mutex.Unlock()
		return websocket.ErrCloseSent
	}
}

func (s *ShipConnectionHandler) closeAndRemoveClient(client *Client) {
	for i := 0; i < len(s.clients); i++ {
		if s.clients[i] == client {
			client.Connection.Close()
			s.removeConnection(i)
			return
		}
	}
}

func (s *ShipConnectionHandler) removeConnection(i int) {
	copy(s.clients[i:], s.clients[i+1:])     // Shift s.clients[i+1:] left one index.
	s.clients[len(s.clients)-1] = nil        // Erase last element (write zero value).
	s.clients = s.clients[:len(s.clients)-1] // Truncate slice.
}

func readActionFromClient(client *Client, actionChan chan request.Action) {
	go (func() {
		action := request.Action{}
		client.Connection.ReadJSON(&action)

		actionChan <- action
	})()
}

func (s *ShipConnectionHandler) handleClient(client *Client) {
	readChan := make(chan request.Action)
	info := s.getInfoForClient(client)
	if info == nil {
		s.mutex.Lock()
		s.closeAndRemoveClient(client)
		s.mutex.Unlock()
		return
	}

	// Write Status to Client
	err := client.Connection.WriteJSON(info)
	if err != nil {
		s.mutex.Lock()
		s.closeAndRemoveClient(client)
		s.mutex.Unlock()
	}

	// Read Action from Client
	readActionFromClient(client, readChan)
	select {
	case action := <-readChan:
		s.game.SetActionForShipId(client.ShipId, action)
		client.ToSlow = false
	case <-time.After(ANSWER_TIME * time.Millisecond):
		client.ToSlow = true
	}
}

func (s *ShipConnectionHandler) setActionForClient(client *Client, action request.Action) {
	s.game.SetActionForShipId(client.ShipId, action)
}

func (s *ShipConnectionHandler) getInfoForClient(client *Client) *response.Info {
	info := s.game.GetInfoForShipId(client.ShipId)
	if info == nil {
		return nil
	}

	if client.ToSlow {
		info.Error = "[WEBSOCKET-1] Response is to slow: max " + strconv.Itoa(ANSWER_TIME) + " milliseconds"
	}

	return info
}

func (s *ShipConnectionHandler) Tick() {
	s.mutex.Lock()
	for i := 0; i < len(s.clients); i++ {
		client := s.clients[i]
		go s.handleClient(client)
	}
	s.mutex.Unlock()
}
