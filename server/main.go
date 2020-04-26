// Package main implements a server for Greeter service.
package main

import (
	"fmt"
	"log"
	"net"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	"github.com/parijatpurohit/sidecar-sql/server/handlers"
	user_views "github.com/parijatpurohit/sidecar-sql/storage/user/views"
	pb "github.com/parijatpurohit/sidecar-sql/zz_generated/go/protogen"
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
	sqlConfig := config.GetSQLConfig()
	userViews := user_views.GetViews(sqlConfig)
	service := &handlers.Service{UserViews: userViews}
	pb.RegisterStorageServiceServer(s, service)
	fmt.Println("starting server on port", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
