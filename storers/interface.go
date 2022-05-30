package storers

import "github.com/thegrandpackard/PackardQuest/models"

type PlayerStore interface {
	GetPlayers() (models.Players, error)
	GetPlayerByName(playerName string) (*models.Player, error)
	GetPlayerByWandID(wandID int) (*models.Player, error)
}
