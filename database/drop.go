package database

import (
	"errors"
	"slices"

	"gorm.io/gorm"
)

func Drop(db *gorm.DB, value []string) (e error) {

	switch value[0] {
	case "all":
		e = dropAll(db)
	default:
		return errors.New("invalid drop value")
	}

	return
}

func dropAll(db *gorm.DB) error {
	err := db.Migrator().DropTable(slices.Concat(base)...)

	if err != nil {
		return err
	}

	return nil
}
