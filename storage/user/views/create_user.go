package user_views

import (
	"github.com/google/uuid"
	"github.com/parijatpurohit/sidecar-sql/storage"
	"github.com/parijatpurohit/sidecar-sql/storage/user/models"
)

func CreateUser(user *models.User) (*models.User, error){
	db := storage.GetDB()
	user.UUID = uuid.New().String()
	if err := db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
