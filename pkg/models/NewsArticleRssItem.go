package bungieapigo

import "time"

type NewsArticleRssItem struct {
	Title            string    `json:"Title"`
	Link             string    `json:"Link"`
	PubDate          time.Time `json:"PubDate"`
	UniqueIdentifier string    `json:"UniqueIdentifier"`
	Description      string    `json:"Description"`
}
