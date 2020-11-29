package logic

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/parijatpurohit/sidecar-sql/vaccinepass/server/transport"
	pb "github.com/parijatpurohit/sidecar-sql/zz_generated/go/protogen"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50010"
)

type ILogic interface {
	GetUserDetails(req *transport.GetUserDetailsRequest) (*transport.GetUserDetailsResponse, error)
	GetUserVaccineDetails(req *transport.GetUserVaccineDetailsRequest) (*transport.GetUserVaccineDetailsResponse, error)
	GetRequiredVaccines(req *transport.GetRequiredVaccineRequest) (*transport.GetRequiredVaccineResponse, error)
}

var (
	logicClient *logicImpl
	once        sync.Once
)

type logicImpl struct {
	address string
	client  pb.StorageServiceClient
}

func initialise() {
	logicClient = &logicImpl{address: "localhost:50010"}
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
	conn, err := grpc.DialContext(ctx, logicClient.address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	logicClient.client = pb.NewStorageServiceClient(conn)
}

func GetClient() ILogic {
	once.Do(func() {
		initialise()
	})
	return logicClient
}
