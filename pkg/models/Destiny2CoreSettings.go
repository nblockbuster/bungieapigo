package bungieapigo

type Destiny2CoreSettings struct {
	CollectionRootNode                     int    `json:"collectionRootNode"`
	BadgesRootNode                         int    `json:"badgesRootNode"`
	RecordsRootNode                        int    `json:"recordsRootNode"`
	MedalsRootNode                         int    `json:"medalsRootNode"`
	MetricsRootNode                        int    `json:"metricsRootNode"`
	ActiveTriumphsRootNodeHash             int    `json:"activeTriumphsRootNodeHash"`
	ActiveSealsRootNodeHash                int    `json:"activeSealsRootNodeHash"`
	LegacyTriumphsRootNodeHash             int    `json:"legacyTriumphsRootNodeHash"`
	LegacySealsRootNodeHash                int    `json:"legacySealsRootNodeHash"`
	MedalsRootNodeHash                     int    `json:"medalsRootNodeHash"`
	ExoticCatalystsRootNodeHash            int    `json:"exoticCatalystsRootNodeHash"`
	LoreRootNodeHash                       int    `json:"loreRootNodeHash"`
	CraftingRootNodeHash                   int    `json:"craftingRootNodeHash"`
	CurrentRankProgressionHashes           []int  `json:"currentRankProgressionHashes"`
	InsertPlugFreeProtectedPlugItemHashes  []int  `json:"insertPlugFreeProtectedPlugItemHashes"`
	InsertPlugFreeBlockedSocketTypeHashes  []int  `json:"insertPlugFreeBlockedSocketTypeHashes"`
	UndiscoveredCollectibleImage           string `json:"undiscoveredCollectibleImage"`
	AmmoTypeHeavyIcon                      string `json:"ammoTypeHeavyIcon"`
	AmmoTypeSpecialIcon                    string `json:"ammoTypeSpecialIcon"`
	AmmoTypePrimaryIcon                    string `json:"ammoTypePrimaryIcon"`
	CurrentSeasonalArtifactHash            int    `json:"currentSeasonalArtifactHash"`
	CurrentSeasonHash                      int    `json:"currentSeasonHash"`
	SeasonalChallengesPresentationNodeHash int    `json:"seasonalChallengesPresentationNodeHash"`
	FutureSeasonHashes                     []int  `json:"futureSeasonHashes"`
	PastSeasonHashes                       []int  `json:"pastSeasonHashes"`
}
