package storers

import (
	"encoding/json"
	"io"
	"os"
	"sync"

	"github.com/thegrandpackard/PackardQuest/models"
)

type triviaQuestionFileStore struct {
	fileName string
	mutex    sync.RWMutex
}

func NewTriviaQuestionFileStore(fileName string) (TriviaQuestionStore, error) {
	// Test opening the file, or create it
	f, err := os.OpenFile(fileName, os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return &triviaQuestionFileStore{fileName: fileName}, err
}

func (s *triviaQuestionFileStore) readTriviaQuestionsFile() (models.TriviaQuestions, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	triviaQuestionsFile, err := os.OpenFile(s.fileName, os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer triviaQuestionsFile.Close()

	triviaQuestionsBytes, err := io.ReadAll(triviaQuestionsFile)
	if err != nil {
		return nil, err
	}

	// Handle empty files
	if len(triviaQuestionsBytes) == 0 {
		return models.TriviaQuestions{}, nil
	}

	var triviaQuestions models.TriviaQuestions
	err = json.Unmarshal(triviaQuestionsBytes, &triviaQuestions)
	if err != nil {
		return nil, err
	}

	return triviaQuestions, err
}

func (s *triviaQuestionFileStore) GetTriviaQuestions() (models.TriviaQuestions, error) {
	return s.readTriviaQuestionsFile()
}

func (s *triviaQuestionFileStore) GetTriviaQuestionByID(triviaQuestionID int) (*models.TriviaQuestion, error) {
	triviaQuestions, err := s.readTriviaQuestionsFile()
	if err != nil {
		return nil, err
	}

	for _, triviaQuestion := range triviaQuestions {
		if triviaQuestion.ID == triviaQuestionID {
			return triviaQuestion, nil
		}
	}

	return nil, errTriviaQuestionNotExists
}
