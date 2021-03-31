package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	deposit "github.com/deposit-services/proto"
	"github.com/olivere/elastic"
)

type Query struct{}

//in GRPC pattern this query must be read from other than write for now we use elasticsearch
func (qry Query) GetDeposit() []*deposit.Deposit {
	var depositList = []*deposit.Deposit{}

	ctx := context.Background()
	searchSource := elastic.NewSearchSource()
	// searchSource.Query(elastic.NewMatchQuery("name", "Doe"))

	searchService := elasticConn.Search().Index(elIndex).SearchSource(searchSource)

	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
	}

	for _, hit := range searchResult.Hits.Hits {
		var depo *deposit.Deposit
		errJs := json.Unmarshal(hit.Source, &depo)
		if errJs != nil {
			fmt.Println("[Getting Depo][Unmarshal] Err=", err)
		}

		depositList = append(depositList, depo)
	}

	log.Println(depositList)

	return depositList
}
