package handlers

import (
	user_views "github.com/parijatpurohit/sidecar-sql/storage/user/views"
	pb "github.com/parijatpurohit/sidecar-sql/zz_generated/go/protogen"
)

type Service struct {
	pb.UnimplementedStorageServiceServer
	UserViews user_views.IViews
}
