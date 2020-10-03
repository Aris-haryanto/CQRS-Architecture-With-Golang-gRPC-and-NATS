package database

import (
	"log"

	"github.com/deposit-services/nats"
	deposit "github.com/deposit-services/proto"
)

type Query struct {
	stream nats.Stream
}

func (qry Query) GetDeposit() []*deposit.Deposit {
	var depositList = []*deposit.Deposit{}
	if err := db.Find(&depositList).Error; err != nil {
		log.Fatalln(err)
	}

	qry.stream.Publish("log", "Success Get All List")

	return depositList
}
