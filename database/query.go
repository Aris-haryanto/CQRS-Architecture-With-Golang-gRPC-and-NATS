package database

import (
	"log"

	deposit "github.com/deposit-services/proto"
)

type Query struct{}

func (get Query) GetDeposit() *deposit.Deposit {
	if err := db.Find(&deposit.Deposit{}).Error; err != nil {
		log.Fatalln(err)
	}

	return &deposit.Deposit{}
}
