package models

type HogwartsHouse string

const (
	HogwartsHouseGryffindor = "Gryffindor"
	HogwartsHouseHufflepuff = "Hufflepuff"
	HogwartsHouseRavenclaw  = "Ravenclaw"
	HogwartsHouseSlytherin  = "Slytherin"
)

var (
	HogwartsHouses = []HogwartsHouse{
		HogwartsHouseGryffindor,
		HogwartsHouseHufflepuff,
		HogwartsHouseRavenclaw,
		HogwartsHouseSlytherin,
	}
)
