package github

import (
	"fmt"
	
	"github.com/jexia/discord-bot/internal/pkg/discord"
	"github.com/jexia/discord-bot/internal/pkg/events"
)

// StartWatchingGithubReleases subscribes to the GitHub release event channel
func StartWatchingGithubReleases() {
	githubRelease := make(chan events.DataEvent)
	events.Queue.Subscribe("github.release", githubRelease)
	for {
		select {

		case d := <-githubRelease:
			go handleGithubRelease(d.Data.(Webhook))
		}
	}
}

// handleGithubRelease is a helper function to create the embed that will be sent
// to Discord containing event data
func handleGithubRelease(wh Webhook) {
	if wh.Action != "published" {
		return
	}
	var payload discord.APIEmbedPayload
	_ = payload.Prepare(map[string]interface{}{
		"color": 0xF9B200,
		"title": fmt.Sprintf("%v (%v)", wh.Release.Name, wh.Release.TagName),
		// "url": "https://jexia.com",
		"author": map[string]string{
			"name":     fmt.Sprintf("New release of %v", wh.Repository.FullName),
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
	}, wh.ChannelID)

	events.Queue.Publish("discord.send_response", payload)
}
