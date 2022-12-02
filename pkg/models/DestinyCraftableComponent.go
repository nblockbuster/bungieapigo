package bungieapigo

type DestinyCraftableComponent struct {
	Visible bool `json:"visible"`

	// If the requirements are not met for crafting this item, these will index into the list of failure
	// strings.
	FailedRequirementIndexes []int `json:"failedRequirementIndexes"`

	// Plug item state for the crafting sockets.
	Sockets []DestinyCraftableSocketComponent `json:"sockets"`
}
