package bungieapigo

type DestinyLoadoutItemComponent struct {
	ItemInstanceId int64 `json:"itemInstanceId,string"`
	PlugItemHashes []int `json:"plugItemHashes"`
}
