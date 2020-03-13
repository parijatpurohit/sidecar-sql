package views

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/parijatpurohit/sidecar-sql/generated/storage/dto"
	"github.com/parijatpurohit/sidecar-sql/implementation2/user"

)



// request, response needs to be generated from config
// this method can be templatised
func FindByRollAndNameRequest(request *dto.FindByRollAndNameRequest) (*dto.FindByRollAndNameResponse, error){
	//TODO : this connection should be created on startup and singleton pattern should be used
	db, err := gorm.Open("mysql", "root:@(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()



	var u1 user.Users


	exp := db.Where("roll=? AND name IN (?)", request.Query.Roll, request.Query.Name).Find(&u1).Debug().QueryExpr()
	fmt.Println(exp)
	rows, err := db.Where("roll=? AND name IN (?)", request.Query.Roll, request.Query.Name).Find(&u1).Rows()
	if err != nil {
		return nil, err
	}


	for rows.Next() {
		var tempUser user.Users
		err := db.ScanRows(rows, &tempUser)
		if err!= nil {
			fmt.Println(err)
		}
		fmt.Println("user", tempUser)
	}

	return nil, nil
}

