package database

import (
	"github.com/deposit-services/api"
	"github.com/deposit-services/nats"
)

type Command struct {
	stream nats.Stream
}

func (cmd Command) CreateDeposit(amount int64, from string) error {
	db.Create(&api.Deposit{Amount: amount, From: from})

	cmd.stream.Publish("log", "Success Deposit")
	return nil
}

func (cmd Command) ApproveDeposit(IDdeposit int64) error {
	db.Model(&api.Deposit{}).Where("id = ?", IDdeposit).Update("approve", 1)

	cmd.stream.Publish("log", "Deposit Approve")

	return nil
}
