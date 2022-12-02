package bungieapigo

type CollectibleDefinitions struct {

	// Defines a
	CollectibleDefinition DestinyCollectibleDefinition `json:"CollectibleDefinition"`

	// So much of what you see in Destiny is actually an Item used in a new and creative way. This is the
	// definition for Items in Destiny, which started off as just entities that could exist in your
	// Inventory but ended up being the backing data for so much more: quests, reward previews, slots,
	// and subclasses.
	// In practice, you will want to associate this data with "live" item data from a Bungie.Net
	// Platform call: these definitions describe the item in generic, non-instanced terms: but an
	// actual instance of an item can vary widely from these generic definitions.
	DestinyInventoryItemDefinition DestinyInventoryItemDefinition `json:"DestinyInventoryItemDefinition"`
}
