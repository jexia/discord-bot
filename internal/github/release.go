package github

import (
	"fmt"

	"github.com/baileyjm02/jexia-discord-bot/internal/types"
)

// TODO: Add comments
func StartWatchingGithubReleases() {
	githubRelease := make(chan types.DataEvent)
	types.Queue.Subscribe("github.release", githubRelease)
	for {

		select {

		case d := <-githubRelease:
			go handleGithubRelease(d.Data.(types.Webhook))

		}
	}
}

// TODO: Add comments
func handleGithubRelease(wh types.Webhook) {
	if wh.Action != "published" {
		return
	}
	body := "**GraphQL & go-micro transport implementation** \n" +
		"This release introduces two new transport implementations. GraphQL and go-micro.\n" +
		"You are now able to expose flows as GraphQL objects similar to how HTTP endpoints are created.\n\n" +

		"```go\nendpoint 'flow' 'graphql' {\n" +
		"\n\tpath = 'awesome'\n" +
		"}\n" +
		"```\n" +

		"Check out the Graphql and go-micro documentation.\n\n" +

		"**Improved header support**\n" +
		"Further improvements are made for header support.\n" +
		"Input headers could be defined and referenced. Checkout the input specs for more information.\n\n" +

		"**Improved proxy forwarding**\n" +
		"Proxy forwarding implementations have been improved to allow greater flexibility.\n" +
		"It is now possible to forward requests to other services or Maestro instances.\n" +
		"Check out the hubs example for more information.\n"

	DiscordAPIPayload := createDiscordAPIPayload(map[string]interface{}{
		"color": 0xF9B200,
		"title": fmt.Sprintf("New Release (%v)", wh.Release.TagName),
		// "url": "https://jexia.com",
		"author": map[string]string{
			"name":     wh.Repository.FullName,
			"icon_url": wh.Repository.Owner.AvatarURL,
			"url":      wh.Release.HTMLURL,
		},
		"timestamp":   wh.Release.CreatedAt,
		"description": body,
		"fields": []interface{}{
			map[string]interface{}{
				"name":   "Assets",
				"value":  "Some value here",
				"inline": false,
			},
		},
		"footer": map[string]string{
			"text": "Sent via Github",
		},
		"thumbnail": map[string]string{
			"url": wh.Sender.AvatarURL,
		},
	})

	types.Queue.Publish("discord.send_response", DiscordAPIPayload)
}
