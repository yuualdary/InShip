package models

import "time"

type Users struct {
	ID            int `gorm:"primary_key"`
	Name          string
	Email         string
	Bod           time.Time
	Initial       string
	Password      string
	IsVerif       bool
	Profile_photo string
	Role          string
}