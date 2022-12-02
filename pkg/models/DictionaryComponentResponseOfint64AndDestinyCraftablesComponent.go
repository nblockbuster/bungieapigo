package bungieapigo

type DictionaryComponentResponseOfint64AndDestinyCraftablesComponent struct {
	Data    map[int64]DestinyCraftablesComponent `json:"data"`
	Privacy ComponentPrivacySetting              `json:"privacy"`

	// If true, this component is disabled.
	Disabled bool `json:"disabled"`
}
