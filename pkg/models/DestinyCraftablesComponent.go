package bungieapigo

type DestinyCraftablesComponent struct {

	// A map of craftable item hashes to craftable item state components.
	Craftables map[int]DestinyCraftableComponent `json:"craftables"`

	// The hash for the root presentation node definition of craftable item categories.
	CraftingRootNodeHash int `json:"craftingRootNodeHash"`
}
