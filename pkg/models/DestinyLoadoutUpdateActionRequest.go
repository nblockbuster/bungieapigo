package bungieapigo

type DestinyLoadoutUpdateActionRequest struct {
	ColorHash int `json:"colorHash"`
	IconHash  int `json:"iconHash"`
	NameHash  int `json:"nameHash"`

	// The index of the loadout for this action request.
	LoadoutIndex int `json:"loadoutIndex"`

	CharacterId    int64                `json:"characterId,string"`
	MembershipType BungieMembershipType `json:"membershipType"`
}
