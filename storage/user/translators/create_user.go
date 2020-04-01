package translators

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/parijatpurohit/sidecar-sql/storage/user/models"
	pb "github.com/parijatpurohit/sidecar-sql/zz_generated/go"
	"time"
)

func TranslateUser_CreateUser(user *pb.User) *models.User {
	var createdAt time.Time
	var updatedAt time.Time
	var deletedAt time.Time
	if user.CreatedAt != nil {
		createdAt, _ = ptypes.Timestamp(user.CreatedAt)
	}
	if user.UpdatedAt != nil {
		updatedAt, _ = ptypes.Timestamp(user.UpdatedAt)
	}
	if user.DeletedAt != nil {
		deletedAt, _ = ptypes.Timestamp(user.DeletedAt)
	}
	return &models.User{
		Name:                 user.Name,
		Roll:                 user.Roll,
		CreatedAt:            &createdAt,
		UpdatedAt:            &updatedAt,
		DeletedAt:            &deletedAt,
	}
}

func TranslateUser_CreateUserResponse(user *models.User) *pb.User_CreateUserResponse{
	return &pb.User_CreateUserResponse{UUID: user.UUID}
}

