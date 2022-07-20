package storers

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thegrandpackard/PackardQuest/models"
)

func createTestPlayersFile(t *testing.T, filename string) PlayerStore {
	s, err := NewFileStore(filename)

	players := []*models.Player{
		{
			ID:     1,
			Name:   "Player One",
			WandID: 1001,
		},
		{
			ID:     2,
			Name:   "Player Two",
			WandID: 1002,
		},
	}

	for _, player := range players {
		err = s.CreatePlayer(player)
		assert.NoError(t, err)
	}

	return s
}

func TestCreatePlayer(t *testing.T) {
	s := createTestPlayersFile(t, "players.json")
	defer os.Remove("players.json")

	type test struct {
		player *models.Player
		err    error
	}

	tests := []test{
		{
			player: &models.Player{
				ID:     1,
				Name:   "Player 1",
				WandID: 1001,
			},
			err: errPlayerExists,
		},
		{
			player: &models.Player{
				ID:     3,
				Name:   "Player Three",
				WandID: 1003,
			},
			err: nil,
		},
	}

	for _, tc := range tests {
		err := s.CreatePlayer(tc.player)
		assert.Equal(t, tc.err, err)

		if err == nil {
			player, err := s.GetPlayerByID(tc.player.ID)
			assert.NoError(t, err)
			assert.Equal(t, tc.player, player)
		}
	}
}

func TestGetPlayers(t *testing.T) {
	s := createTestPlayersFile(t, "players.json")
	defer os.Remove("players.json")

	players, err := s.GetPlayers()
	assert.NoError(t, err)
	assert.NotEmpty(t, players)
	assert.Len(t, players, 2)
}

func TestGetPlayerByID(t *testing.T) {
	s := createTestPlayersFile(t, "players.json")
	defer os.Remove("players.json")

	type test struct {
		id     int
		exists bool
		err    error
	}

	tests := []test{
		{
			id:     1,
			exists: true,
			err:    nil,
		},
		{
			id:     0,
			exists: false,
			err:    nil,
		},
	}

	for _, tc := range tests {
		player, err := s.GetPlayerByID(tc.id)
		assert.Equal(t, tc.exists, player != nil)
		assert.Equal(t, tc.err, err)
	}
}

func TestGetPlayerByName(t *testing.T) {
	s := createTestPlayersFile(t, "players.json")
	defer os.Remove("players.json")

	type test struct {
		name   string
		exists bool
		err    error
	}

	tests := []test{
		{
			name:   "Player One",
			exists: true,
			err:    nil,
		},
		{
			name:   "Nonexistant User",
			exists: false,
			err:    nil,
		},
	}

	for _, tc := range tests {
		player, err := s.GetPlayerByName(tc.name)
		assert.Equal(t, tc.exists, player != nil)
		assert.Equal(t, tc.err, err)
	}
}

func TestGetPlayerByWandID(t *testing.T) {
	s := createTestPlayersFile(t, "players.json")
	defer os.Remove("players.json")

	type test struct {
		wandID int
		exists bool
		err    error
	}

	tests := []test{
		{
			wandID: 1001,
			exists: true,
			err:    nil,
		},
		{
			wandID: 0,
			exists: false,
			err:    nil,
		},
	}

	for _, tc := range tests {
		player, err := s.GetPlayerByWandID(tc.wandID)
		assert.Equal(t, tc.exists, player != nil)
		assert.Equal(t, tc.err, err)
	}
}

func TestUpdatePlayer(t *testing.T) {
	s := createTestPlayersFile(t, "players.json")
	defer os.Remove("players.json")

	type test struct {
		player *models.Player
		err    error
	}

	tests := []test{
		{
			player: &models.Player{
				ID:     1,
				Name:   "Player 1",
				WandID: 1001,
			},
			err: nil,
		},
		{
			player: &models.Player{
				ID:     0,
				Name:   "Player 0",
				WandID: 1000,
			},
			err: errPlayerNotExists,
		},
	}

	for _, tc := range tests {
		// Update player
		err := s.UpdatePlayer(tc.player)
		assert.Equal(t, tc.err, err)

		if err == nil {

			// Get updated player
			player, err := s.GetPlayerByID(tc.player.ID)
			assert.NoError(t, err)
			assert.NotNil(t, player)
			assert.Equal(t, tc.player, player)
		}
	}
}

func TestDeletePlayer(t *testing.T) {
	s := createTestPlayersFile(t, "players.json")
	defer os.Remove("players.json")

	type test struct {
		id  int
		err error
	}

	tests := []test{
		{
			id:  1,
			err: nil,
		},
		{
			id:  0,
			err: errPlayerNotExists,
		},
	}

	for _, tc := range tests {
		// Delete player
		err := s.DeletePlayer(tc.id)
		assert.Equal(t, tc.err, err)

		if tc.err == nil {
			// Get deleted player
			player, err := s.GetPlayerByID(tc.id)
			assert.NoError(t, err)
			assert.Nil(t, player)
		}
	}
}
