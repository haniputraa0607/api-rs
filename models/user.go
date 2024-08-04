package models

import "time"

type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"type:varchar(255);not null;index;unique" json:"username"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	CreatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP()" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP()" json:"updated_at"`
}
