package bungieapigo

// If you want to report a player causing trouble in a game, this request will let you report that
// player and the specific PGCR in which the trouble was caused, along with why.
// Please don't do this just because you dislike the person! I mean, I know people will do it
// anyways, but can you like take a good walk, or put a curse on them or something? Do me a solid and
// reconsider.
// Note that this request object doesn't have the actual PGCR ID nor your Account/Character ID in
// it. We will infer that information from your authentication information and the PGCR ID that
// you pass into the URL of the reporting endpoint itself.
type DestinyReportOffensePgcrRequest struct {

    // So you've decided to report someone instead of cursing them and their descendants. Well, okay
    // then. This is the category or categorie(s) of infractions for which you are reporting the user.
    // These are hash identifiers that map to DestinyReportReasonCategoryDefinition entries.
    ReasonCategoryHashes []int `json:"reasonCategoryHashes"`


    // If applicable, provide a more specific reason(s) within the general category of problems
    // provided by the reasonHash. This is also an identifier for a reason. All reasonHashes provided
    // must be children of at least one the reasonCategoryHashes provided.
    ReasonHashes []int `json:"reasonHashes"`


    // Within the PGCR provided when calling the Reporting endpoint, this should be the character ID
    // of the user that you thought was violating terms of use. They must exist in the PGCR provided.
    OffendingCharacterId int64 `json:"offendingCharacterId"`

}

