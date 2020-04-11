package translators

import (
	"github.com/parijatpurohit/sidecar-sql/storage/user/models"
	pb "github.com/parijatpurohit/sidecar-sql/zz_generated/go"
)

func TranslateUser_DeleteUsersRequest(in *pb.User_DeleteUsersRequest) []*models.User {
	entities := in.Entities
	var res []*models.User
	for _, entity := range entities {
		res = append(res, TranslateUser_Proto(entity))
	}
	return res
}

func TranslateUser_DeleteUsersResponse(users []*models.User) *pb.User_DeleteUsersResponse {
	return &pb.User_DeleteUsersResponse{Count: int64(len(users))}
}
