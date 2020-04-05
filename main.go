package main

import (
	"github.com/baileyjm02/jexia-discord-bot/internal/api"
)

func main() {
	// Start the API process, doing this starts the HTTP server
	// for webhooks and connects the bot to Discord's gateway
	api.Start()
}
