package bungieapigo

import "time"

type TwitchDropHistoryResponse struct {
	Title       string    `json:"Title"`
	Description string    `json:"Description"`
	CreatedAt   time.Time `json:"CreatedAt"`
	ClaimState  int       `json:"ClaimState"`
}
