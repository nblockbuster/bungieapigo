package bungieapigo

// If the objective has a known UI label, this enumeration will represent it.
type DestinyObjectiveUiStyle int

const (
	DestinyObjectiveUiStyleNone                        = 0
	DestinyObjectiveUiStyleHighlighted                 = 1
	DestinyObjectiveUiStyleCraftingWeaponLevel         = 2
	DestinyObjectiveUiStyleCraftingWeaponLevelProgress = 3
	DestinyObjectiveUiStyleCraftingWeaponTimestamp     = 4
	DestinyObjectiveUiStyleCraftingMementos            = 5
	DestinyObjectiveUiStyleCraftingMementoTitle        = 6
)
