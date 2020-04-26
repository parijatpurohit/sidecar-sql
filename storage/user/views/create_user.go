package user_views

import (
	"github.com/google/uuid"
	"github.com/parijatpurohit/sidecar-sql/storage/user/models"
)

func (v *viewsImpl) CreateUser(user *models.User) (*models.User, error) {
	db := v.db
	if user.UUID == "" {
		user.UUID = uuid.New().String()
	}
	if err := db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
