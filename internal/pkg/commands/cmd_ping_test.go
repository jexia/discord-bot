package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/jexia/discord-bot/internal/pkg/discord"
)

// toByteArray is a helper function to convert an interface to an array of bytes
func toByteArray(structToConvert interface{}) []byte {
	byteArray, err := json.Marshal(structToConvert)
	if err != nil {
		log.Fatal("Test Preperation Failed: Unable to convert to []byte")
	}
	return byteArray
}

// parse is a helper command to convert a timestap string to a time
func parse(timestamp string) time.Time {
	time, err := time.Parse(time.RFC3339, string(timestamp))
	if err != nil {
		log.Fatal("Test Preperation Failed: Unable to parse timestamp")
	}
	return time
}

// TestPingCommand is a unit test to ensure that the payload returned is correct and the channel ID is passed correctly
func TestPingCommand(t *testing.T) {
	type valuesReturned struct {
		Payload discord.APIPayload
		Err     error
	}
	type test struct {
		input      discord.Message
		parameters []string
		expected   valuesReturned
	}

	var tests = map[string]test{
		"no parameters": {
			discord.Message{ChannelID: "0", SentAt: "2002-10-02T21:20:05.000Z"},
			nil,
			valuesReturned{
				discord.APIPayload{APIURL: "/channels/0/messages", Payload: toByteArray(discord.Message{Content: "Pong! `0ms`"})},
				nil,
			},
		},
		"parameters": {
			discord.Message{ChannelID: "0", SentAt: "2002-10-02T21:20:05.000Z"},
			[]string{"Some", "Random", "Parameters"},
			valuesReturned{
				discord.APIPayload{APIURL: "/channels/0/messages", Payload: toByteArray(discord.Message{Content: "Pong! `0ms`"})},
				nil,
			},
		},
		"10 millisecond ping": {
			discord.Message{ChannelID: "0", SentAt: "2002-10-02T21:20:04.990Z"},
			nil,
			valuesReturned{
				discord.APIPayload{APIURL: "/channels/0/messages", Payload: toByteArray(discord.Message{Content: "Pong! `10ms`"})},
				nil,
			},
		},
		"1 second ping": {
			discord.Message{ChannelID: "0", SentAt: "2002-10-02T21:20:04.00Z"},
			nil,
			valuesReturned{
				discord.APIPayload{APIURL: "/channels/0/messages", Payload: toByteArray(discord.Message{Content: "Pong! `1000ms`"})},
				nil,
			},
		},
		"normal special characters": {
			discord.Message{ChannelID: "$&\"$&", SentAt: "2002-10-02T21:20:05.000Z"},
			nil,
			valuesReturned{
				discord.APIPayload{APIURL: "/channels/$&\"$&/messages", Payload: toByteArray(discord.Message{Content: "Pong! `0ms`"})},
				nil,
			},
		},
		"extra special characters": {
			discord.Message{ChannelID: "		", SentAt: "2002-10-02T21:20:05.000Z"},
			nil,
			valuesReturned{
				discord.APIPayload{APIURL: "/channels/		/messages", Payload: toByteArray(discord.Message{Content: "Pong! `0ms`"})},
				nil,
			},
		},
	}

	// Loop through and run each test case
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			val, err := pingCommand(test.input, test.parameters, parse("2002-10-02T21:20:05.000Z"))
			output := valuesReturned{val, err}
			if !cmp.Equal(output, test.expected) {
				t.Error(fmt.Sprintf("Test Failed: %v", name))
			}
		})
	}
}
