package bungieapigo

type DestinyArrangementRegionFilterDefinition struct {
	ArtArrangementRegionHash    int         `json:"artArrangementRegionHash"`
	ArtArrangementRegionIndex   int         `json:"artArrangementRegionIndex"`
	StatHash                    int         `json:"statHash"`
	ArrangementIndexByStatValue map[int]int `json:"arrangementIndexByStatValue"`
}
