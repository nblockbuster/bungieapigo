package bungieapigo

type PartnerRewardHistoryResponse struct {
	PartnerOffers []PartnerOfferSkuHistoryResponse `json:"PartnerOffers"`
	TwitchDrops   []TwitchDropHistoryResponse      `json:"TwitchDrops"`
}
