package translators

import (
	"github.com/parijatpurohit/sidecar-sql/storage/user/models"
	pb "github.com/parijatpurohit/sidecar-sql/zz_generated/go/protogen"
)

func TranslateUser_UpdateUsersRequest(in *pb.User_UpdateUsersRequest) []*models.User {
	users := in.Users
	var res []*models.User
	for _, entity := range users {
		res = append(res, TranslateUser_Proto(entity))
	}
	return res
}

func TranslateUser_UpdateUsersResponse(users []*models.User) *pb.User_UpdateUsersResponse {
	var entities []*pb.User
	for _, user := range users {
		entities = append(entities, TranslateUserModel(user))
	}
	return &pb.User_UpdateUsersResponse{Users: entities}
}
