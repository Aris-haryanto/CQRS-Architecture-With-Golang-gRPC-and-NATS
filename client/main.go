package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	pb "github.com/deposit-services/proto"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
)

type server struct{}

const (
	aggregate = "deposit"
	baseGRPC  = "localhost:4040"
)

func CreateDeposit(client pb.EventStoreClient) gin.HandlerFunc {
	handler := func(ctx *gin.Context) {
		amount, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter amount"})
		}

		from := string(ctx.Param("b"))
		if from == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter from name"})
		}

		//data yang ada dalam event
		var depositData pb.DepositParam
		depositData.Amount = int64(amount)
		depositData.From = from

		evID, _ := uuid.NewV4()
		agID, _ := uuid.NewV4()
		jsonData, _ := json.Marshal(depositData)

		//define param2 untuk event
		event := &pb.EventParam{
			EventId:       evID.String(),
			EventType:     "deposit-created",
			AggregateId:   agID.String(),
			AggregateType: aggregate,
			EventData:     string(jsonData),
			Channel:       "deposit-created",
		}

		//create event via grpc ke server
		if res, err := client.CreateEvent(context.Background(), event); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"status":  fmt.Sprint("200"),
				"message": fmt.Sprint("Deposit Created"),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": res})
		}
	}
	return handler
}

func ApproveDeposit(client pb.EventStoreClient) gin.HandlerFunc {
	handler := func(ctx *gin.Context) {
		IdDeposit, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter deposit ID"})
		}

		var approveData pb.ApproveParam
		approveData.IdDeposit = int64(IdDeposit)

		evID, _ := uuid.NewV4()
		agID, _ := uuid.NewV4()
		jsonData, _ := json.Marshal(approveData)

		event := &pb.EventParam{
			EventId:       evID.String(),
			EventType:     "deposit-approve",
			AggregateId:   agID.String(),
			AggregateType: aggregate,
			EventData:     string(jsonData),
			Channel:       "deposit-approve",
		}

		if res, err := client.CreateEvent(
			context.Background(), event); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"status":  fmt.Sprint("200"),
				"message": fmt.Sprint("Deposit Approved"),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": res})
		}
	}
	return handler
}

func ListDeposit(list pb.AddServiceClient) gin.HandlerFunc {
	handler := func(ctx *gin.Context) {
		req := &pb.ListDepositParam{}

		if response, err := list.ListDeposit(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"status":  fmt.Sprint(response.Status),
				"message": fmt.Sprint(response.Message),
				"data":    response.Data,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
	return handler
}

func main() {
	conn, err := grpc.Dial(baseGRPC, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewEventStoreClient(conn)
	list := pb.NewAddServiceClient(conn)

	g := gin.Default()

	g.GET("/deposit/:a/:b", CreateDeposit(client))
	g.GET("/approve/:a", ApproveDeposit(client))
	g.GET("/list", ListDeposit(list))

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
