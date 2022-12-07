package bungieapigo

// The reasons why an item cannot be equipped, if any. Many flags can be set, or "None" if
type EquipFailureReason int

const (

	// The item is/was able to be equipped.
	EquipFailureReasonNone = 0

	// This is not the kind of item that can be equipped. Did you try equipping Glimmer or something?
	EquipFailureReasonItemUnequippable = 1

	// This item is part of a "unique set", and you can't have more than one item of that same set type
	// equipped at once. For instance, if you already have an Exotic Weapon equipped, you can't equip a
	// second one in another weapon slot.
	EquipFailureReasonItemUniqueEquipRestricted = 2

	// This item has state-based gating that prevents it from being equipped in certain
	// circumstances. For instance, an item might be for Warlocks only and you're a Titan, or it might
	// require you to have beaten some special quest that you haven't beaten yet. Use the additional
	// failure data passed on the item itself to get more information about what the specific failure
	// case was (See DestinyInventoryItemDefinition and DestinyItemInstanceComponent)
	EquipFailureReasonItemFailedUnlockCheck = 4

	// This item requires you to have reached a specific character level in order to equip it, and you
	// haven't reached that level yet.
	EquipFailureReasonItemFailedLevelCheck = 8

	// This item is 'wrapped' and must be unwrapped before being equipped. NOTE: This value used to be
	// called ItemNotOnCharacter but that is no longer accurate.
	EquipFailureReasonItemWrapped = 16

	// This item is not yet loaded and cannot be equipped yet.
	EquipFailureReasonItemNotLoaded = 32

	// This item is block-listed and cannot be equipped.
	EquipFailureReasonItemEquipBlocklisted = 64

	// This item does not meet the loadout requirements for the current activity
	EquipFailureReasonItemLoadoutRequirementNotMet = 128
)
