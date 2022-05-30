package storers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPlayers(t *testing.T) {
	s, err := NewFileStore("players.json")
	assert.NoError(t, err)

	players, err := s.GetPlayers()
	assert.NoError(t, err)
	assert.NotEmpty(t, players)
}

func TestGetPlayerByName(t *testing.T) {
	s, err := NewFileStore("players.json")
	assert.NoError(t, err)

	type test struct {
		name   string
		exists bool
		err    error
	}

	tests := []test{
		{
			name:   "Player 1",
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
	s, err := NewFileStore("players.json")
	assert.NoError(t, err)

	type test struct {
		wandID int
		exists bool
		err    error
	}

	tests := []test{
		{
			wandID: 1000,
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
