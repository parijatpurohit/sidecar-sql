package sqlconn

import (
	"errors"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/parijatpurohit/sidecar-sql/lib/config"
)

var (
	once = sync.Once{}
	db   *gorm.DB
)

func Initialize(conf *config.SQLConfig) (*gorm.DB, error) {
	var err error
	once.Do(func() {
		db, err = gorm.Open(conf.SQL.Dialect, conf.SQL.ConnectionStr)
		if err != nil {
			err = errors.New(fmt.Sprintf("error creating connection. check database configuration and try again, err: %v", err))
			return
		}
		db.DB().SetMaxIdleConns(conf.SQL.MaxIdle)
		db.DB().SetMaxOpenConns(conf.SQL.MaxOpen)

	})
	return db, err
}
func GetDB() *gorm.DB {
	return db
}
