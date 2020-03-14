package translators

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/parijatpurohit/sidecar-sql/storage/user"
	storage "github.com/parijatpurohit/sidecar-sql/zz_generated/go"
)

func TranslateUsers(inputs []*user.User) []*storage.User {
	var returnVal []*storage.User
	for _, elem := range inputs {
		returnVal = append(returnVal, TranslateUser(elem))
	}
	return returnVal
}

func TranslateUser(user *user.User) *storage.User {
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

	return &storage.User{
		Name:                 user.Name,
		Roll:                 user.Roll,
		CreatedAt:            createdAt,
		UpdatedAt:            updatedAt,
		DeletedAt:            deletedAt,
	}
}
