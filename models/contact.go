package models

import "time"

type Contact struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null;" json:"name"`
	Value     string    `gorm:"type:varchar(255);not null" json:"value"`
	Icon      *string   `gorm:"type:varchar(255)" json:"-"`
	CreatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP()" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP()" json:"updated_at"`
}