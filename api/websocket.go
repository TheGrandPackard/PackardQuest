package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/thegrandpackard/PackardQuest/models"
)

func (a *api) websocketUpgraderPlayer(c *gin.Context) {
	// get player id
	id, err := getIntParam(c, "id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newApiError(err))
		return
	}

	// upgrade connection
	conn, err := a.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	// add connection to client map, closing previous connection for player if set
	if _, ok := a.clients[id]; ok {
		if err := a.clients[id].Close(); err != nil {
			log.Println("error closing previous player connection:", err)
		}
		delete(a.clients, id)
	}
	a.clients[id] = conn
}

type playerUpdate struct {
	Type   string         `json:"type"`
	Player *models.Player `json:"player"`
}

func (a *api) OnPlayerUpdate(player *models.Player) {
	if conn, ok := a.clients[player.ID]; ok {
		playerUpdate := playerUpdate{
			Type:   "playerUpdate",
			Player: player,
		}

		playerUpdateBytes, err := json.Marshal(playerUpdate)
		if err != nil {
			log.Println("marshal player update:", err)
			return
		}

		if err := conn.WriteMessage(websocket.TextMessage, playerUpdateBytes); err != nil {
			log.Println("write player update:", err)
			return
		}
	}
}
