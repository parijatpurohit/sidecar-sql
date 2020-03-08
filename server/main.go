// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/timestamp"

	pb "github.com/sidecar-storage/zz_generated/go"
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
	return &pb.User_FindByRollAndNameResponse{
		Users: []*pb.User{
			{Name: "some name", Roll: 123, CreatedAt: &timestamp.Timestamp{Seconds: 1234567890111, Nanos: 0}},
		},
	}, nil
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
