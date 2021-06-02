package models

type Companies struct {
	ID             int `gorm:"primary_key"`
	CompanyName    string
	CompanyAddress string
	CompanyProfile string
	CompanyPhoto   string
	IsCompany      bool
	UsersID        string
	User           Users `gorm:"foreignKey:UsersID"`

	SocialMedias []SocialMedias
}