package bungieapigo

type DestinyCraftableSocketPlugComponent struct {
	PlugItemHash int `json:"plugItemHash"`

	// Index into the unlock requirements to display failure descriptions
	FailedRequirementIndexes []int `json:"failedRequirementIndexes"`
}
