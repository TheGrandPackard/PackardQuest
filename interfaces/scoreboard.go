package interfaces

import "github.com/thegrandpackard/PackardQuest/models"

type ScoreboardManager interface {
	GetScoreboards() ([]*models.HouseScore, []*models.PlayerScore, error)
}
