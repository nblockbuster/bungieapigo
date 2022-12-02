package bungieapigo

type DestinyPlugItemCraftingRequirements struct {
	UnlockRequirements []DestinyPlugItemCraftingUnlockRequirement `json:"unlockRequirements"`

	// If the plug has a known level requirement, it'll be available here.
	RequiredLevel int `json:"requiredLevel"`

	MaterialRequirementHashes []int `json:"materialRequirementHashes"`
}
