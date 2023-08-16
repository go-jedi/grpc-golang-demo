package main

import (
	"log"
	"net"

	"github.com/rob-bender/go-grpc-test/pb"
	"google.golang.org/grpc"
)

const (
	port string = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port) // подключаемя по tcp на порт 8080
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
	// create a new gRPC server
	grpcServer := grpc.NewServer()
	// register the greet service
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("Server started at %v", lis.Addr())
	//list is the port, the grpc server needs to start there
	if err := grpcServer.Serve(lis); err != nil { // проверка соединения
		log.Fatalf("Failed to start: %v", err)
	}
}
