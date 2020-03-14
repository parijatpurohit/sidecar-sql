package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/parijatpurohit/sidecar-sql/storage/user"
)

func CreateTableUsers() {
	db, err := gorm.Open("mysql", "root:@(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&user.User{})

	// Create
	db.Create(&user.User{Name: "L1212", Roll: 1000})
}

func InsertTableUsers() {
	db, err := gorm.Open("mysql", "root:@(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()


	// Create
	db.Save(&user.User{Name: "L1213", Roll: 3000})
	db.Save(&user.User{Name: "L1213", Roll: 4000})
	db.Save(&user.User{Name: "L1214", Roll: 5000})
	db.Save(&user.User{Name: "L1214", Roll: 6000})
}

func main() {
	// CreateTableUsers()
	// InsertTableUsers()
	// TestFindByRollAndName()
}
