package bungieapigo

type DestinyItemStateRequest struct {
	State bool `json:"state"`

	// The instance ID of the item for this action request.
	ItemId int64 `json:"itemId,string"`

	CharacterId    int64                `json:"characterId,string"`
	MembershipType BungieMembershipType `json:"membershipType"`
}
