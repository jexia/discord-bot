package commands

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	// Import sub-packages containing commands, this isn't the best way to do this
	// but I can't think of any other way that allows us to keep a nice file structure
	// and modular design
	_ "github.com/baileyjm02/jexia-discord-bot/internal/discord/commands/util"

	"github.com/baileyjm02/jexia-discord-bot/internal/types"
)

// StartCommands subscribes to the discord message_create event on the queue and
// waits for a response, also subscribes to the internal send_response event which
// then sends the payload to Discord's API
func StartCommands() {
	// Declare channels to send received events too and subscribe to the channel's events
	githubRelease := make(chan types.DataEvent)
	types.Queue.Subscribe("discord.message_create", githubRelease)
	sendResponse := make(chan types.DataEvent)
	types.Queue.Subscribe("discord.send_response", sendResponse)

	// Create a blocking loop that processes events when they occur
	for {
		select {
		case d := <-githubRelease:
			// Start goroutines to prevent blocks within the loop
			go checkCommand(d.Data.(types.Message))

		case d := <-sendResponse:
			go send(d.Data.(types.DiscordAPIPayload))
		}
	}
}

// checkCommand checks if the message event received contains the prefix and
// a subsequent command from the array of pre-populated commands
func checkCommand(m types.Message) {
	// Checks if the message starts with the command prefix
	if !strings.HasPrefix(m.Content, types.Prefix) {
		// If not, return prematurely finishing the function
		return
	}
	// Remove the prefix to be left with the message content
	content := strings.TrimPrefix(m.Content, types.Prefix)
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
func runCommand(m types.Message, command []string) {
	if cmd, ok := types.Commands[command[0]]; ok {
		defer func() {
			if err := recover(); err != nil {
				var payload types.DiscordAPIPayload
				_ = payload.Prepare(fmt.Sprintf("Unhandled panic occurred: `%v`", err), m.ChannelID)
				types.Queue.Publish("discord.send_response", payload)
				return
			}
		}()
		parameters := command[1:]
		response, err := cmd.Run(m, parameters)
		if err != nil {
			var payload types.DiscordAPIPayload
			_ = payload.Prepare(fmt.Sprintf("Command Error: `%v`",  err.Error()), m.ChannelID)
			types.Queue.Publish("discord.send_response", payload)
			return
		}
		types.Queue.Publish("discord.send_response", response)
	}
}

// TODO: Add comment
func send(payload types.DiscordAPIPayload) {
	req, err := http.NewRequest("POST", "https://discordapp.com/api/v6"+payload.APIURL, bytes.NewBuffer(payload.Payload))
	req.Header.Set("Authorization", "Bot "+types.Token)
	req.Header.Set("User-Agent", "DiscordBot (Jexia, 0.0.1)")
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(resp)

	defer resp.Body.Close()
}
