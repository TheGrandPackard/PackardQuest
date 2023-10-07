package managers

import "github.com/thegrandpackard/PackardQuest/models"

type PlayerManager interface {
	SetSubscriber(PlayerManagerSubscriber)

	GetPlayers() (models.Players, error)
	GetPlayerByName(playerName string) (*models.Player, error)
	GetPlayerByID(playerID int) (*models.Player, error)
	GetPlayerByWandID(wandID int) (*models.Player, error)
	CreatePlayer(playerName string, wandID int) (*models.Player, error)
	UpdatePlayer(id int, req models.UpdatePlayerRequest) (*models.Player, error)

	GetScoreboards() ([]*models.HouseScore, []*models.PlayerScore, error)
}

type PlayerManagerSubscriber interface {
	OnPlayerUpdate(*models.Player)
}
