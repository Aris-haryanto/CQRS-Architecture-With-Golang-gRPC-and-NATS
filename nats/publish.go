package nats

import "fmt"

func (pb Stream) Publish(types string, result string) error {
	// Simple Synchronous Publisher
	sc.Publish(types, []byte(fmt.Sprintf("%s", result)))
	return nil
}
