package models

type Player struct {
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	WandID        int           `json:"wandID"`
	House         HogwartsHouse `json:"house"`
	Progress      Progress      `json:"progress"`
	TriviaAnswers map[int]bool  `json:"triviaAnswers"`
}

func (p *Player) GetScore(answerPoints map[int]int) int {
	playerScore := 0
	// iterate over all trivia questions the player has answered
	for questionID, correct := range p.TriviaAnswers {
		// if the answer was correct and the answerScores has a score for the question id, add it up
		if points, ok := answerPoints[questionID]; ok && correct {
			playerScore += points
		}
	}

	return playerScore
}

type Progress struct {
	SortingHat bool `json:"sortingHat"`
	Pensieve   bool `json:"pensieve"`
}

type Players []*Player

type PlayerResponse struct {
	Player *Player `json:"player"`
}
type RegisterPlayerRequest struct {
	Name   string `json:"name"`
	WandID int    `json:"wandId"`
}

type UpdatePlayerRequest struct {
	Name          *string        `json:"name"`
	House         *HogwartsHouse `json:"house"`
	WandID        *int           `json:"wandId"`
	Progress      *Progress      `json:"progress"`
	TriviaAnswers map[int]bool   `json:"triviaAnswers"`
}
