package bungieapigo

type DestinySocialCommendationDefinition struct {

	// Many Destiny*Definition contracts - the "first order" entities of Destiny that have their own
	// tables in the Manifest Database - also have displayable information. This is the base class for
	// that display information.
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayProperties"`

	CardImagePath string `json:"cardImagePath"`

	// Represents a color whose RGBA values are all represented as values between 0 and 255.
	Color DestinyColor `json:"color"`

	DisplayPriority            int `json:"displayPriority"`
	ActivityGivingLimit        int `json:"activityGivingLimit"`
	ParentCommendationNodeHash int `json:"parentCommendationNodeHash"`

	// The display properties for the the activities that this commendation is available in.
	DisplayActivities []DestinyDisplayPropertiesDefinition `json:"displayActivities"`

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
