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
	triviaQuestions, err := t.store.GetTriviaQuestions()
	if err != nil {
		return nil, err
	}

	if len(triviaQuestions) == 0 {
		return nil, errors.New("no trivia questions found")
	}

	return triviaQuestions[0], nil
}

func (t *triviaQuestionManager) AnswerQuestion(playerID int, answer *models.PlayerAnswer) (bool, error) {
	// TODO
	return false, errors.New("not implemented")
}
