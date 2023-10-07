package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
			log.Print(err)
		}
		delete(a.clients, id)
	}
	a.clients[id] = conn
}
