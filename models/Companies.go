package models

import "gorm.io/gorm"

type Companies struct {
	gorm.Model
	CompanyName    string
	CompanyAddress string
	CompanyProfile string
	CompanyPhoto   string
	IsCompany      bool


	SocialMedias []SocialMedias
}