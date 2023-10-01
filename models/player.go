package models

import "math/rand"

type Player struct {
	ID       int           `json:"id"`
	Name     string        `json:"name"`
	WandID   int           `json:"wandID"`
	House    HogwartsHouse `json:"house"`
	Progress Progress      `json:"progress"`
}

func (p *Player) GetScore() int {
	// TODO: calculate player score based on trivia
	return rand.Intn(100)
}

type Progress struct {
	SortingHat bool
	Pensieve   bool
}

type Players []*Player

type PlayerScore struct {
	Name  string
	House HogwartsHouse
	Score int
}
