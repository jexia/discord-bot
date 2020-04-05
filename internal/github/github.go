package github

import (
	"fmt"
	"os"
	"encoding/json"

	"github.com/baileyjm02/jexia-discord-bot/internal/types"
)

// TODO: Add comments
func StartWatching() {
	go StartWatchingGithubReleases()
}

// TODO: Add comments
func createDiscordAPIPayload(embed map[string]interface{}) types.DiscordAPIPayload {
	channel := os.Getenv("channel")
	payload := types.Message{Embed: embed}
	byteArray, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	return types.DiscordAPIPayload{Payload: byteArray, APIURL: fmt.Sprintf("/channels/%v/messages", channel)}
}