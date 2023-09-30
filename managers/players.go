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

func (p *playerManager) GetPlayerByID(playerID int) (*models.Player, error) {
	return p.store.GetPlayerByID(playerID)
}

func (p *playerManager) GetPlayerByWandID(wandID int) (*models.Player, error) {
	return p.store.GetPlayerByWandID(wandID)
}

func (p *playerManager) getPlayerHouse() models.HogwartsHouse {
	// Randomly place a player into a house, but balance house distribution
	// Houses: Gryffindor, Slytherin, Ravenclaw, Hufflepuff

	return models.HogwartsHouseGryffindor
}

func (p *playerManager) CreatePlayer(playerName string, wandID int) (*models.Player, error) {
	house := p.getPlayerHouse()

	player := &models.Player{Name: playerName, WandID: wandID, House: house}
	if err := p.store.CreatePlayer(player); err != nil {
		return nil, err
	}

	return player, nil
}
