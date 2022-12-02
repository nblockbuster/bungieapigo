package bungieapigo

type DestinyBaseItemComponentSetOfuint32 struct {
	Objectives DictionaryComponentResponseOfuint32AndDestinyItemObjectivesComponent `json:"objectives"`
	Perks      DictionaryComponentResponseOfuint32AndDestinyItemPerksComponent      `json:"perks"`
}
