package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thegrandpackard/PackardQuest/models"
)

type triviaQuestionResponse struct {
	Question *models.TriviaQuestion `json:"question"`
}

func (a *api) getPlayerTriviaQuestion(c *gin.Context) {
	id, err := getIntParam(c, "id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newApiError(err))
		return
	}

	resp, err := a.triviaQuestionManager.GetQuestionForPlayer(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, newApiError(err))
		return
	}

	c.JSON(http.StatusOK, triviaQuestionResponse{Question: resp})
}

type triviaAnswerResponse struct {
	Correct bool `json:"correct"`
}

func (a *api) answerTriviaQuestion(c *gin.Context) {
	id, err := getIntParam(c, "id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newApiError(err))
		return
	}

	req := &models.PlayerAnswer{}
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newApiError(err))
		return
	}

	resp, err := a.triviaQuestionManager.AnswerQuestion(id, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, newApiError(err))
		return
	}
	// TODO: 404 if player not found

	c.JSON(http.StatusOK, triviaAnswerResponse{Correct: resp})
}
