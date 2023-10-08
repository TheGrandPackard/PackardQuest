package storers

import (
	"encoding/json"
	"io"
	"os"
	"sync"

	"github.com/thegrandpackard/PackardQuest/models"
)

type playerFileStore struct {
	fileName string
	mutex    sync.RWMutex
}

func NewPlayerFileStore(fileName string) (PlayerStore, error) {
	// Test opening the file, or create it
	f, err := os.OpenFile(fileName, os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return &playerFileStore{fileName: fileName}, err
}

func (s *playerFileStore) readPlayersFile() (models.Players, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	playersFile, err := os.OpenFile(s.fileName, os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer playersFile.Close()

	playersBytes, err := io.ReadAll(playersFile)
	if err != nil {
		return nil, err
	}

	// Handle empty files
	if len(playersBytes) == 0 {
		return models.Players{}, nil
	}

	var players models.Players
	err = json.Unmarshal(playersBytes, &players)
	if err != nil {
		return nil, err
	}

	return players, err
}

func (s *playerFileStore) writePlayersFile(players models.Players) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	playersBytes, err := json.Marshal(players)
	if err != nil {
		return err
	}

	return os.WriteFile(s.fileName, playersBytes, 0)
}

func (s *playerFileStore) getNextPlayerID() (int, error) {
	maxId := 0

	players, err := s.readPlayersFile()
	if err != nil {
		return 0, err
	}

	for _, player := range players {
		if player.ID > maxId {
			maxId = player.ID
		}
	}

	return maxId + 1, nil
}

func (s *playerFileStore) CreatePlayer(player *models.Player) error {
	// TODO: This sucks. Use UUIDs instead
	nextId, err := s.getNextPlayerID()
	if err != nil {
		return err
	}
	player.ID = nextId

	players, err := s.readPlayersFile()
	if err != nil {
		return err
	}

	// Check if player exists by ID
	for _, p := range players {
		if player.ID == p.ID {
			return errPlayerExists
		}
	}

	players = append(players, player)
	return s.writePlayersFile(players)
}

func (s *playerFileStore) GetPlayers() (models.Players, error) {
	return s.readPlayersFile()
}

func (s *playerFileStore) GetPlayerByID(playerID int) (*models.Player, error) {
	players, err := s.readPlayersFile()
	if err != nil {
		return nil, err
	}

	for _, player := range players {
		if player.ID == playerID {
			return player, nil
		}
	}

	return nil, errPlayerNotExists
}

func (s *playerFileStore) GetPlayerByName(playerName string) (*models.Player, error) {
	players, err := s.readPlayersFile()
	if err != nil {
		return nil, err
	}

	for _, player := range players {
		if player.Name == playerName {
			return player, nil
		}
	}

	return nil, errPlayerNotExists
}

func (s *playerFileStore) GetPlayerByWandID(wandID int) (*models.Player, error) {
	players, err := s.readPlayersFile()
	if err != nil {
		return nil, err
	}

	for _, player := range players {
		if player.WandID == wandID {
			return player, nil
		}
	}

	return nil, errPlayerNotExists
}

func (s *playerFileStore) UpdatePlayer(player *models.Player) error {
	players, err := s.readPlayersFile()
	if err != nil {
		return err
	}

	// Check if player exists by ID and replace it
	for _, p := range players {
		if player.ID == p.ID {
			*p = *player
			return s.writePlayersFile(players)
		}
	}

	return errPlayerNotExists
}

func (s *playerFileStore) DeletePlayer(playerID int) error {
	players, err := s.readPlayersFile()
	if err != nil {
		return err
	}

	// Check if player exists by ID
	for i, p := range players {
		if playerID == p.ID {
			players = append(players[:i], players[i+1:]...)
			return s.writePlayersFile(players)
		}
	}

	return errPlayerNotExists
}
