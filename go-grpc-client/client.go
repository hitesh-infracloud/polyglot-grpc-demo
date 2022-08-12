package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "polyglot-grpc-demo/addition"
)

func main() {
	conn, err := grpc.Dial("localhost:10000", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
