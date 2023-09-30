package managers

import (
	"math"
	"math/rand"

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

// Randomly place a player into a house, but balance house distribution
func getRandomlyDistributedPlayerHouse(players []*models.Player) models.HogwartsHouse {
	if len(players) == 0 {
		houseIdx := rand.Intn(len(models.HogwartsHouses))
		return models.HogwartsHouses[houseIdx]
	}

	// Count players for each house
	houseCountMap := map[models.HogwartsHouse]int{}
	for _, house := range models.HogwartsHouses {
		houseCountMap[house] = 0
	}
	for _, player := range players {
		houseCountMap[player.House]++
	}

	// Determine eligible houses
	houses := []models.HogwartsHouse{}
	minHouseCount := math.MaxInt
	for _, count := range houseCountMap {
		if count < minHouseCount {
			minHouseCount = count
		}
	}
	for house, count := range houseCountMap {
		if count == minHouseCount {
			houses = append(houses, house)
		}
	}

	// Pick random house from eligible houses
	houseIdx := rand.Intn(len(houses))
	return houses[houseIdx]
}

func (p *playerManager) CreatePlayer(playerName string, wandID int) (*models.Player, error) {
	// Get all players
	players, err := p.store.GetPlayers()
	if err != nil {
		return nil, err
	}

	// Get player house
	house := getRandomlyDistributedPlayerHouse(players)

	player := &models.Player{Name: playerName, WandID: wandID, House: house}
	if err := p.store.CreatePlayer(player); err != nil {
		return nil, err
	}

	return player, nil
}
