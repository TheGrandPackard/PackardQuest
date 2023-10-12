package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thegrandpackard/PackardQuest/models"
)

func (a *api) getScoreboard(c *gin.Context) {
	houses, players, err := a.scorebaordManager.GetScoreboards()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, newApiError(err))
		return
	}

	c.JSON(http.StatusOK, models.ScoreboardResponse{
		Houses:  houses,
		Players: players,
	})
}
