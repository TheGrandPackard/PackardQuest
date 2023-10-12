package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thegrandpackard/PackardQuest/models"
)

func (a *api) getPlayerByID(c *gin.Context) {
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

	c.JSON(http.StatusOK, models.PlayerResponse{Player: resp})
}

func (a *api) getPlayerByWandID(c *gin.Context) {
	wandID, err := getIntParam(c, "wandID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newApiError(err))
		return
	}

	resp, err := a.playerManager.GetPlayerByWandID(wandID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, newApiError(err))
		return
	}
	// TODO: 404 if player not found

	c.JSON(http.StatusOK, models.PlayerResponse{Player: resp})
}

func (a *api) registerPlayer(c *gin.Context) {
	req := models.RegisterPlayerRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newApiError(err))
		return
	}

	resp, err := a.playerManager.CreatePlayer(req.Name, req.WandID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, newApiError(err))
		return
	}

	c.JSON(http.StatusCreated, models.PlayerResponse{Player: resp})
}

func (a *api) updatePlayer(c *gin.Context) {
	id, err := getIntParam(c, "id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newApiError(err))
		return
	}

	req := models.UpdatePlayerRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newApiError(err))
		return
	}

	resp, err := a.playerManager.UpdatePlayer(id, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, newApiError(err))
		return
	}

	c.JSON(http.StatusCreated, models.PlayerResponse{Player: resp})
}
