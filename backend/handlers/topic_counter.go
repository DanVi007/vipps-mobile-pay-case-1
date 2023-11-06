package handlers

import (
	"backend/constants"
	"backend/structs"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func TopicCounterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAmountOfTopicInArticle(w, r)
	default:
		log.Println("No implementation for method " + r.Method)
		http.Error(w, "No implementation for method "+r.Method, http.StatusNotImplemented)
	}
}

// Retrieves the amount of times a topic appears in an article.
//
// example request:
// GET /topic_count?topic=coronavirus
//
// example response:
// 200 OK
// 5
func getAmountOfTopicInArticle(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	topic := queryValues.Get("topic")
	if topic == "" {
		log.Println("no topic query parameter provided")
		http.Error(w, "bad request: topic query parameter not provided", http.StatusBadRequest)
		return
	}

	response, err := http.Get(constants.EXTERNAL_TOPIC_ARTICLE_API + topic)
	if err != nil {
		errMsg := "could not get article from external api: "
		log.Println(errMsg + err.Error())
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		errMsg := "could not read response body: "
		log.Println(errMsg + err.Error())
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	var articleResponse structs.WikipediaArticleResponse
	err = json.Unmarshal(responseBody, &articleResponse)
	if err != nil {
		log.Println("could not unmarshal response body: " + err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	log.Println(articleResponse.Parse.Text)

}
