package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	pb "polyglot-grpc-demo/addition"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// load certificate of CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("certs/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, errors.New("failed to add server CA's certificate")
	}

	// create credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("unable to load tls credentials", err)
	}

	conn, err := grpc.Dial("localhost:10000", grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatal("unable to connect to server localhost:10000")
	}
	client := pb.NewAdditionClient(conn)

	response, err := client.Add(context.Background(), &pb.AddRequest{Number: 10, AnotherNumber: 20})
	if err != nil {
		log.Fatal("error while invoking grpc add service", err)
	}

	log.Println("add response from server", response.Sum)
}
