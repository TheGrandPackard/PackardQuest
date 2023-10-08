package managers

import (
	"sort"

	"github.com/thegrandpackard/PackardQuest/interfaces"
	"github.com/thegrandpackard/PackardQuest/models"
)

type scoreboardManager struct {
	playerManager         interfaces.PlayerManager
	triviaQuestionManager interfaces.TriviaQuestionManager
}

func NewScoreboardManager(
	playerManager interfaces.PlayerManager,
	triviaQuestionManager interfaces.TriviaQuestionManager,
) (interfaces.ScoreboardManager, error) {
	return &scoreboardManager{
		playerManager:         playerManager,
		triviaQuestionManager: triviaQuestionManager,
	}, nil
}

func (s *scoreboardManager) GetScoreboards() ([]*models.HouseScore, []*models.PlayerScore, error) {
	// Get all players
	players, err := s.playerManager.GetPlayers()
	if err != nil {
		return nil, nil, err
	}

	// Get all trivia questions and build map with score
	triviaQuestions, err := s.triviaQuestionManager.GetQuestions()
	if err != nil {
		return nil, nil, err
	}
	triviaQuestionPoints := map[int]int{}
	for _, question := range triviaQuestions {
		triviaQuestionPoints[question.ID] = question.Points
	}

	// Get house and player scores
	houseScores := []*models.HouseScore{}
	playerScores := []*models.PlayerScore{}

	houseScoreMap := map[models.HogwartsHouse]int{}
	for _, house := range models.HogwartsHouses {
		houseScoreMap[house] = 0
	}
	for _, player := range players {
		score := player.GetScore(triviaQuestionPoints)
		playerScores = append(playerScores, &models.PlayerScore{
			Name:  player.Name,
			House: player.House,
			Score: score,
		})

		houseScoreMap[player.House] += score
	}

	// Flatten house score map
	for house, score := range houseScoreMap {
		houseScores = append(houseScores, &models.HouseScore{
			Name:  house,
			Score: score,
		})
	}

	// Sort house scores descending
	sort.Slice(houseScores, func(i, j int) bool {
		return houseScores[i].Score > houseScores[j].Score
	})

	// Player house scores descending
	sort.Slice(playerScores, func(i, j int) bool {
		return playerScores[i].Score > playerScores[j].Score
	})

	return houseScores, playerScores, nil
}
