package handlers

import (
	"context"
	"log"

	"github.com/parijatpurohit/sidecar-sql/storage/user/translators"
	user_views "github.com/parijatpurohit/sidecar-sql/storage/user/views"
	pb "github.com/parijatpurohit/sidecar-sql/zz_generated/go"
)

// User_FindByRollAndName implements storage.StorageServiceServer
func (s *Service) User_FindByRollAndName(ctx context.Context, in *pb.User_FindByRollAndNameRequest) (*pb.User_FindByRollAndNameResponse, error) {
	log.Println("User_FindByRollAndName")
	query := translators.TranslateUser_FindByRollAndNameRequest(in)
	res, err := user_views.FindByRollAndName(query)
	if err != nil {
		return nil, err
	}
	return translators.TranslateUser_FindByRollAndNameResponse(res), nil
}

func (s *Service) User_CreateUser(ctx context.Context, in *pb.User_CreateUserRequest) (*pb.User_CreateUserResponse, error) {
	log.Println("User_CreateUser")
	req := translators.TranslateUser_CreateUserRequest(in)
	res, err := user_views.CreateUser(req)
	if err != nil {
		return nil, err
	}
	return translators.TranslateUser_CreateUserResponse(res), nil
}

func (s *Service) User_UpdateUsers(ctx context.Context, in *pb.User_UpdateUsersRequest) (*pb.User_UpdateUsersResponse, error) {
	log.Println("User_UpdateUsers")
	req := translators.TranslateUser_UpdateUsersRequest(in)
	res, err := user_views.UpdateUsers(ctx, req)
	if err != nil {
		return nil, err
	}
	return translators.TranslateUser_UpdateUsersResponse(res), nil
}

func (s *Service) User_DeleteUsers(ctx context.Context, in *pb.User_DeleteUsersRequest) (*pb.User_DeleteUsersResponse, error) {
	log.Println("User_DeleteUsers")
	req := translators.TranslateUser_DeleteUsersRequest(in)
	res, err := user_views.DeleteUsers(ctx, req)
	if err != nil {
		return nil, err
	}
	return translators.TranslateUser_DeleteUsersResponse(res), nil
}
