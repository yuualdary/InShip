package models

import "time"

type Otps struct {
	ID             int `gorm:"primary_key"`
	Value   string
	UsersID string 
	User Users `gorm:"foreignKey:UsersID"`
	Expired time.Time
}


