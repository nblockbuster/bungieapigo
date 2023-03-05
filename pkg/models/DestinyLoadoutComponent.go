package bungieapigo

type DestinyLoadoutComponent struct {
	ColorHash int                           `json:"colorHash"`
	IconHash  int                           `json:"iconHash"`
	NameHash  int                           `json:"nameHash"`
	Items     []DestinyLoadoutItemComponent `json:"items"`
}
