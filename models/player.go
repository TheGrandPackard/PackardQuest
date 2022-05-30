package models

type Player struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	WandID int    `json:"wand_id"`
}

type Players []*Player
