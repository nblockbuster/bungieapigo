package bungieapigo

type DestinyItemTransferRequest struct {
	ItemReferenceHash int  `json:"itemReferenceHash"`
	StackSize         int  `json:"stackSize"`
	TransferToVault   bool `json:"transferToVault"`

	// The instance ID of the item for this action request.
	ItemId int64 `json:"itemId,string"`

	CharacterId    int64                `json:"characterId,string"`
	MembershipType BungieMembershipType `json:"membershipType"`
}
