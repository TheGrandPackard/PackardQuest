package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thegrandpackard/PackardQuest/models"
)

type playerResponse struct {
	Player *models.Player `json:"player"`
}

func (a *api) getPlayer(c *gin.Context) {
	id, err := getIntParam(c, "id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newApiError(err))
		return
	}

	resp, err := a.playerManager.GetPlayerByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, newApiError(err))
		return
	}
	// TODO: 404 if player not found

	c.JSON(http.StatusOK, playerResponse{Player: resp})
}

type registerPlayerRequest struct {
	PlayerName string `json:"playerName"`
	WandID     int    `json:"wandId"`
}

func (a *api) registerPlayer(c *gin.Context) {
	req := registerPlayerRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newApiError(err))
		return
	}

	resp, err := a.playerManager.CreatePlayer(req.PlayerName, req.WandID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, newApiError(err))
		return
	}

	c.JSON(http.StatusCreated, playerResponse{Player: resp})
}
