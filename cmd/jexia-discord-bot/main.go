package main

import (
	"github.com/baileyjm02/jexia-discord-bot/internal/bot"
	"github.com/baileyjm02/jexia-discord-bot/internal/pkg/logger"
)

func main() {
	// Format the logger so we can import logrus and have a correct output in other areas
	logger.Format()

	// Start the API process, doing this starts the HTTP server
	// for webhooks and connects the bot to Discord's gateway
	bot.Start()
}
