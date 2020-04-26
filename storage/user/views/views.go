package user_views

import (
	"context"
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/parijatpurohit/sidecar-sql/lib/sqlconn"
	"github.com/parijatpurohit/sidecar-sql/storage/user/models"
)

var (
	once  = sync.Once{}
	views IViews
)

type IViews interface {
	CreateUser(user *models.User) (*models.User, error)
	DeleteUsers(ctx context.Context, users []*models.User) ([]*models.User, error)
	FindByRollAndName(query *models.User_FindByRollAndNameQuery) ([]*models.User, error)
	UpdateUsers(ctx context.Context, users []*models.User) ([]*models.User, error)
}
type viewsImpl struct {
	db *gorm.DB
}

func GetViews() IViews {
	once.Do(func() {
		db := sqlconn.GetDB()
		views = &viewsImpl{db: db}
	})
	return views
}
