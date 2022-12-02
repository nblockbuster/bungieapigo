package bungieapigo

type BungieRewardDisplay struct {
	UserRewardAvailabilityModel UserRewardAvailabilityModel `json:"UserRewardAvailabilityModel"`
	ObjectiveDisplayProperties  RewardDisplayProperties     `json:"ObjectiveDisplayProperties"`
	RewardDisplayProperties     RewardDisplayProperties     `json:"RewardDisplayProperties"`
}
