package database

import (
	"api-rs/models"
)

var base []interface{} = []interface{}{
	&models.User{},
	&models.Contact{},
	&models.Partner{},
}
