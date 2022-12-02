package bungieapigo

type DestinyBaseItemComponentSetOfint64 struct {
	Objectives DictionaryComponentResponseOfint64AndDestinyItemObjectivesComponent `json:"objectives"`
	Perks      DictionaryComponentResponseOfint64AndDestinyItemPerksComponent      `json:"perks"`
}
