package models

type TriviaQuestion struct {
	ID            int            `json:"id"`
	Prompt        string         `json:"prompt"`
	Answers       []TriviaAnswer `json:"answers"`
	CorrectAnswer int            `json:"correctAnswer"`
	Points        int            `json:"points"`
	AudioFile     string         `json:"audioFile"`
}

type TriviaQuestions []*TriviaQuestion

type TriviaAnswer struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type PlayerAnswer struct {
	QuestionID int `json:"questionId"`
	Answer     int `json:"answer"`
}

type TriviaQuestionResponse struct {
	Question *TriviaQuestion `json:"question"`
}

type TriviaAnswerResponse struct {
	Correct bool `json:"correct"`
}
