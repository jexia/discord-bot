package commands

import (
	"fmt"
	"strings"
	"github.com/sirupsen/logrus"

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
	logrus.Println("Checking Command")
	logrus.Println("prefix is: "+discord.Prefix)
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

// TODO: Refractor into smaller functions and improve error handling when the error is handled (err.New)
func runCommand(m discord.Message, command []string) {
	logrus.Println("Running command: "+command[0] )

	if cmd, ok := Commands[command[0]]; ok {
		defer func() {
			if err := recover(); err != nil {
				var payload discord.APIPayload
				_ = payload.Prepare(fmt.Sprintf("Unhandled panic occurred: `%v`", err), m.ChannelID)
				events.Queue.Publish("discord.send_response", payload)
				return
			}
		}()
		parameters := command[1:]
		response, err := cmd.Run(m, parameters)
		logrus.Println(response)
		if err != nil {
			var payload discord.APIPayload
			_ = payload.Prepare(fmt.Sprintf("Command Error: `%v`", err.Error()), m.ChannelID)
			events.Queue.Publish("discord.send_response", payload)
			return
		}
		logrus.Println("Response sent")
		events.Queue.Publish("discord.send_response", response)
	}
}

// TODO: Create 'handleError' function
func handleError() {}
