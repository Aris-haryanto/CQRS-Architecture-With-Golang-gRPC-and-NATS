package main

import (
	"encoding/json"
	"log"
	"runtime"

	db "github.com/deposit-services/database"
	deposit "github.com/deposit-services/proto"
	"github.com/nats-io/nats.go"
)

const (
	subscribeChannel = "deposit-approve"

	grpcUri = "localhost:4040"
)

func main() {
	sc, _ := nats.Connect(nats.DefaultURL)

	// Simple Async Subscriber
	sc.Subscribe(subscribeChannel, func(m *nats.Msg) {
		approveParam := deposit.ApproveParam{}
		err := json.Unmarshal(m.Data, &approveParam)
		if err != nil {
			log.Println(err)
		}

		log.Println(approveParam)
		isErr := db.Command{}.ApproveDeposit(approveParam.AggregateId)
		if isErr != nil {
			log.Println(isErr)
		}
	})

	// Keep the connection alive
	runtime.Goexit()
}
