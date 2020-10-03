package main

import (
	"fmt"
	"runtime"

	"github.com/nats-io/nats.go"
)

func main() {
	sc, _ := nats.Connect(nats.DefaultURL)

	// Simple Async Subscriber
	sc.Subscribe("log", func(m *nats.Msg) {
		fmt.Printf("%s\n", m.Data)
	})

	// Keep the connection alive
	runtime.Goexit()

}
