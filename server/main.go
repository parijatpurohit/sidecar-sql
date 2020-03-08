/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/timestamp"

	pb "github.com/sidecar-storage/generated"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedStorageServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) User_FindByRollAndName(ctx context.Context, in *pb.User_FindByRollAndNameRequest) (*pb.User_FindByRollAndNameResponse, error) {
	return &pb.User_FindByRollAndNameResponse{
		Users: []*pb.User{
			{
				Name:      "some name",
				Roll:      123,
				CreatedAt: &timestamp.Timestamp{Seconds: 1234567890111, Nanos: 0},
			},
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
