package main

import (
	"log"
	"net"

	"github.com/fahimsGit/gRpctest/handler"
	pb "github.com/fahimsGit/gRpctest/proto"
	"google.golang.org/grpc"
)

type Server struct {
}

func main() {

	listen, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("Server starting failed %v", err)
	}

	grpcServer := grpc.NewServer()
	serv := handler.NewServer()

	pb.RegisterArticleServiceServer(grpcServer, &serv)

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("Grpc server listen failed %v", err)
	}

}
