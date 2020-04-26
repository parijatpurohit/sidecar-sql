package translators

import (
	"github.com/parijatpurohit/sidecar-sql/storage/user/models"
	pb "github.com/parijatpurohit/sidecar-sql/zz_generated/go/protogen"
)

func TranslateUser_CreateUserRequest(in *pb.User_CreateUserRequest) *models.User {
	return TranslateUser_Proto(in.User)
}

func TranslateUser_CreateUserResponse(user *models.User) *pb.User_CreateUserResponse {
	return &pb.User_CreateUserResponse{UUID: user.UUID}
}
