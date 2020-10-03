package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	deposit "github.com/deposit-services/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := deposit.NewAddServiceClient(conn)
	g := gin.Default()

	g.GET("/deposit/:a/:b", func(ctx *gin.Context) {
		amount, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter amount"})
		}

		from := string(ctx.Param("b"))
		if from == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter from name"})
		}

		req := &deposit.DepositParam{Amount: int64(amount), From: from}
		if response, err := client.Deposit(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"status":  fmt.Sprint(response.Status),
				"message": fmt.Sprint(response.Message),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

	})

	g.GET("/approve/:a", func(ctx *gin.Context) {
		IdDeposit, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter deposit ID"})
		}

		req := &deposit.ApproveParam{IdDeposit: int64(IdDeposit)}

		if response, err := client.Approve(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"status":  fmt.Sprint(response.Status),
				"message": fmt.Sprint(response.Message),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/list", func(ctx *gin.Context) {
		req := &deposit.ListDepositParam{}

		if response, err := client.ListDeposit(ctx, req); err == nil {
			// var io io.Writer
			// a, _ := json.Marshal(response.Data) //get json byte array
			// n := len(a)                         //Find the length of the byte array
			// s := string(a[:n])                  //convert to string

			ctx.JSON(http.StatusOK, gin.H{
				"status":  fmt.Sprint(response.Status),
				"message": fmt.Sprint(response.Message),
				"data":    response.Data,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}