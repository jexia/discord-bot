package commands

import (
	"github.com/baileyjm02/jexia-discord-bot/internal/pkg/discord"
)

// Command is a type which contains namely the runable function and name of the command
type Command struct {
	Name        string
	Usage       string
	Description string
	Category    string
	NeedArgs    bool
	Args        map[string]bool
	OwnerOnly   bool
	Enabled     bool
	Run         func(m discord.Message, parameters []string) (discord.APIPayload, error)
}

var (
	// Commands is an array of command objects once they have been registered
	Commands map[string]Command
)

// Register registers the commands to the Commands array by command name
func (c Command) Register() {
	if Commands == nil {
		Commands = make(map[string]Command)
	}

	if c.Enabled {
		Commands[c.Name] = c
	}
}
