package commands

import (
	"errors"
	"fmt"

	"github.com/baileyjm02/jexia-discord-bot/internal/pkg/discord"
)

// TODO: Add comments
func RepoCommand(m discord.Message, parameters []string) (discord.APIPayload, error) {
	var payload discord.APIPayload
	var val string

	if len(parameters) < 1 {
		return discord.APIPayload{}, errors.New("No parameters passed.")
	}

	if parameters[0] == "add" {
		val = fmt.Sprintf("Added new repo: `%v`", parameters[1])
	}
	if parameters[0] == "remove" {
		val = fmt.Sprintf("Removed repo: `%v`", parameters[1])
	}
	err := payload.Prepare(val, m.ChannelID)
	return payload, err
}

// TODO: Add comments
func init() {
	repo := Command{
		"repo",
		"repo",
		"add or remove repos.",
		"General",
		false,
		map[string]bool{
			"action": true,
			"repo":   true,
		},
		false,
		true,
		RepoCommand,
	}

	repo.Register()
}
