package user_views

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	storage "github.com/parijatpurohit/sidecar-sql/storage"
	"github.com/parijatpurohit/sidecar-sql/storage/user/models"
)



// request, response needs to be generated from config
// this method can be converted to a template
func FindByRollAndName(query *models.User_FindByRollAndNameQuery) ([]*models.User, error){
	db := storage.GetDB()

	var u1 models.User
	rows, err := db.Where("roll=? AND name IN (?)", query.Roll, query.Name).Find(&u1).Rows()
	if err != nil {
		return nil, err
	}

	var res  []*models.User
	for rows.Next() {
		var tempUser models.User
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

