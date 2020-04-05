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
	
	DiscordAPIPayload := createDiscordAPIPayload(map[string]interface{}{
		"color": 0xF9B200,
		"title": fmt.Sprintf("%v (%v)", wh.Release.Name, wh.Release.TagName),
		// "url": "https://jexia.com",
		"author": map[string]string{
			"name": fmt.Sprintf("New release of %v", wh.Repository.FullName),
			"icon_url": wh.Repository.Owner.AvatarURL,
			"url":      wh.Release.HTMLURL,
		},
		"timestamp":   wh.Release.PublishedAt,
		"description": wh.Release.Body,
		// TODO: Loop though download links / files
		// "fields": []interface{}{
		// 	map[string]interface{}{
		// 		"name":   "Assets",
		// 		"value":  "Some value here",
		// 		"inline": false,
		// 	},
		// },
		"footer": map[string]string{
			"text": "Sent via Github",
		},
	})

	types.Queue.Publish("discord.send_response", DiscordAPIPayload)
}
