package handlers

import "net/http"

func TopicCounterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAmountOfTopicInArticle(w,r);
	default:
		log.Println("No implementation for method " + r.Method)

		


}

func getAmountOfTopicInArticle(w http.ResponseWriter, r *http.Request) {

}