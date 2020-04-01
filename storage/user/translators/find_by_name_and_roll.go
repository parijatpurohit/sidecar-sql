package translators

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/parijatpurohit/sidecar-sql/storage/user/models"
	"github.com/parijatpurohit/sidecar-sql/utils"
	pb "github.com/parijatpurohit/sidecar-sql/zz_generated/go"
)

func TranslateUser_FindByRollAndNameQuery(in *pb.User_FindByRollAndNameQuery) *models.User_FindByRollAndNameQuery {
	res := models.User_FindByRollAndNameQuery{
		Roll:      in.Roll,
		DeletedAt: utils.GetTimestamp(in.DeletedAt),
		UpdatedAt: utils.GetTimestamp(in.UpdatedAt),
		Name:      in.Name,
	}
	return &res
}

func TranslateUser_FindByRollAndNameResponse(in []*models.User) *pb.User_FindByRollAndNameResponse {
	return &pb.User_FindByRollAndNameResponse{Users: TranslateUsers(in)}
}
func TranslateUsers(in []*models.User) []*pb.User {
	var returnVal []*pb.User
	for _, elem := range in {
		returnVal = append(returnVal, TranslateUser(elem))
	}
	return returnVal
}

func TranslateUser(user *models.User) *pb.User {
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
		Name:                 user.Name,
		Roll:                 user.Roll,
		CreatedAt:            createdAt,
		UpdatedAt:            updatedAt,
		DeletedAt:            deletedAt,
	}
}

