// Package main implements a client for Greeter service.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/parijatpurohit/sidecar-sql/zz_generated/go/protogen"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50010"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewStorageServiceClient(conn)

	name := "random_"
	createRes, err := c.User_CreateUser(context.Background(), &pb.User_CreateUserRequest{User: &pb.User{UUID: "12345", Name: name, Roll: 1234}})
	if err != nil {
		fmt.Println("err here\n--------", err)
	}

	fmt.Println(createRes)

	//findRes, err := c.User_FindByRollAndName(context.Background(), &pb.User_FindByRollAndNameRequest{
	//	Query: &pb.User_FindByRollAndName_Query{Name: []string{name, "blah"}, Roll: 1234},
	//})
	//if err != nil {
	//	fmt.Println("err here\n--------", err)
	//}
	user := &pb.User{
		UUID: "12345",
		Name: "UpdateName2",
	}
	updateRes, err := c.User_UpdateUsers(context.Background(), &pb.User_UpdateUsersRequest{Users: []*pb.User{user}})
	if err != nil {
		fmt.Println("err here\n--------", err)
	}

	fmt.Println(updateRes)

	deleteRes, err := c.User_DeleteUsers(context.Background(), &pb.User_DeleteUsersRequest{Users: []*pb.User{user}})
	if err != nil {
		fmt.Println("err here\n--------", err)
	}

	fmt.Println(deleteRes)

}
