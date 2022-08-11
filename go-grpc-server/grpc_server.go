package main

import (
	"context"
	"github.com/mahendrabagul/polyglot-grpc-demo/go-grpc-server/addition"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type additionServer struct {
}

func (s *additionServer) Add(c context.Context, addRequest *addition.AddRequest) (*addition.AddResponse, error) {
	log.Printf("received add request for %d and %d\n", addRequest.Number, addRequest.AnotherNumber)
	result := addRequest.Number + addRequest.AnotherNumber
	response := addition.AddResponse{
		Sum: int64(result),
	}
	return &response, nil
}

func newAddServer() *additionServer {
	return &additionServer{}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:10000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	addition.RegisterAdditionServer(grpcServer, newAddServer())
	go func() {
		err := grpcServer.Serve(lis)
		log.Fatalf("failed to start server: %v", err)
	}()

	log.Println("server started and listening at localhost:10000")

	<-interrupt()
}

func interrupt() chan os.Signal {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	return interrupt
}
