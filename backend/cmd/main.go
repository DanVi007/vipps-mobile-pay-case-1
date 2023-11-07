package main

import (
	"backend/constants"
	"backend/handlers"
	"log"
	"net/http"
	"os"
)

// main function and entry point for application
// handler paths are defined here
// server port is defined here
func main() {
	http.HandleFunc(constants.TOPIC_COUNTER_PATH, handlers.TopicCounterHandler)
	http.HandleFunc(constants.TOPIC_COUNTER_PATH+"/", handlers.TopicCounterHandler)

	port := os.Getenv("PORT")
	if port == "" {
		defaultPort := "8080"
		log.Println("No PORT environment variable detected, defaulting to " + defaultPort)
		port = defaultPort
	}

	log.Println("Starting server on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("could not start server: %s\n", err.Error())
	}
}
