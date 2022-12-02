package bungieapigo

type DestinyPresentationNodeRecordChildEntry struct {
	RecordHash int `json:"recordHash"`

	// Use this value to sort the presentation node children in ascending order.
	NodeDisplayPriority int `json:"nodeDisplayPriority"`
}
