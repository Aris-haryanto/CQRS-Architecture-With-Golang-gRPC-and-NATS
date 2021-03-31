package database

import (
	"github.com/deposit-services/api"
	"github.com/deposit-services/nats"
	event "github.com/deposit-services/proto"
)

type EventStore struct {
	stream nats.Stream
}

func (cmd EventStore) CreateEvent(event *event.EventParam) error {
	db.Create(&api.EventStore{
		ID:            event.EventId,
		EventType:     event.EventType,
		AggregateID:   event.AggregateId,
		AggregateType: event.AggregateType,
		EventData:     event.EventData,
		Channel:       event.Channel,
	})

	return nil
}
