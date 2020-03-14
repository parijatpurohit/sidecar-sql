package user_views

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/parijatpurohit/sidecar-sql/storage/user"
	storage "github.com/parijatpurohit/sidecar-sql/zz_generated/go"
)



// request, response needs to be generated from config
// this method can be templatized
func FindByRollAndNameRequest(request *storage.User_FindByRollAndNameRequest) ([]*user.User, error){
	//TODO : this connection should be created on startup and singleton pattern should be used
	db, err := gorm.Open("mysql", "root:@(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()

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
		res = append(res, &tempUser)
	}

	return res, nil
}

