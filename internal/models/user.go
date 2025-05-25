package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `json:"email" gorm:"type:varchar(255);not null;uniqueIndex"`
	Name      string `json:"name" gorm:"type:varchar(255);not null"`
	Picture   string `json:"picture" gorm:"type:varchar(255);default:'https://example.com/default.png'"`
}

type GoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}