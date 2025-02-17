package bungieapigo

// Many actions relating to items require you to expend materials: - Activating a talent node -
// Inserting a plug into a socket The items will refer to material requirements by a
// materialRequirementsHash in these cases, and this is the definition for those requirements
// in terms of the item required, how much of it is required and other interesting info. This is one
// of the rare/strange times where a single contract class is used both in definitions *and* in
// live data response contracts. I'm not sure yet whether I regret that.
type DestinyMaterialRequirement struct {

	// The hash identifier of the material required. Use it to look up the material's
	// DestinyInventoryItemDefinition.
	ItemHash int `json:"itemHash"`

	// If True, the material will be removed from the character's inventory when the action is
	// performed.
	DeleteOnAction bool `json:"deleteOnAction"`

	// The amount of the material required.
	Count int `json:"count"`

	// If true, the material requirement count value is constant. Since The Witch Queen expansion,
	// some material requirement counts can be dynamic and will need to be returned with an API call.
	CountIsConstant bool `json:"countIsConstant"`

	// If True, this requirement is "silent": don't bother showing it in a material requirements
	// display. I mean, I'm not your mom: I'm not going to tell you you *can't* show it. But we won't show
	// it in our UI.
	OmitFromRequirements bool `json:"omitFromRequirements"`
}
