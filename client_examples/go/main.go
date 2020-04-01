// Package main implements a client for Greeter service.
package main

import (
	"context"
	"fmt"
	pb "github.com/parijatpurohit/sidecar-sql/zz_generated/go"
	"log"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewStorageServiceClient(conn)

	res, err := c.User_FindByRollAndName(context.Background(), &pb.User_FindByRollAndNameRequest{
		Query: &pb.User_Query{Name: []string{"L1213"}, Roll: 3000},
	})
	if err != nil {
		fmt.Println("err here\n--------", err)
	}
	fmt.Println("I REACHED HERE")
	fmt.Println(res)
}
