package bungieapigo

type UserRewardAvailabilityModel struct {
	AvailabilityModel  RewardAvailabilityModel `json:"AvailabilityModel"`
	IsAvailableForUser bool                    `json:"IsAvailableForUser"`
	IsUnlockedForUser  bool                    `json:"IsUnlockedForUser"`
}
