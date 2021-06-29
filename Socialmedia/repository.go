package Socialmedia

import (
	"InShip/models"

	"gorm.io/gorm"
)


type Repository interface {
	Save(socmed models.SocialMedias)(models.SocialMedias, error)
	UpdateSocmed(socmed models.SocialMedias)(models.SocialMedias, error)
	FindByUserID(UserID int)(models.SocialMedias, error)
	FindByCompaniesID(CompanyID int) (models.SocialMedias, error)

}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {

	return &repository{db}
}

func (r *repository)Save(socmed models.SocialMedias)(models.SocialMedias, error){

	err := r.db.Create(&socmed).Error

	if err !=nil{

		return socmed, err
	}

	return socmed, nil

}

func (r *repository)UpdateSocmed(socmed models.SocialMedias)(models.SocialMedias, error){

	err := r.db.Save(&socmed).Error

	if err !=nil{

		return socmed, err
	}

	return socmed, nil

}

func (r *repository)FindByUserID(UserID int)(models.SocialMedias, error){

	var socmed models.SocialMedias
	err:= r.db.Where("users_id = ?", UserID).Find(&socmed).Error

	if err != nil {
		return socmed, err
	}

	return socmed,nil
}


func (r *repository)FindByCompaniesID(CompanyID int) (models.SocialMedias, error){

	var socmed models.SocialMedias
	err:= r.db.Where("companies_id = ?", CompanyID).Find(&socmed).Error

	if err != nil {
		return socmed, err
	}

	return socmed,nil
}


