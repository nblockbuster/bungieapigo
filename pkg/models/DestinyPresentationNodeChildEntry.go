package bungieapigo

type DestinyPresentationNodeChildEntry struct {
	PresentationNodeHash int `json:"presentationNodeHash"`

	// Use this value to sort the presentation node children in ascending order.
	NodeDisplayPriority int `json:"nodeDisplayPriority"`
}
