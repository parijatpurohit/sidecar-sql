package user_views

import (
	"context"
	"sync"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
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
	sqlConfig *config.SQLConfig
}

func GetViews(sqlConfig *config.SQLConfig) IViews {
	once.Do(func() {
		views = &viewsImpl{sqlConfig: sqlConfig}
	})
	return views
}
