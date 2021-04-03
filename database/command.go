package database

import (
	"context"
	"encoding/json"
	"log"

	"github.com/deposit-services/api"
)

type Command struct{}

func (cmd Command) CreateDeposit(amount int64, from string, aggregate_id string) error {
	InsertToElastic(&api.ElDeposit{Amount: amount, From: from, AggregateID: aggregate_id})
	db.Create(&api.Deposit{Amount: amount, From: from, AggregateID: aggregate_id})
	return nil
}

func (cmd Command) ApproveDeposit(AggregateID string) error {
	UpadateToElastic(AggregateID)
	db.Model(&api.Deposit{}).Where("aggregate_id = ?", AggregateID).Update("approve", 1)

	return nil
}

func InsertToElastic(deposit *api.ElDeposit) {
	ctx := context.Background()

	elData, _ := json.Marshal(deposit)
	js := string(elData)
	_, elErr := elasticConn.Index().
		Index(elIndex).
		Id(deposit.AggregateID).
		BodyJson(js).
		Do(ctx)

	if elErr != nil {
		log.Println(elErr)
	}

}
func UpadateToElastic(AggregateID string) {
	ctx := context.Background()

	_, elErr := elasticConn.Update().Index(elIndex).Id(AggregateID).Doc(map[string]interface{}{"approve": 1}).Do(ctx)
	if elErr != nil {
		log.Println(elErr)
	}
}
