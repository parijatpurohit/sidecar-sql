package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/parijatpurohit/sidecar-sql/generated/storage/user"
	"github.com/parijatpurohit/sidecar-sql/pb/storage/builder"
	"time"
)

type Users struct {
	Name string
	Roll int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

//func main() {
//	db, err := gorm.Open("mysql", "root:@(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
//	if err != nil {
//		fmt.Println(err)
//		panic("failed to connect database")
//	}
//	defer db.Close()
//
//	// Migrate the schema
//	db.AutoMigrate(&Users{})
//
//	// Create
//	//db.Create(&Users{Name: "L1212", Roll: 1000})
//
//	// Read
//	var users Users
//	////db.First(&users, 1) // find product with id 1
//	res1 := db.First(&users, "name = ?", "L1212") // find product with code l1212
//	fmt.Println(res1.Value)
//	//// Update - update product's price to 2000
//	//db.Model(&users).Update("Roll", 2000)
//
//	rows, err := db.Find(&users, "name = ?", "L1212").Rows()
//	if err != nil {
//		log.Printf("error in getting sql rows, err %+v", err)
//	}
//	defer rows.Close()
//	for rows.Next() {
//		var user Users
//		db.ScanRows(rows, &user)
//		fmt.Println(user)
//	}
//	//// Delete - delete product
//	//db.Delete(&users)
//}


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
func TestFindByRollAndName() {
	_, _ = builder.FindByRollAndNameRequest(&user.FindByRollAndNameRequest{
		Query:&user.Query{
			Roll: 2000,
			Name: []string{"L1212"},
		},
	})
}
func main() {
	CreateTableUsers()
}