package bungieapigo

// If an item can have an action performed on it (like "Dismantle"), it will be defined here if you
// care.
type DestinyItemCraftingBlockDefinition struct {

	// A reference to the item definition that is created when crafting with this 'recipe' item.
	OutputItemHash int `json:"outputItemHash"`

	// A list of socket type hashes that describes which sockets are required for crafting with this
	// recipe.
	RequiredSocketTypeHashes []int `json:"requiredSocketTypeHashes"`

	FailedRequirementStrings []string `json:"failedRequirementStrings"`

	// A reference to the base material requirements for crafting with this recipe.
	BaseMaterialRequirements int `json:"baseMaterialRequirements"`

	// A list of 'bonus' socket plugs that may be available if certain requirements are met.
	BonusPlugs []DestinyItemCraftingBlockBonusPlugDefinition `json:"bonusPlugs"`
}
