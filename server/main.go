// Package main implements a server for Greeter service.
package main

import (
	"context"
	"github.com/parijatpurohit/sidecar-sql/storage/user/translators"
	"github.com/parijatpurohit/sidecar-sql/storage/user/views"
	"log"
	"net"

	pb "github.com/parijatpurohit/sidecar-sql/zz_generated/go"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedStorageServiceServer
}

// User_FindByRollAndName implements storage.StorageServiceServer
func (s *server) User_FindByRollAndName(ctx context.Context, in *pb.User_FindByRollAndNameRequest) (*pb.User_FindByRollAndNameResponse, error) {
	users, _ := user_views.FindByRollAndNameRequest(in)
	return &pb.User_FindByRollAndNameResponse{Users: translators.TranslateUsers(users)}, nil
}



func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStorageServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
