package main

import (
	"context"
	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"polyglot-grpc-demo/addition"
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

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// load server certificate and private key
	serverCert, err := tls.LoadX509KeyPair("certs/server-cert.pem", "certs/server-key.pem")
	if err != nil {
		return nil, err
	}

	// create credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:10000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatalf("failed to generate tls credentials: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))
	addition.RegisterAdditionServer(grpcServer, newAddServer())

	reflection.Register(grpcServer)

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
