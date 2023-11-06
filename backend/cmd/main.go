package main

import (
	constants "backend"
	"backend/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc(constants.TOPIC_COUNT_PATH, handlers.TopicCounterHandler)
	http.HandleFunc(constants.TOPIC_COUNT_PATH+"/", handlers.TopicCounterHandler)

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("No PORT environment variable detected, defaulting to 8080")
		port = "8080"
	}

	log.Println("Starting server on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("could not start server: %s\n", err.Error())
	}
}
