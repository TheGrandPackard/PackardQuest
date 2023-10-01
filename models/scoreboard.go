package models

type HouseScore struct {
	Name  HogwartsHouse `json:"name"`
	Score int           `json:"int"`
}

type PlayerScore struct {
	Name  string        `json:"name"`
	House HogwartsHouse `json:"house"`
	Score int           `json:"score"`
}
