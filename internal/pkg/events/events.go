package events

import (
	"sync"
)

// Queue contains an array of all subscribable channels
var Queue = &EventBus{
	subscribers: map[string]DataChannelSlice{},
}

// EventBus stores the information about subscribers interested for a particular topic
type EventBus struct {
	Bus
	subscribers map[string]DataChannelSlice
	rm          sync.RWMutex
}

// Bus an interface for publishing and subscribing to events
type Bus interface {
	Publish(topic string, data interface{})
	Subscribe(topic string, ch DataChannel)
}

// DataEvent is the type sent over an event, importantly it contains the topic name as there
// are multiple channels
type DataEvent struct {
	Data  interface{}
	Topic string
}

// DataChannel is a channel which can accept an DataEvent
type DataChannel chan DataEvent

// DataChannelSlice is a slice of DataChannels
type DataChannelSlice []DataChannel

// Publish pushes events onto the event channel matching the topic name
func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.rm.RLock()
	if chans, found := eb.subscribers[topic]; found {
		// this is done because the slices refer to same array even though they are passed by value
		// thus we are creating a new slice with our elements thus preserve locking correctly.
		channels := append(DataChannelSlice{}, chans...)
		go func(data DataEvent, dataChannelSlices DataChannelSlice) {
			for _, ch := range dataChannelSlices {
				ch <- data
			}
		}(DataEvent{Data: data, Topic: topic}, channels)
	}
	eb.rm.RUnlock()
}

// Subscribe allows channels to be watched as events to be received
func (eb *EventBus) Subscribe(topic string, ch DataChannel) {
	eb.rm.Lock()
	if prev, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(prev, ch)
	} else {
		eb.subscribers[topic] = append([]DataChannel{}, ch)
	}
	eb.rm.Unlock()
}
