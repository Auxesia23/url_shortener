package models

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	Original  string `json:"original"`
	Shortened string `json:"shortened"`
	UserID    uint   `json:"user_id"`
	
	User   User `json:"user" gorm:"foreignKey:UserID;references:ID;onDelete:CASCADE"`
}
