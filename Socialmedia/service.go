package Socialmedia

import (
	"InShip/Users"
	"InShip/models"
)


type Service interface {
	CreateSocmed(UserID int, input CreateSocialInput) (models.SocialMedias,error)
}

type service struct {
	repository Repository
	UserRepository Users.Repository
}

func NewService(repository Repository,UserRepository Users.Repository) *service {
	return &service{repository, UserRepository}
}


// func(s *service)CreateSocmed(UserID int, input CreateSocialInput) (models.SocialMedias,error){

// 	CreateSocmed := models.SocialMedias{}

	
// }
