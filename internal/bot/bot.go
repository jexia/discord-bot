package bot

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/jexia/discord-bot/internal/pkg/commands"
	"github.com/jexia/discord-bot/internal/pkg/discord"
	"github.com/jexia/discord-bot/internal/pkg/github"
)

// Start initiates the HTTP server for webhooks and requests the bot to start
func Start() {
	// Tell the world we're running
	logrus.Info("Bot started")

	// Starts separate process that listen for events on the queue
	go github.StartWatching()
	go discord.StartBot()

	// Load in commands and start the subscriber
	go commands.StartSubscriber()

	// Add the endpoint for github webhook payloads
	http.HandleFunc("/github", github.WebhookListener)

	// Start the HTTP server ()
	address := os.Getenv("address")
	logrus.Fatal(http.ListenAndServe(address, nil))
}
