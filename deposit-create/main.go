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
	subscribeChannel = "deposit-created"

	grpcUri = "localhost:4040"
)

func main() {
	sc, _ := nats.Connect(nats.DefaultURL)

	// subscribe chanel sesuai yang dikirim dari service client dalam contoh ini "deposit-created"
	sc.Subscribe(subscribeChannel, func(m *nats.Msg) {
		depositParam := deposit.DepositParam{}
		err := json.Unmarshal(m.Data, &depositParam)
		if err != nil {
			log.Println(err)
		}

		log.Println(depositParam)

		//insert deposit yang dikirim dari client ke database dan ke elastic seach sebagai pattern CQRS
		isErr := db.Command{}.CreateDeposit(depositParam.Amount, depositParam.From)
		if isErr != nil {
			log.Println(isErr)
		}
	})

	// Keep the connection alive
	runtime.Goexit()
}
