package nats

import (
	event "github.com/deposit-services/proto"
)

func (pb Stream) Publish(types string, event *event.EventParam) error {
	// Simple Synchronous Publisher

	//publish to service receiver
	sc.Publish(types, []byte(event.EventData))

	//publish to log
	sc.Publish("log", []byte(event.EventData))
	return nil
}
