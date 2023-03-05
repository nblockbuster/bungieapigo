package bungieapigo

type DestinySocialCommendationsComponent struct {
	TotalScore                   int         `json:"totalScore"`
	ScoreDetailValues            []int       `json:"scoreDetailValues"`
	CommendationNodeScoresByHash map[int]int `json:"commendationNodeScoresByHash"`
	CommendationScoresByHash     map[int]int `json:"commendationScoresByHash"`
}
