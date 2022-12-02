package bungieapigo

import "time"

type RewardAvailabilityModel struct {
	HasExistingCode        bool                      `json:"HasExistingCode"`
	RecordDefinitions      []DestinyRecordDefinition `json:"RecordDefinitions"`
	CollectibleDefinitions []CollectibleDefinitions  `json:"CollectibleDefinitions"`
	IsOffer                bool                      `json:"IsOffer"`
	HasOffer               bool                      `json:"HasOffer"`
	OfferApplied           bool                      `json:"OfferApplied"`
	DecryptedToken         string                    `json:"DecryptedToken"`
	IsLoyaltyReward        bool                      `json:"IsLoyaltyReward"`
	ShopifyEndDate         time.Time                 `json:"ShopifyEndDate"`
	GameEarnByDate         time.Time                 `json:"GameEarnByDate"`
	RedemptionEndDate      time.Time                 `json:"RedemptionEndDate"`
}
