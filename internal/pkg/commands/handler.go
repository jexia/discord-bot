package commands

import (
	"fmt"
	"strings"

	"github.com/baileyjm02/jexia-discord-bot/internal/pkg/discord"
	"github.com/baileyjm02/jexia-discord-bot/internal/pkg/events"
)

// StartSubscriber subscribes to the discord message_create event on the queue and
// waits for a response, also subscribes to the internal send_response event which
// then sends the payload to Discord's API
func StartSubscriber() {
	// Declare channels to send received events too and subscribe to the channel's events
	githubRelease := make(chan events.DataEvent)
	events.Queue.Subscribe("discord.message_create", githubRelease)

	// Create a blocking loop that processes events when they occur
	for {
		select {
		case d := <-githubRelease:
			// Start goroutines to prevent blocks within the loop
			go checkCommand(d.Data.(discord.Message))
		}
	}
}

// checkCommand checks if the message event received contains the prefix and
// a subsequent command from the array of pre-populated commands
func checkCommand(m discord.Message) {
	// Checks if the message starts with the command prefix
	if !strings.HasPrefix(m.Content, discord.Prefix) {

		// If not, return prematurely finishing the function
		return
	}
	// Remove the prefix to be left with the message content
	content := strings.TrimPrefix(m.Content, discord.Prefix)
	// Split the message by whitespace, allows us to access the command
	// and subsequent inputs
	command := strings.Split(content, " ")

	// Check if the command array contains at least one value, this would be
	// the command name
	if len(command) < 1 {
		// If not, return prematurely finishing the function
		return
	}

	// Run the command in a non-blocking function, passing the message and
	// the command parameters, containing the command name
	go runCommand(m, command)
}

// runCommand takes a array of strings and checks if the first value is a command.
// If a command is found, it is then run with the remaining parameters of the array.
func runCommand(m discord.Message, command []string) {
	// Searches for command if it is registered in the array
	if cmd, ok := Commands[command[0]]; ok {
		// Handles an unexpected panic error
		defer func() {
			if err := recover(); err != nil {
				handleError(m, fmt.Sprintf("Unhandled panic occurred: `%v`", err))
				return
			}
		}()
		// Removes the command name from the array of parameters
		parameters := command[1:]
		// Runs the command using the function passed through into the type
		response, err := cmd.Run(m, parameters)
		// Handles an _expected_ error
		if err != nil {
			handleError(m, fmt.Sprintf("Command Error: `%v`", err.Error()))
			return
		}
		// Publishes the response to be sent to discord
		events.Queue.Publish("discord.send_response", response)
	}
}

// handleError is a helper function for returning useful error information to the user
// when a command fails to exectue correctly
func handleError(m discord.Message, errorText string) {
	var payload discord.APIPayload
	_ = payload.Prepare(errorText, m.ChannelID)
	events.Queue.Publish("discord.send_response", payload)
}
