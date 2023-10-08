package managers

import (
	"errors"

	"github.com/thegrandpackard/PackardQuest/interfaces"
	"github.com/thegrandpackard/PackardQuest/models"
	"github.com/thegrandpackard/PackardQuest/storers"
)

type triviaQuestionManager struct {
	store         storers.TriviaQuestionStore
	playerManager interfaces.PlayerManager
}

func NewTriviaQuestionManager(
	triviaStore storers.TriviaQuestionStore,
	playerManager interfaces.PlayerManager,
) (interfaces.TriviaQuestionManager, error) {
	return &triviaQuestionManager{
		store:         triviaStore,
		playerManager: playerManager,
	}, nil
}

func (t *triviaQuestionManager) GetQuestions() (models.TriviaQuestions, error) {
	return t.store.GetTriviaQuestions()
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
	// get trivia question
	triviaQuestion, err := t.store.GetTriviaQuestionByID(answer.QuestionID)
	if err != nil {
		return false, err
	}

	// check answer
	correct := answer.Answer == triviaQuestion.CorrectAnswer

	// update player
	if _, err := t.playerManager.UpdatePlayer(playerID, models.UpdatePlayerRequest{
		TriviaAnswers: map[int]bool{answer.Answer: correct},
	}); err != nil {
		return false, err
	}

	return correct, nil
}
