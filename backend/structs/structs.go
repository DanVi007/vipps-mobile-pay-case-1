package structs

// WikipediaArticleResponse is the struct for the response from the wikipedia api
type WikipediaArticleResponse struct {
	Parse struct {
		Title  string `json:"title"`
		Pageid int    `json:"pageid"`
		Text   struct {
			Asterisk string `json:"*"`
		} `json:"text"`
	}
}
