package bungieapigo

// Describes the type of progression that a vendor has.
type DestinyVendorProgressionType int

const (

	// The original rank progression from token redemption.
	DestinyVendorProgressionTypeDefault = 0

	// Progression from ranks in ritual content. For example: Crucible (Shaxx), Gambit (Drifter),
	// and Season 13 Battlegrounds (War Table).
	DestinyVendorProgressionTypeRitual = 1

	// A vendor progression with no seasonal refresh. For example: Xur in the Eternity destination
	// for the 30th Anniversary.
	DestinyVendorProgressionTypeNoSeasonalRefresh = 2
)
