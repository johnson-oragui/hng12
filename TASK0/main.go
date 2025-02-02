package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
)

type Response struct {
	Email         string `json:"email"`
	CurrentTime   string `json:"current_datetime"`
	GitHubRepoURL string `json:"github_url"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// set CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	response := Response{
		Email:         "johnson.oragui",
		CurrentTime:   time.Now().UTC().Format(time.RFC3339),
		GitHubRepoURL: "https://github.com/johnson-oragui/hng12",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/api/v1/task0", handler)
	// logger
	loggedRouter := handlers.LoggingHandler(log.Writer(), http.DefaultServeMux)

	log.Println("Server is starting...")
	http.ListenAndServe(":7001", loggedRouter)
}
