package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fahimsGit/gRpctest/proto"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client not started %v", err)
	}
	defer conn.Close()

	client := proto.NewArticleServiceClient(conn)

	/*createReq := proto.CreateArticleReq{
		Name:          "Names",
		PublisherName: "Publisher",
	}*/

	getReq := proto.GetArticleReq{
		Id: "f47009c2-9bb4-4da1-a246-8353619c2c6b",
	}
	/*	createResp, err := client.CreateArticle(context.Background(), &createReq)

		if err != nil {
			log.Fatalf("Server response error %v", err)
		}

		fmt.Printf("Response from Server ", createResp)*/

	getResp, err := client.GetArticle(context.Background(), &getReq)
	if err != nil {
		log.Fatalf("Server response error %v", err)
	}

	fmt.Println("Response from Server ", getResp)
}
