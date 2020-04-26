package translators

import (
	"github.com/parijatpurohit/sidecar-sql/storage/user/models"
	pb "github.com/parijatpurohit/sidecar-sql/zz_generated/go/protogen"
)

func TranslateUser_DeleteUsersRequest(in *pb.User_DeleteUsersRequest) []*models.User {
	users := in.Users
	var res []*models.User
	for _, user := range users {
		res = append(res, TranslateUser_Proto(user))
	}
	return res
}

func TranslateUser_DeleteUsersResponse(users []*models.User) *pb.User_DeleteUsersResponse {
	return &pb.User_DeleteUsersResponse{Count: int64(len(users))}
}
