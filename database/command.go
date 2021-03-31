package database

import (
	"context"
	"encoding/json"
	"log"

	"github.com/deposit-services/api"
)

type Command struct{}

func (cmd Command) CreateDeposit(amount int64, from string) error {
	InsertToElastic(&api.ElDeposit{Amount: amount, From: from})
	db.Create(&api.Deposit{Amount: amount, From: from})
	return nil
}

func (cmd Command) ApproveDeposit(IDdeposit int64) error {
	db.Model(&api.Deposit{}).Where("id = ?", IDdeposit).Update("approve", 1)

	return nil
}

func InsertToElastic(deposit *api.ElDeposit) {
	ctx := context.Background()

	elData, _ := json.Marshal(deposit)
	js := string(elData)
	_, elErr := elasticConn.Index().
		Index(elIndex).
		BodyJson(js).
		Do(ctx)

	if elErr != nil {
		log.Println(elErr)
	}

}
