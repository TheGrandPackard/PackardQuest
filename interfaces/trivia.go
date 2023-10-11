package interfaces

import "github.com/thegrandpackard/PackardQuest/models"

type TriviaQuestionManager interface {
	GetQuestions() (models.TriviaQuestions, error)
	GetQuestionForPlayer(playerID int) (*models.TriviaQuestion, error)

	AnswerQuestion(playerID int, answer *models.PlayerAnswer) (bool, error)
}
