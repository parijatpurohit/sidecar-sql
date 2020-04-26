package sqlconn

import (
	"errors"
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
)

var (
	once = sync.Once{}
	db   *gorm.DB
)

func Initialize(conf *config.SQLConfig) (*gorm.DB, error) {
	var err error
	once.Do(func() {
		db, err = gorm.Open(conf.Dialect, conf.ConnectionStr)
		db.DB().SetMaxIdleConns(conf.MaxIdle)
		db.DB().SetMaxOpenConns(conf.MaxOpen)
		if err != nil {
			err = errors.New(fmt.Sprintf("error creating connection. check database configuration and try again, err: %v", errStr))
		}
	})
	return db, err
}
func GetDB() *gorm.DB {
	return db
}
