package handlers

import (
	"context"
	"github.com/parijatpurohit/sidecar-sql/storage/user/translators"
	user_views "github.com/parijatpurohit/sidecar-sql/storage/user/views"
	pb "github.com/parijatpurohit/sidecar-sql/zz_generated/go"
)

// User_FindByRollAndName implements storage.StorageServiceServer
func (s *Service) User_FindByRollAndName(ctx context.Context, in *pb.User_FindByRollAndNameRequest) (*pb.User_FindByRollAndNameResponse, error) {
	query := translators.TranslateUser_FindByRollAndNameQuery(in.Query)
	res, err := user_views.FindByRollAndName(query)
	if err != nil {
		return nil, err
	}
	return translators.TranslateUser_FindByRollAndNameResponse(res), nil
}

func (s *Service) User_CreateUser(ctx context.Context, in *pb.User_CreateUserRequest) (*pb.User_CreateUserResponse, error) {
	req := translators.TranslateUser_CreateUser(in.User)
	res, err := user_views.CreateUser(req)
	if err != nil {
		return nil, err
	}
	return translators.TranslateUser_CreateUserResponse(res), nil
}