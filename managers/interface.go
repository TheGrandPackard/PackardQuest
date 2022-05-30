package managers

import "github.com/thegrandpackard/PackardQuest/models"

type PlayerManager interface {
	GetPlayers() (models.Players, error)
	GetPlayerByName(playerName string) (*models.Player, error)
	GetPlayerByWandID(wandID int) (*models.Player, error)
}
