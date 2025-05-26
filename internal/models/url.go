package models

import (
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	Original  string `json:"original" gorm:"type:varchar(255);not null"`
	Shortened string `json:"shortened" gorm:"type:varchar(255);not null;uniqueIndex"`
	UserEmail    string   `json:"user_email" gorm:"type:varchar(255);not null"`
	
	User   User `json:"user" gorm:"foreignKey:UserEmail;references:Email;onDelete:CASCADE"`
}




