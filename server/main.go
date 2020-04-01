// Package main implements a server for Greeter service.
package main

import (
	"github.com/parijatpurohit/sidecar-sql/server/handlers"
	"log"
	"net"

	pb "github.com/parijatpurohit/sidecar-sql/zz_generated/go"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStorageServiceServer(s, &handlers.Service{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}


}
