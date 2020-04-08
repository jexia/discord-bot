package github

import (
	"encoding/json"
	"net/http"

	"github.com/baileyjm02/jexia-discord-bot/internal/pkg/events"
)

// StartWatching is a helper function that will call all events allowing
// then to subscribe then to their related event
func StartWatching() {
	go StartWatchingGithubReleases()
}

// WebhookListener is the endpoint for which the GitHub webhook events should be sent
// It checks if we support the sent event and handles it accordingly
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
		rw.WriteHeader(204)
		return
	default:
		rw.WriteHeader(501) // Return 501 Not Implemented Request as we don't support that function
		return
	}
}
