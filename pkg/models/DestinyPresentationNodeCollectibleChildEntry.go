package bungieapigo

type DestinyPresentationNodeCollectibleChildEntry struct {
	CollectibleHash int `json:"collectibleHash"`

	// Use this value to sort the presentation node children in ascending order.
	NodeDisplayPriority int `json:"nodeDisplayPriority"`
}
