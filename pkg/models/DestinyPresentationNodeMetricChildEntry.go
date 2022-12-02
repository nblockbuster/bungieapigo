package bungieapigo

type DestinyPresentationNodeMetricChildEntry struct {
	MetricHash int `json:"metricHash"`

	// Use this value to sort the presentation node children in ascending order.
	NodeDisplayPriority int `json:"nodeDisplayPriority"`
}
