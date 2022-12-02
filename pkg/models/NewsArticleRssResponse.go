package bungieapigo

type NewsArticleRssResponse struct {
	NewsArticles           []NewsArticleRssItem `json:"NewsArticles"`
	CurrentPaginationToken int                  `json:"CurrentPaginationToken"`
	NextPaginationToken    int                  `json:"NextPaginationToken"`
	ResultCountThisPage    int                  `json:"ResultCountThisPage"`
}
