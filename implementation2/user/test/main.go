package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/parijatpurohit/sidecar-sql/generated/storage/dto"
	"github.com/parijatpurohit/sidecar-sql/implementation2/user"
	"github.com/parijatpurohit/sidecar-sql/implementation2/user/views"
)

func CreateTableUsers() {
	db, err := gorm.Open("mysql", "root:@(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&user.Users{})

	// Create
	db.Create(&user.Users{Name: "L1212", Roll: 1000})
}

func InsertTableUsers() {
	db, err := gorm.Open("mysql", "root:@(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()


	// Create
	db.Save(&user.Users{Name: "L1213", Roll: 3000})
	db.Save(&user.Users{Name: "L1213", Roll: 4000})
	db.Save(&user.Users{Name: "L1214", Roll: 5000})
	db.Save(&user.Users{Name: "L1214", Roll: 6000})
}
func TestFindByRollAndName() {
	_, _ = views.FindByRollAndNameRequest(&dto.User_FindByRollAndNameRequest{
		Query:&dto.User_Query{
			Roll: 3000,
			Name: []string{"L1213"},
		},
	})
}

func main() {
	TestFindByRollAndName()
}
