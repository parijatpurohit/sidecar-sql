package translators

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/parijatpurohit/sidecar-sql/storage/user/models"
	"github.com/parijatpurohit/sidecar-sql/utils"
	pb "github.com/parijatpurohit/sidecar-sql/zz_generated/go/protogen"
)

func TranslateUserModel(user *models.User) *pb.User {
	var createdAt *timestamp.Timestamp
	var updatedAt *timestamp.Timestamp
	var deletedAt *timestamp.Timestamp
	if user.CreatedAt != nil {
		createdAt, _ = ptypes.TimestampProto(*user.CreatedAt)
	}
	if user.UpdatedAt != nil {
		updatedAt, _ = ptypes.TimestampProto(*user.UpdatedAt)
	}
	if user.DeletedAt != nil {
		deletedAt, _ = ptypes.TimestampProto(*user.DeletedAt)
	}
	return &pb.User{
		UUID:      user.UUID,
		Name:      user.Name,
		Roll:      user.Roll,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}

func TranslateUser_Proto(user *pb.User) *models.User {
	return &models.User{
		UUID:      user.UUID,
		Name:      user.Name,
		Roll:      user.Roll,
		CreatedAt: utils.GetTimestamp(user.CreatedAt),
		UpdatedAt: utils.GetTimestamp(user.UpdatedAt),
		DeletedAt: utils.GetTimestamp(user.DeletedAt),
	}
}
