package storers

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/thegrandpackard/PackardQuest/models"
)

type fileStore struct {
	fileName string
}

func NewFileStore(fileName string) (PlayerStore, error) {
	// Test opening the file
	f, err := os.Open(fileName)
	defer f.Close()

	return &fileStore{fileName: fileName}, err
}

func readPlayersFile(fileName string) (models.Players, error) {
	playersFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	playersBytes, err := ioutil.ReadAll(playersFile)
	if err != nil {
		return nil, err
	}

	var players models.Players
	err = json.Unmarshal(playersBytes, &players)
	if err != nil {
		return nil, err
	}

	return players, err
}

func (s *fileStore) GetPlayers() (models.Players, error) {
	return readPlayersFile(s.fileName)
}

func (s *fileStore) GetPlayerByName(playerName string) (*models.Player, error) {
	players, err := readPlayersFile(s.fileName)
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
	players, err := readPlayersFile(s.fileName)
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
