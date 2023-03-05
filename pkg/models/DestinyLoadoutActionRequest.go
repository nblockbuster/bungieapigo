package bungieapigo

type DestinyLoadoutActionRequest struct {

	// The index of the loadout for this action request.
	LoadoutIndex int `json:"loadoutIndex"`

	CharacterId    int64                `json:"characterId,string"`
	MembershipType BungieMembershipType `json:"membershipType"`
}
