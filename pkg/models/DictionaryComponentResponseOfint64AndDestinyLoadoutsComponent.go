package bungieapigo

type DictionaryComponentResponseOfint64AndDestinyLoadoutsComponent struct {
	Data    map[int64]DestinyLoadoutsComponent `json:"data"`
	Privacy ComponentPrivacySetting            `json:"privacy"`

	// If true, this component is disabled.
	Disabled bool `json:"disabled"`
}
