package bungieapigo

type DestinyCraftableSocketComponent struct {
	PlugSetHash int `json:"plugSetHash"`

	// Unlock state for plugs in the socket plug set definition
	Plugs []DestinyCraftableSocketPlugComponent `json:"plugs"`
}
