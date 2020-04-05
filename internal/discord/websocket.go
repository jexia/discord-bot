package discord

import (
	"os"

	"github.com/baileyjm02/jexia-discord-bot/internal/discord/commands"
	"github.com/baileyjm02/jexia-discord-bot/internal/types"
)

// StartBot gets the environment variables needed for the Discord gateway, creates
// the new session type and starts the initial websocket connection to Discord's gateway
func StartBot() {
	// Get the environment variables and assing them to the types package so when other funcions 
	// request the value they are given the correct values of token and prefix we want instead of nil
	types.Token = os.Getenv("token")
	types.Prefix = os.Getenv("prefix")

	// Create a new Discord session type with some pre-set variables
	var newSession = &types.Session{URL: "wss://gateway.discord.gg/?v=6&encoding=json", Token: types.Token}

	// Connect to the Discord gateway
	go newSession.Connect()

	// Load in commands and start the subscriber
	go commands.StartCommands()
}