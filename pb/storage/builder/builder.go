package builder

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/parijatpurohit/sidecar-sql/generated/storage/user"
)

func FindByRollAndNameRequest(request *user.FindByRollAndNameRequest) (*user.FindByRollAndNameResponse, error){
	//TODO : this connection should be created on startup and singelton
	db, err := gorm.Open("mysql", "root:@(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()

	var users *user.User
	rows, err := db.Find(&users, "roll=? AND name IN (?)", request.Query.Roll, request.Query.Name).Rows()
	if err != nil {
		return nil, err
	}


	for rows.Next() {
		var tempUser user.User
		err := db.ScanRows(rows, &tempUser)
		if err!= nil {
			fmt.Println(err)
		}
		fmt.Println(tempUser)
	}

	return nil, nil
}
