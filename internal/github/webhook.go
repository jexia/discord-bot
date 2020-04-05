package github

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/baileyjm02/jexia-discord-bot/internal/types"
)

// TODO: Add comments
func Webhook(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	switch event := req.Header.Get("X-GitHub-Event"); event {
	case "release":
		var wh types.Webhook
		err := decoder.Decode(&wh)
		if err != nil {
			panic(err)
		}
		types.Queue.Publish("github.release", wh)
		log.Println("Webhook called")
	default:
		rw.WriteHeader(400) // Return 400 Bad Request.
		return
	}
}