package helpers

import (
	"log"

	gevent "github.com/gookit/event"
)

type EventType string

const (
	DataReceived EventType = "DATA_RECEIVED"
)

type Event struct {
	events []EventType
}

func NewEventHelper() *Event {
	return &Event{
		events: []EventType{DataReceived},
	}
}

func (e Event) Subscribe(event EventType) {
	gevent.On(string(event), gevent.ListenerFunc(func(e gevent.Event) error {
		rocket := e.Get("rocket")
		log.Printf("Event received for %s", rocket)
		return nil
	}))
}

func (e Event) Trigger(event EventType, args map[string]any) {
	gevent.Trigger(string(event), args)
}
