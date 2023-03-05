package bungieapigo

type DestinyLoadoutConstantsDefinition struct {

	// Many Destiny*Definition contracts - the "first order" entities of Destiny that have their own
	// tables in the Manifest Database - also have displayable information. This is the base class for
	// that display information.
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayProperties"`

	// This is the same icon as the one in the display properties, offered here as well with a more
	// descriptive name.
	WhiteIconImagePath string `json:"whiteIconImagePath"`

	// This is a color-inverted version of the whiteIconImagePath.
	BlackIconImagePath string `json:"blackIconImagePath"`

	// The maximum number of loadouts available to each character. The loadouts component API
	// response can return fewer loadouts than this, as more loadouts are unlocked by reaching higher
	// Guardian Ranks.
	LoadoutCountPerCharacter int `json:"loadoutCountPerCharacter"`

	// A list of the socket category hashes to be filtered out of loadout item preview displays.
	LoadoutPreviewFilterOutSocketCategoryHashes []int `json:"loadoutPreviewFilterOutSocketCategoryHashes"`

	// A list of the socket type hashes to be filtered out of loadout item preview displays.
	LoadoutPreviewFilterOutSocketTypeHashes []int `json:"loadoutPreviewFilterOutSocketTypeHashes"`

	// A list of the loadout name hashes in index order, for convenience.
	LoadoutNameHashes []int `json:"loadoutNameHashes"`

	// A list of the loadout icon hashes in index order, for convenience.
	LoadoutIconHashes []int `json:"loadoutIconHashes"`

	// A list of the loadout color hashes in index order, for convenience.
	LoadoutColorHashes []int `json:"loadoutColorHashes"`

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
