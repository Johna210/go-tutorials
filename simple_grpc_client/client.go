package main

import (
	"context"
	"log"
	mainpb "simplegrpcclient/proto/gen"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Did not connect:", err)
	}
	defer conn.Close()

	client := mainpb.NewCalculateClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &mainpb.AddRequest{
		A: 10,
		B: 20,
	}
	res, err := client.Add(ctx, req)
	if err != nil {
		log.Fatalln("Could not add", err)
	}

	log.Println("Sum:", res.Result)
}
