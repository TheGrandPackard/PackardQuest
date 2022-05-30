package managers

import (
	"github.com/thegrandpackard/PackardQuest/models"
	"github.com/thegrandpackard/PackardQuest/storers"
)

type playerManager struct {
	store storers.PlayerStore
}

func NewPlayerManager(playerStore storers.PlayerStore) (PlayerManager, error) {
	return &playerManager{store: playerStore}, nil
}

func (p *playerManager) GetPlayers() (models.Players, error) {
	return p.store.GetPlayers()
}

func (p *playerManager) GetPlayerByName(playerName string) (*models.Player, error) {
	return p.store.GetPlayerByName(playerName)
}

func (p *playerManager) GetPlayerByWandID(wandID int) (*models.Player, error) {
	return p.store.GetPlayerByWandID(wandID)
}
