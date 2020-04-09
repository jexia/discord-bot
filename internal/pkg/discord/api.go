package discord

import (
	"bytes"
	"net/http"

	"github.com/jexia/discord-bot/internal/pkg/events"
)

// StartSubscriber subscribes to the discord message_create event on the queue and
// waits for a response, also subscribes to the internal send_response event which
// then sends the payload to Discord's API
func StartSubscriber() {
	// Declare channels to send received events to and subscribe to the channel's events
	sendResponse := make(chan events.DataEvent)
	events.Queue.Subscribe("discord.send_response", sendResponse)

	// Create a blocking loop that processes events when they occur
	for {
		select {
		case d := <-sendResponse:
			go sendMessage(d.Data)
		}
	}
}

// sendMessage matches the payload to its type and then uses the Discord API to send it to
// correct channel passed
func sendMessage(payload interface{}) {
	var data []byte
	var url string
	switch v := payload.(type) {
	case APIPayload:
		data = v.Payload
		url = v.APIURL
	case APIEmbedPayload:
		data = v.Payload
		url = v.APIURL
	}
	req, err := http.NewRequest("POST", "https://discordapp.com/api/v6"+url, bytes.NewBuffer(data))
	req.Header.Set("Authorization", "Bot "+Token)
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

	defer resp.Body.Close()
}
