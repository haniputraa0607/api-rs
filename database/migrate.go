package database

import (
	"errors"
	"slices"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB, value []string) (e error) {

	switch value[0] {
	case "all":
		e = migrateAll(db)
	default:
		return errors.New("invalid migration value")
	}

	return
}

func migrateAll(db *gorm.DB) error {
	err := db.AutoMigrate(slices.Concat(base)...)

	if err != nil {
		return err
	}

	return nil
}
