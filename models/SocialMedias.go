package models

type SocialMedias struct {
	ID          int `gorm:"primary_key"`
	SocType     string
	SocLink     string
	UsersID     string
	CompaniesID string
	User        Users     `gorm:"foreignKey:UsersID"`
	Company     Companies `gorm:"foreignKey:CompaniesID"`
}