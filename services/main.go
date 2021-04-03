package main

import (
	"context"
	"fmt"
	"net"

	db "github.com/deposit-services/database"
	event "github.com/deposit-services/proto"
	"github.com/pkg/errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/deposit-services/nats"
)

type server struct {
	ev     db.EventStore
	stream nats.Stream
	query  db.Query
}

func main() {

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	event.RegisterEventStoreServer(srv, &server{})
	event.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

//we need to define whenever we not use this
func (s *server) Approve(ctx context.Context, request *event.ApproveParam) (*event.Response, error) {
	return &event.Response{Status: int32(200), Message: string("Approve"), Data: []*event.Deposit{}}, nil
}

//we need to define whenever we not use this
func (s *server) Deposit(ctx context.Context, request *event.DepositParam) (*event.Response, error) {
	return &event.Response{Status: int32(200), Message: string("Deposit"), Data: []*event.Deposit{}}, nil
}

func (s *server) ListDeposit(ctx context.Context, request *event.ListDepositParam) (*event.Response, error) {
	listDeposit := s.query.GetDeposit()

	fmt.Println(listDeposit)
	if listDeposit == nil {
		return &event.Response{Status: int32(500), Message: string("Error get list deposit"), Data: []*event.Deposit{}}, nil
	}

	return &event.Response{Status: int32(200), Message: string("Get All List"), Data: listDeposit}, nil
}

//we need to define whenever we not use this
func (s *server) GetEvents(ctx context.Context, eventData *event.EventFilter) (*event.EventResponse, error) {

	return &event.EventResponse{}, nil
}

func (s *server) CreateEvent(ctx context.Context, eventData *event.EventParam) (*event.ResponseParam, error) {
	//insert kedalam event store dalam hal ini kita create ke table eventstore mysql
	//  dengan data yang dikirim dari client eventData
	createEvent := s.ev.CreateEvent(eventData)

	//publish event melalui nats streaming dengan data yang dikirim dari client eventData
	go s.stream.Publish(eventData.Channel, eventData)

	fmt.Println(createEvent)
	if createEvent == nil {
		return &event.ResponseParam{}, errors.Wrap(createEvent, "error from RPC server")
	}

	return &event.ResponseParam{}, nil
}
