package userrepository

import (
	"api-rs/models"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() (users []models.User, err error)
	GetUser(username string) (users *models.User, err error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUsers() (users []models.User, err error) {
	var db = r.db.Model(&users)

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) GetUser(username string) (users *models.User, err error) {
	var user models.User
	err = r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}

		return nil, err
	}

	return &user, nil
}
