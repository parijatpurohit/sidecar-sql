// Package main implements a client for Greeter service.
package main

import (
	"context"
	"fmt"
	pb "github.com/parijatpurohit/sidecar-sql/zz_generated/go"
	"google.golang.org/grpc"
	"log"
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

	name  := "random_"
	createRes, err := c.User_CreateUser(context.Background(), &pb.User_CreateUserRequest{User: &pb.User{Name: name, Roll: 1234}})
	if err != nil {
		fmt.Println("err here\n--------", err)
	}

	fmt.Println(createRes)

	findRes, err := c.User_FindByRollAndName(context.Background(), &pb.User_FindByRollAndNameRequest{
		Query: &pb.User_FindByRollAndNameQuery{Name: []string{name, "blah"}, Roll: 1234},
	})
	if err != nil {
		fmt.Println("err here\n--------", err)
	}

	fmt.Println("find response", findRes)

	for _, res := range findRes.Users {
		res.Name = res.Name + "_updated"
	}
	updateRes, err := c.User_UpdateUsers(context.Background(), &pb.User_UpdateUsersRequest{Entities: findRes.Users})
	if err != nil {
		fmt.Println("err here\n--------", err)
	}

	fmt.Println(updateRes)

	deleteRes, err := c.User_DeleteUsers(context.Background(), &pb.User_DeleteUsersRequest{Entities: updateRes.Entities})
	if err != nil {
		fmt.Println("err here\n--------", err)
	}

	fmt.Println(deleteRes)

}
