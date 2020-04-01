package translators

import (
	"github.com/parijatpurohit/sidecar-sql/storage/user/models"
	"github.com/parijatpurohit/sidecar-sql/utils"
	pb "github.com/parijatpurohit/sidecar-sql/zz_generated/go"
)

func TranslateUser_CreateUser(user *pb.User) *models.User {
	return &models.User{
		Name:      user.Name,
		Roll:      user.Roll,
		CreatedAt: utils.GetTimestamp(user.CreatedAt),
		UpdatedAt: utils.GetTimestamp(user.UpdatedAt),
		DeletedAt: utils.GetTimestamp(user.DeletedAt),
	}
}

func TranslateUser_CreateUserResponse(user *models.User) *pb.User_CreateUserResponse {
	return &pb.User_CreateUserResponse{UUID: user.UUID}
}
