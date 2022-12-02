package bungieapigo

type DestinyItemComponentSetOfint64 struct {
	Instances      DictionaryComponentResponseOfint64AndDestinyItemInstanceComponent       `json:"instances"`
	RenderData     DictionaryComponentResponseOfint64AndDestinyItemRenderComponent         `json:"renderData"`
	Stats          DictionaryComponentResponseOfint64AndDestinyItemStatsComponent          `json:"stats"`
	Sockets        DictionaryComponentResponseOfint64AndDestinyItemSocketsComponent        `json:"sockets"`
	ReusablePlugs  DictionaryComponentResponseOfint64AndDestinyItemReusablePlugsComponent  `json:"reusablePlugs"`
	PlugObjectives DictionaryComponentResponseOfint64AndDestinyItemPlugObjectivesComponent `json:"plugObjectives"`
	TalentGrids    DictionaryComponentResponseOfint64AndDestinyItemTalentGridComponent     `json:"talentGrids"`
	PlugStates     DictionaryComponentResponseOfuint32AndDestinyItemPlugComponent          `json:"plugStates"`
	Objectives     DictionaryComponentResponseOfint64AndDestinyItemObjectivesComponent     `json:"objectives"`
	Perks          DictionaryComponentResponseOfint64AndDestinyItemPerksComponent          `json:"perks"`
}
