package github

import (
	"encoding/json"
	"net/http"

	"github.com/baileyjm02/jexia-discord-bot/internal/pkg/events"
)

// TODO: Add comments
func StartWatching() {
	go StartWatchingGithubReleases()
}

// TODO: Add comments
func WebhookListener(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	switch event := req.Header.Get("X-GitHub-Event"); event {
	case "release":
		var wh Webhook
		err := decoder.Decode(&wh)
		if err != nil {
			panic(err)
		}
		events.Queue.Publish("github.release", wh)
	default:
		rw.WriteHeader(400) // Return 400 Bad Request.
		return
	}
}
