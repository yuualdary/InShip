package Users

import (
	"InShip/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)


type Service interface {
	RegisterUser(input RegisterInput) (models.Users, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterInput) (models.Users, error){

	User := models.Users{}
	User.Name = input.Name
	User.Email = input.Email
	User.Bod = input.BOD
	
	GenereateHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil{

		return User, err
	}

	User.Password = string(GenereateHash)
	User.Role = "candidate"

	CheckMail,err := s.repository.FindUserEmail(User.Email)

	if err !=nil {
		return CheckMail,err

	}

	if CheckMail.ID != 0{

		return CheckMail, errors.New("Email Already Used")
	}

	NewUser, err := s.repository.RegisterUser(User)

	if err != nil {

		return NewUser, err
	}
	return NewUser, nil

}