package commands

import (
	"fmt"
	"time"

	"github.com/baileyjm02/jexia-discord-bot/internal/types"
)

// TODO: Add comment
func pingCommand(m types.Message, parameters []string) (types.DiscordAPIPayload, error) {
	var check time.Time
	var payload types.DiscordAPIPayload
	editedAt, _ := m.EditedAt.Parse()
	sent, _ := m.SentAt.Parse()

	if editedAt != check {
		sent = editedAt
	}

	elapsed := time.Since(sent)
	val := fmt.Sprintf("Pong! `%vms`", int64(elapsed/time.Millisecond))
	err := payload.Prepare(val, m.ChannelID)
	return payload, err
}

// TODO: Add comment
func init() {
	ping := types.Command{
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
