package bungieapigo

// An artificial construct of our own creation, to try and put some order on top of Medals and keep
// them from being one giant, unmanageable and unsorted blob of stats.
// Unfortunately, we haven't had time to do this evaluation yet in Destiny 2, so we're short on
// Medal Tiers. This will hopefully be updated over time, if Medals continue to exist.
type DestinyMedalTierDefinition struct {

	// The name of the tier.
	TierName string `json:"tierName"`

	// If you're rendering medals by tier, render them in this order (ascending)
	Order int `json:"order"`

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
