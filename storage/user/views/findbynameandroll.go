package user_views

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	storage "github.com/parijatpurohit/sidecar-sql/storage"
	"github.com/parijatpurohit/sidecar-sql/storage/user"
	pb "github.com/parijatpurohit/sidecar-sql/zz_generated/go"
)



// request, response needs to be generated from config
// this method can be converted to a template
func FindByRollAndNameRequest(request *pb.User_FindByRollAndNameRequest) ([]*user.User, error){
	//TODO : this connection should be created on startup and singleton pattern should be used
	db := storage.GetDB()

	var u1 user.User
	rows, err := db.Where("roll=? AND name IN (?)", request.Query.Roll, request.Query.Name).Find(&u1).Rows()
	if err != nil {
		return nil, err
	}

	var res  []*user.User
	for rows.Next() {
		var tempUser user.User
		err := db.ScanRows(rows, &tempUser)
		if err!= nil {
			fmt.Println(err)
		}
		fmt.Println(tempUser)
		res = append(res, &tempUser)
	}
	err = rows.Close()
	if err != nil {
		fmt.Println(err)
	}

	return res, nil
}

