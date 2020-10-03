package main

import (
	"context"
	"net"

	db "github.com/deposit-services/database"
	deposit "github.com/deposit-services/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	command db.Command
	query   db.Query
}

func main() {

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	deposit.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

func (s *server) Deposit(ctx context.Context, request *deposit.DepositParam) (*deposit.Response, error) {
	err := s.command.CreateDeposit(request.GetAmount(), request.GetFrom())
	if err != nil {
		return &deposit.Response{Status: int32(500), Message: string("Error create deposit")}, err
	}

	return &deposit.Response{Status: int32(200), Message: string("Success create deposit")}, nil
}

func (s *server) Approve(ctx context.Context, request *deposit.ApproveParam) (*deposit.Response, error) {
	depositId := s.command.ApproveDeposit(request.GetIdDeposit())
	if depositId == nil {
		return &deposit.Response{Status: int32(500), Message: string("Error approve deposit")}, depositId
	}

	return &deposit.Response{Status: int32(200), Message: string("Success approve deposit")}, nil
}

func (s *server) ListDeposit(ctx context.Context, request *deposit.ListDepositParam) (*deposit.Response, error) {
	listDeposit := s.query.GetDeposit()
	if listDeposit == nil {
		return &deposit.Response{Status: int32(500), Message: string("Error approve deposit"), Data: deposit.Deposit{}}, nil
	}

	return &deposit.Response{Status: int32(200), Message: string("Success approve deposit"), Data: listDeposit}, nil
}
