package bungieAPI

// Gender is a social construct, and as such we have definitions for Genders. Right now there
// happens to only be two, but we'll see what the future holds.
type DestinyGenderDefinition struct {

	// This is a quick reference enumeration for all of the currently defined Genders. We use the
	// enumeration for quicker lookups in related data, like
	// DestinyClassDefinition.genderedClassNames.
	GenderType DestinyGender `json:"genderType"`

	// Many Destiny*Definition contracts - the "first order" entities of Destiny that have their own
	// tables in the Manifest Database - also have displayable information. This is the base class for
	// that display information.
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayProperties"`

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
