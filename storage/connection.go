package storage

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"sync"
)

var (
	once = sync.Once{}
	db   *gorm.DB
)

func GetDB() *gorm.DB {
	var err error
	once.Do(func() {
		db, err = gorm.Open("mysql", "root:@(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)
		if err != nil {
			panic(fmt.Errorf("error creating connection. check database configuration and try again, err: %v", err))
		}
	})
	return db
}
