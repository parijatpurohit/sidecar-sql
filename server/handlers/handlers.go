package handlers

import (
	pb "github.com/parijatpurohit/sidecar-sql/zz_generated/go/protogen"
	Job_views "github.com/parijatpurohit/sidecar-sql/zz_generated/go/storage/job/views"
	User_Views "github.com/parijatpurohit/sidecar-sql/zz_generated/go/storage/user/views"
)

type Service struct {
	pb.UnimplementedStorageServiceServer
	UserViews User_Views.IUserViews
	JobViews  Job_views.IJobViews
}
