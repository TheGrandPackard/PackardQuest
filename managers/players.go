package managers

import (
	"errors"
	"math"
	"math/rand"

	"github.com/thegrandpackard/PackardQuest/interfaces"
	"github.com/thegrandpackard/PackardQuest/models"
	"github.com/thegrandpackard/PackardQuest/storers"
)

type playerManager struct {
	store      storers.PlayerStore
	subscriber interfaces.PlayerManagerSubscriber
}

func NewPlayerManager(playerStore storers.PlayerStore) (interfaces.PlayerManager, error) {
	return &playerManager{
		store: playerStore,
	}, nil
}

func (p *playerManager) SetSubscriber(subscriber interfaces.PlayerManagerSubscriber) {
	p.subscriber = subscriber
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

func (p *playerManager) UpdatePlayer(id int, req models.UpdatePlayerRequest) (*models.Player, error) {
	// get player
	player, err := p.GetPlayerByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		if *req.Name == "" {
			return nil, errors.New("name cannot be empty")
		}

		player.Name = *req.Name
	}

	if req.House != nil {
		if *req.House == "" {
			return nil, errors.New("house cannot be empty")
		}
		if !models.IsValidHouse(*req.House) {
			return nil, errors.New("house must be valid")
		}

		player.House = *req.House
	}

	if req.WandID != nil {
		if *req.WandID == 0 {
			return nil, errors.New("wand id cannot be 0")
		}

		player.WandID = *req.WandID
	}

	if req.Progress != nil {
		player.Progress = *req.Progress
	}

	if req.TriviaAnswers != nil {
		if player.TriviaAnswers == nil {
			player.TriviaAnswers = map[int]bool{}
		}

		for id, answer := range req.TriviaAnswers {
			player.TriviaAnswers[id] = answer
		}
	}

	// update player
	err = p.store.UpdatePlayer(player)
	if err != nil {
		return nil, err
	}

	// trigger event to subscribers
	if p.subscriber != nil {
		go p.subscriber.OnPlayerUpdate(player)
	}

	return player, nil
}
