package handlers

import (
	"backend/constants"
	"backend/structs"
	"encoding/json"
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

	client := &http.Client{}
	req, err := http.NewRequest("GET", constants.EXTERNAL_TOPIC_ARTICLE_API, nil)
	if err != nil {
		log.Println("could not create request: " + err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	query := req.URL.Query()
	query.Add(constants.EXTERNAL_TOPIC_ARTICLE_API_PARAM, topic)

	req.URL.RawQuery = query.Encode()

	response, err := client.Do(req)
	if err != nil {
		errMsg := "could not get article from external api: "
		log.Println(errMsg + err.Error())
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	var articleResponse structs.WikipediaArticleResponse

	err = json.NewDecoder(response.Body).Decode(&articleResponse)
	if err != nil {
		log.Println("could not unmarshal response body: " + err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	log.Println(articleResponse.Parse.Text.Asterisk)

}
