package models

type Player struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	WandID   int      `json:"wand_id"`
	House    string   `json:"house"`
	Progress Progress `json:"progress"`
}

type Progress struct {
	SortingHat bool
	Pensieve   bool
}

type Players []*Player
