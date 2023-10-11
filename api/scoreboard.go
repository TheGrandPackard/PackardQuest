package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thegrandpackard/PackardQuest/models"
)

type scoreboardResponse struct {
	Houses  []*models.HouseScore  `json:"houses"`
	Players []*models.PlayerScore `json:"players"`
}

func (a *api) getScoreboard(c *gin.Context) {
	houses, players, err := a.scorebaordManager.GetScoreboards()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, newApiError(err))
		return
	}

	c.JSON(http.StatusOK, scoreboardResponse{
		Houses:  houses,
		Players: players,
	})
}
