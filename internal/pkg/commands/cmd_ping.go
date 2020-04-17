package commands

import (
	"fmt"
	"time"

	"github.com/jexia/discord-bot/internal/pkg/discord"
)

// pingCommand is a simple command returning the response time of the bot
func pingCommand(m discord.Message, parameters []string, received time.Time) (discord.APIPayload, error) {
	var payload discord.APIPayload
	sent, _ := m.SentAt.Parse()

	elapsed := received.Sub(sent)
	val := fmt.Sprintf("Pong! `%vms`", int64(elapsed/time.Millisecond))
	err := payload.Prepare(val, m.ChannelID)
	return payload, err
}

// This init command registers the ping command to the array of Commands so it can be called
func init() {
	ping := Command{
		"ping",
		"ping",
		"see how long the bot takes to respond.",
		"General",
		false,
		map[string]bool{},
		false,
		true,
		pingCommand,
	}

	ping.Register()
}
