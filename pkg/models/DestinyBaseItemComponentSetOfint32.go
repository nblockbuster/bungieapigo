package bungieapigo

type DestinyBaseItemComponentSetOfint32 struct {
	Objectives DictionaryComponentResponseOfint32AndDestinyItemObjectivesComponent `json:"objectives"`
	Perks      DictionaryComponentResponseOfint32AndDestinyItemPerksComponent      `json:"perks"`
}
