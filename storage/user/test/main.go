package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/parijatpurohit/sidecar-sql/storage/user/models"
)

func CreateTableUsers() {
	db, err := gorm.Open("mysql", "root:@(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.User{})
}

func main() {
	CreateTableUsers()
	// InsertTableUsers()
	// TestFindByRollAndName()
}
