package bungieapigo

type DestinyPresentationNodeCraftableChildEntry struct {
	CraftableItemHash int `json:"craftableItemHash"`

	// Use this value to sort the presentation node children in ascending order.
	NodeDisplayPriority int `json:"nodeDisplayPriority"`
}
