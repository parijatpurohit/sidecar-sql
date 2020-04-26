package sqlconn

import (
	"fmt"
	"sync"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"

	"github.com/jinzhu/gorm"
)

var (
	once = sync.Once{}
	db   *gorm.DB
)

func GetDB(conf *config.SQLConfig) *gorm.DB {
	var err error
	once.Do(func() {
		db, err = gorm.Open(conf.Dialect, conf.ConnectionStr)
		db.DB().SetMaxIdleConns(conf.MaxIdle)
		db.DB().SetMaxOpenConns(conf.MaxOpen)
		if err != nil {
			panic(fmt.Errorf("error creating connection. check database configuration and try again, err: %v", err))
		}
	})
	return db
}
