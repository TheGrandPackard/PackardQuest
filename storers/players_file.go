package storers

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	"github.com/thegrandpackard/PackardQuest/models"
)

type fileStore struct {
	fileName string
	mutex    sync.RWMutex
}

func NewFileStore(fileName string) (PlayerStore, error) {
	// Test opening the file, or create it
	f, err := os.OpenFile(fileName, os.O_CREATE, 0644)
	defer f.Close()

	return &fileStore{fileName: fileName}, err
}

func (s *fileStore) readPlayersFile() (models.Players, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	playersFile, err := os.OpenFile(s.fileName, os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer playersFile.Close()

	playersBytes, err := ioutil.ReadAll(playersFile)
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

func (s *fileStore) writePlayersFile(players models.Players) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	playersBytes, err := json.Marshal(players)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(s.fileName, playersBytes, 0)
}

func (s *fileStore) getNextPlayerID() (int, error) {
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

func (s *fileStore) CreatePlayer(player *models.Player) error {
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

func (s *fileStore) GetPlayers() (models.Players, error) {
	return s.readPlayersFile()
}

func (s *fileStore) GetPlayerByID(playerID int) (*models.Player, error) {
	players, err := s.readPlayersFile()
	if err != nil {
		return nil, err
	}

	for _, player := range players {
		if player.ID == playerID {
			return player, nil
		}
	}

	return nil, nil
}

func (s *fileStore) GetPlayerByName(playerName string) (*models.Player, error) {
	players, err := s.readPlayersFile()
	if err != nil {
		return nil, err
	}

	for _, player := range players {
		if player.Name == playerName {
			return player, nil
		}
	}

	return nil, nil
}

func (s *fileStore) GetPlayerByWandID(wandID int) (*models.Player, error) {
	players, err := s.readPlayersFile()
	if err != nil {
		return nil, err
	}

	for _, player := range players {
		if player.WandID == wandID {
			return player, nil
		}
	}

	return nil, nil
}

func (s *fileStore) UpdatePlayer(player *models.Player) error {
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

func (s *fileStore) DeletePlayer(playerID int) error {
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
