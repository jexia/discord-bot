package discord

import (
	"os"
)

// StartBot gets the environment variables needed for the Discord gateway, creates
// the new session type and starts the initial websocket connection to Discord's gateway
func StartBot() {
	// Get the environment variables and assign their values
	token := os.Getenv("token")
	prefix := os.Getenv("prefix")

	// Create a new Discord session type with some pre-set variables
	var discordSession = &Session{URL: "wss://gateway.discord.gg/?v=6&encoding=json", Token: token, Prefix: prefix}

	// Connect to the Discord gateway
	go discordSession.Connect()
	go StartSubscriber()
}
