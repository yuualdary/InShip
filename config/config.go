package config

import (
	"InShip/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var DB *gorm.DB

func ConnectDatabase(){

	var(
	
		
		Users = models.Users{}
		Companies = models.Companies{}
		Otps = models.Otps{}
		SocialMedias = models.SocialMedias{}

	)
	

	dsn := "root:@tcp(127.0.0.1:3306)/inship?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: false,})//set true kalau mau dimatiin

	if err != nil{
		log.Fatal(err.Error())
	}

	

	db.AutoMigrate(&Users,&Companies,&Otps,&SocialMedias)

	fmt.Println("Connecting To Database...")

	DB = db



}