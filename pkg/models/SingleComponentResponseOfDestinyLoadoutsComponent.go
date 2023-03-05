package bungieapigo

type SingleComponentResponseOfDestinyLoadoutsComponent struct {
	Data    DestinyLoadoutsComponent `json:"data"`
	Privacy ComponentPrivacySetting  `json:"privacy"`

	// If true, this component is disabled.
	Disabled bool `json:"disabled"`
}
