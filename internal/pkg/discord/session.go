package discord

import (
	"context"
	"encoding/json"
	"time"

	"github.com/sirupsen/logrus"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"

	"github.com/baileyjm02/jexia-discord-bot/internal/pkg/events"
)

var (
	// Token is the string used for authenticating against Discord's websocket
	Token string
	// Prefix is the value to prefix commands so they are recognised
	Prefix string
)

// Session is the type which holds Discord's current session
type Session struct {
	Ctx      context.Context
	URL      string
	Conn     *websocket.Conn
	Interval time.Duration
	Sequence int
	ID       string
	Token    string
	Prefix   string
}

// startHeartbeat is a helper function that sends a payload every X milliseconds as
// requested by Discord
func (s *Session) startHeartbeat() {
	for {
		time.Sleep(s.Interval * time.Millisecond)
		heartbeat := Heartbeat{
			OPCode: 1,
			Data:   s.Sequence,
		}
		s.send(heartbeat)
	}
}

// send is a helper function to send a payload back over the websocket
func (s *Session) send(payload interface{}) error {
	err := wsjson.Write(s.Ctx, s.Conn, payload)
	if err != nil {
		logrus.Error(err)
	}
	return err
}

// read is a helper function to read all payloads sent across the websocket
func (s *Session) read() (Payload, error) {
	var payload Payload
	err := wsjson.Read(s.Ctx, s.Conn, &payload)
	if err != nil {
		logrus.Fatalf("s.read: %v", err)
		return Payload{}, err
	}
	return payload, nil
}

// spin continuously loops and accepts any incoming payloads
func (s *Session) spin(exit chan bool) {
	for {
		err := s.accept()
		if err != nil {
			s.reconnect()
			exit <- true
			break
		}
	}
}

// accept takes the payload, requests it to be read and passes it
// onto the deploy function where it is handled
func (s *Session) accept() error {
	payload, err := s.read()
	if err != nil {
		logrus.Error(err)
		return nil
	}
	go s.deploy(payload)
	return nil
}

// deploy handles the payload and calls functions based on the payload event
func (s *Session) deploy(payload Payload) {
	s.Sequence = payload.Sequence
	switch payload.OPCode {

	// Reconnect
	case 7:
		{
			go s.reconnect()
		}

	// InvalID Session
	case 9:
		{
			go s.reconnect()
		}

	// Heartbeat ACK
	case 11:
		{
			// Heartbeat ACK
		}

	// Hello
	case 10:
		{
			var mID interface{}
			json.Unmarshal(payload.Data, &mID)
			data := mID.(map[string]interface{})
			s.Interval = time.Duration(data["heartbeat_interval"].(float64))
			go s.startHeartbeat()
			go s.identifySelf()
		}

	// Dispatch
	case 0:
		{
			if payload.Event != "" {
				switch payload.Event {
				case "READY":
					var mID interface{}
					json.Unmarshal(payload.Data, &mID)
					data := mID.(map[string]interface{})
					s.ID = data["session_id"].(string)
				case "RESUMED":
				case "MESSAGE_CREATE":
					var message Message
					json.Unmarshal(payload.Data, &message)
					events.Queue.Publish("discord.message_create", message)
					var err error
					if err != nil {
						logrus.Error(err)
					}
				}
			}
		}
	}
}

// reconnect is a helper function that waits 500 milliseconds and calls the connect
// function again
func (s *Session) reconnect() {
	time.Sleep(500 * time.Millisecond)
	s.Connect()
}

// identifySelf is a helper command which sends the credentials of the bot to
// Discord so it can be authenticated correctly
func (s *Session) identifySelf() {
	data, err := json.Marshal(map[string]interface{}{
		"token": s.Token,
		"properties": map[string]string{
			"$os":      "jexia",
			"$browser": "jexia",
			"$device":  "jexia",
		},
		"presence": map[string]interface{}{
			"game": map[string]interface{}{
				"name": " our GitHub events",
				"type": 2,
			},
			"since": 91879201,
			"afk":   false,
		},
	})
	if err != nil {
		logrus.Error(err)
	}
	payload := Payload{
		2,
		data,
		0,
		"",
	}
	err = s.send(&payload)
	if err != nil {
		logrus.Error(err)
	}
}

// Connect dials Discord's websocket and creates a connection within the session
func (s *Session) Connect() {
	// Set global variables
	Token = s.Token
	Prefix = s.Prefix

	Ctx := context.Background()
	s.Ctx = Ctx

	c, _, err := websocket.Dial(s.Ctx, s.URL, nil)
	if err != nil {
		logrus.Error(err)
		go s.reconnect()
		return
	}
	s.Conn = c

	for {
		exit := make(chan bool)
		go s.spin(exit)
		<-exit
	}

	defer s.Conn.Close(websocket.StatusNormalClosure, "")
}
