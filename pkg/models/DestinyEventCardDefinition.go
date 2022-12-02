package bungieapigo

// Defines the properties of an 'Event Card' in Destiny 2, to coincide with a seasonal event for
// additional challenges, premium rewards, a new seal, and a special title. For example:
// Solstice of Heroes 2022.
type DestinyEventCardDefinition struct {

	// Many Destiny*Definition contracts - the "first order" entities of Destiny that have their own
	// tables in the Manifest Database - also have displayable information. This is the base class for
	// that display information.
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayProperties"`

	LinkRedirectPath string `json:"linkRedirectPath"`

	// Represents a color whose RGBA values are all represented as values between 0 and 255.
	Color DestinyColor `json:"color"`

	Images                       DestinyEventCardImages `json:"images"`
	TriumphsPresentationNodeHash int                    `json:"triumphsPresentationNodeHash"`
	SealPresentationNodeHash     int                    `json:"sealPresentationNodeHash"`
	TicketCurrencyItemHash       int                    `json:"ticketCurrencyItemHash"`
	TicketVendorHash             int                    `json:"ticketVendorHash"`
	TicketVendorCategoryHash     int                    `json:"ticketVendorCategoryHash"`
	EndTime                      int64                  `json:"endTime,string"`

	// The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not
	// globally.
	// When entities refer to each other in Destiny content, it is this hash that they are referring to.
	Hash int `json:"hash"`

	// The index of the entity as it was found in the investment tables.
	Index int `json:"index"`

	// If this is true, then there is an entity with this identifier/type combination, but BNet is not
	// yet allowed to show it. Sorry!
	Redacted bool `json:"redacted"`
}
