// Package main implements a server for Greeter service.
package server

import (
	"log"
	"net"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	"github.com/parijatpurohit/sidecar-sql/lib/sqlconn"
	pb "github.com/parijatpurohit/sidecar-sql/zz_generated/go/protogen"
	"github.com/parijatpurohit/sidecar-sql/zz_generated/go/server/handlers"
	job_views "github.com/parijatpurohit/sidecar-sql/zz_generated/go/storage/job/views"
	jobapplication_views "github.com/parijatpurohit/sidecar-sql/zz_generated/go/storage/jobapplication/views"
	user_views "github.com/parijatpurohit/sidecar-sql/zz_generated/go/storage/user/views"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func Serve() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	initializeDB()

	userViews := user_views.GetViews()
	jobViews := job_views.GetViews()
	jobapplicationViews := jobapplication_views.GetViews()
	service := &handlers.Service{
		UserViews:           userViews,
		JobViews:            jobViews,
		JobApplicationViews: jobapplicationViews,
	}
	pb.RegisterStorageServiceServer(s, service)
	log.Printf("starting server on port: %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initializeDB() {
	sqlConfig := config.GetSQLConfig()
	_, err := sqlconn.Initialize(sqlConfig)
	if err != nil {
		log.Fatalf("error initialising storage: %v", err)
	}
}
