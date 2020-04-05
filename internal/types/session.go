package types
import (
	"context"
	"encoding/json"
	"log"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// TODO: Add comments
var Token string
// TODO: Add comments
var Prefix string

// TODO: Add comments
type Session struct {
	Ctx      context.Context
	URL      string
	Conn     *websocket.Conn
	Interval time.Duration
	Sequence int
	ID       string
	Token    string
}

// TODO: Add comments
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

// TODO: Add comments
func (s *Session) send(payload interface{}) error {
	err := wsjson.Write(s.Ctx, s.Conn, payload)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// TODO: Add comments
func (s *Session) read() (Payload, error) {
	var payload Payload
	err := wsjson.Read(s.Ctx, s.Conn, &payload)
	if err != nil {
		log.Fatalf("s.read: %v", err)
		return Payload{}, err
	}
	return payload, nil
}

// TODO: Add comments
func (s *Session) Spin(exit chan bool) {
	for {
		err := s.Accept()
		if err != nil {
			s.ReConnect()
			exit <- true
			break
		}
	}
}

// TODO: Add comments
func (s *Session) Accept() error {
	payload, err := s.read()
	if err != nil {
		// return err - no need to return this error
		return nil
	}
	// log.Printf("Got message: %#v\n", payload)
	go s.Deploy(payload)
	return nil
}

// TODO: Add comments
func (s *Session) Deploy(payload Payload) {
	s.Sequence = payload.Sequence
	switch payload.OPCode {

	// ReConnect
	case 7:
		{
			go s.ReConnect()
		}

	// InvalID Session
	case 9:
		{
			go s.ReConnect()
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
			go s.IdentifySelf()
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
					Queue.Publish("discord.message_create", message)
					var err error
					if err != nil {
						// Handled elsewhere
					}
				}
			}
		}
	}
}

// TODO: Add comments
func (s *Session) ReConnect() {
	time.Sleep(500 * time.Millisecond)
	s.Connect()
}

// TODO: Add comments
func (s *Session) IdentifySelf() {
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
		log.Fatalln(err)
	}
	payload := Payload{
		2,
		data,
		0,
		"",
	}
	err = s.send(&payload)
	if err != nil {
	}
}

// TODO: Add comments
func (s *Session) Connect() {
	Ctx := context.Background()
	s.Ctx = Ctx

	c, _, err := websocket.Dial(s.Ctx, s.URL, nil)
	if err != nil {
		log.Fatal(err)
	}
	s.Conn = c

	for {
		exit := make(chan bool)
		go s.Spin(exit)
		<-exit
	}

	defer s.Conn.Close(websocket.StatusNormalClosure, "")
}
