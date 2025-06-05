package models

import "gorm.io/gorm"

type Analytic struct{
	gorm.Model
	ShortenedUrl string `json:"shortened_url" gorm:"type:varchar(255);not null"`
	IpAddress string `json:"ip_address" gorm:"type:varchar(255);not null"`
	Country string `json:"country" gorm:"type:varchar(255);not null"`
	UserAgent string `json:"user_agent" gorm:"type:varchar(255);not null"`
}