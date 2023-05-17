package storers

import (
	"errors"

	"github.com/thegrandpackard/PackardQuest/models"
)

type PlayerStore interface {
	CreatePlayer(*models.Player) error
	GetPlayers() (models.Players, error)
	GetPlayerByID(playerID int) (*models.Player, error)
	GetPlayerByName(playerName string) (*models.Player, error)
	GetPlayerByWandID(wandID int) (*models.Player, error)
	UpdatePlayer(*models.Player) error
	DeletePlayer(playerID int) error
}

var (
	errPlayerExists    error = errors.New("player already exists")
	errPlayerNotExists error = errors.New("player not exists")
)
