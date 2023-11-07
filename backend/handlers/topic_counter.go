package handlers

import (
	"backend/constants"
	"backend/structs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/k3a/html2text"
)

// Handler for /topic_counter
func TopicCounterHandler(w http.ResponseWriter, r *http.Request) {
	// simple enabling of cors to frontend
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
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
	// retrieve topic from query parameter
	queryValues := r.URL.Query()
	topic := queryValues.Get("topic")
	if topic == "" {
		log.Println("no topic query parameter provided")
		http.Error(w, "bad request: topic query parameter not provided", http.StatusBadRequest)
		return
	}

	// get article from external api
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

	// decode response into struct
	var articleResponse structs.WikipediaArticleResponse

	err = json.NewDecoder(response.Body).Decode(&articleResponse)
	defer response.Body.Close()
	if err != nil {
		log.Println("could not unmarshal response body: " + err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// decode html article into plain text
	htmlArticle := articleResponse.Parse.Text.Asterisk

	articleInPlainText := html2text.HTML2Text(htmlArticle)

	// count amount of times topic appears in article plain text
	count := strings.Count(articleInPlainText, topic)

	// write to user
	_, err = fmt.Fprint(w, count)
	if err != nil {
		log.Println("could not write response: " + err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
