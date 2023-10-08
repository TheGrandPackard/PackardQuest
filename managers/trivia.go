package managers

import (
	"errors"

	"github.com/thegrandpackard/PackardQuest/models"
	"github.com/thegrandpackard/PackardQuest/storers"
)

type triviaQuestionManager struct {
	store storers.TriviaQuestionStore
}

func NewTriviaQuestionManager(triviaStore storers.TriviaQuestionStore) (TriviaQuestionManager, error) {
	return &triviaQuestionManager{
		store: triviaStore,
	}, nil
}

func (t *triviaQuestionManager) GetQuestionForPlayer(playerID int) (*models.TriviaQuestion, error) {
	// TODO
	return nil, errors.New("not implemented")
}

func (t *triviaQuestionManager) AnswerQuestion(playerID int, answer *models.PlayerAnswer) (bool, error) {
	// TODO
	return false, errors.New("not implemented")
}
