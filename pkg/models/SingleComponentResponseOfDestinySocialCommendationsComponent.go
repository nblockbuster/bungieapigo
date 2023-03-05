package bungieapigo

type SingleComponentResponseOfDestinySocialCommendationsComponent struct {
	Data    DestinySocialCommendationsComponent `json:"data"`
	Privacy ComponentPrivacySetting             `json:"privacy"`

	// If true, this component is disabled.
	Disabled bool `json:"disabled"`
}
