package user_views

import (
	"context"
	"database/sql"
	"github.com/parijatpurohit/sidecar-sql/storage"
	"github.com/parijatpurohit/sidecar-sql/storage/user/models"
	"github.com/parijatpurohit/sidecar-sql/utils"
	"time"
)

func DeleteUsers(ctx context.Context, users []*models.User) ([]*models.User, error) {
	db := storage.GetDB().BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	for _, user := range users {
		user.UpdatedAt = utils.GetTimePtr(time.Now())
		user.DeletedAt = utils.GetTimePtr(time.Now())
		db.Model(&user).UpdateColumns(user)
	}
	db.Commit()
	return users, nil
}
