package handlers

import (
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

func getAmountOfTopicInArticle(w http.ResponseWriter, r *http.Request) {

}
