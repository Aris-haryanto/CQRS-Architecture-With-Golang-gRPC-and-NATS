package nats

import (
	"github.com/nats-io/nats.go"
)

type Stream struct{}

var sc *nats.Conn

func init() {
	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}

	sc = conn
}
