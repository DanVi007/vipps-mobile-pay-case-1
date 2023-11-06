package structs

type WikipediaArticleResponse struct {
	Parse struct {
		Title  string `json:"title"`
		Pageid int    `json:"pageid"`
		Text   struct {
			Asterisk interface{} `json:"*"`
		} `json:"text"`
	}
}
