package bungieapigo

import "time"

type DestinyPostGameCarnageReportData struct {

	// Date and time for the activity.
	Period time.Time `json:"period"`

	// If this activity has "phases", this is the phase at which the activity was started. This value is
	// only valid for activities before the Beyond Light expansion shipped. Subsequent activities
	// will not have a valid value here.
	StartingPhaseIndex int `json:"startingPhaseIndex"`

	// True if the activity was started from the beginning, if that information is available and the
	// activity was played post Witch Queen release.
	ActivityWasStartedFromBeginning bool `json:"activityWasStartedFromBeginning"`

	// Details about the activity.
	ActivityDetails DestinyHistoricalStatsActivity `json:"activityDetails"`

	// Collection of players and their data for this activity.
	Entries []DestinyPostGameCarnageReportEntry `json:"entries"`

	// Collection of stats for the player in this activity.
	Teams []DestinyPostGameCarnageReportTeamEntry `json:"teams"`
}
