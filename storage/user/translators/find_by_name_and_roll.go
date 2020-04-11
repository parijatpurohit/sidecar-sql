package translators

import (
	"github.com/parijatpurohit/sidecar-sql/storage/user/models"
	pb "github.com/parijatpurohit/sidecar-sql/zz_generated/go"
)

func TranslateUser_FindByRollAndNameRequest(in *pb.User_FindByRollAndNameRequest) *models.User_FindByRollAndNameQuery {
	query := in.Query
	res := models.User_FindByRollAndNameQuery{
		Roll: query.Roll,
		Name: query.Name,
	}
	return &res
}

func TranslateUser_FindByRollAndNameResponse(in []*models.User) *pb.User_FindByRollAndNameResponse {
	return &pb.User_FindByRollAndNameResponse{Users: TranslateUsers(in)}
}
func TranslateUsers(in []*models.User) []*pb.User {
	var returnVal []*pb.User
	for _, elem := range in {
		returnVal = append(returnVal, TranslateUserModel(elem))
	}
	return returnVal
}
