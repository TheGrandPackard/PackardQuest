package models

type HouseScore struct {
	Name  HogwartsHouse `json:"name"`
	Score int           `json:"score"`
}

type PlayerScore struct {
	Name  string        `json:"name"`
	House HogwartsHouse `json:"house"`
	Score int           `json:"score"`
}
