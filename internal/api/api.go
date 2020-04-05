package api

import (
	"log"
	"os"
	"fmt"
	"net/http"

	"github.com/baileyjm02/jexia-discord-bot/internal/discord"
	"github.com/baileyjm02/jexia-discord-bot/internal/github"
)

// Start initiates the HTTP server for webhooks and requests the bot to start
func Start() {
	// Tell the world we're running 
	log.Println("API process called")

	// Starts separate process that listen for events on the queue
	github.StartWatching()
	discord.StartBot()

	// Add the endpoint for github webhook payloads
	http.HandleFunc("/github", github.Webhook)
	http.HandleFunc("/xyz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there, this is working.")
	})

	// Start the HTTP server ()
	port := os.Getenv("port")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}