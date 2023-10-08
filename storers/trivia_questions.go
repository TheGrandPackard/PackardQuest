package storers

import (
	"errors"

	"github.com/thegrandpackard/PackardQuest/models"
)

type TriviaQuestionStore interface {
	GetTriviaQuestions() (models.TriviaQuestions, error)
	GetTriviaQuestionByID(id int) (*models.TriviaQuestion, error)
}

var (
	errTriviaQuestionNotExists error = errors.New("question not exists")
)
