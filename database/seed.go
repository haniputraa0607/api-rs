package database

import (
	"api-rs/models"
	"api-rs/utility"
	"errors"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB, value []string) (e error) {

	switch value[0] {
	case "all":
		e = seedAll(db)
	default:
		return errors.New("invalid seed value")
	}

	return
}

func seedAll(db *gorm.DB) error {
	err := createUser(db)
	if err != nil {
		return err
	}

	return nil
}

func createUser(db *gorm.DB) error {
	users := []models.User{
		{
			Username: "admin",
			Password: utility.HashPassword("12345"),
		},
	}

	err := db.Create(&users).Error
	if err != nil {
		return err
	}

	return nil

}
