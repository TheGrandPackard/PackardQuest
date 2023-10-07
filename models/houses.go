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

func IsValidHouse(house HogwartsHouse) bool {
	for _, h := range HogwartsHouses {
		if house == h {
			return true
		}
	}

	return false
}
