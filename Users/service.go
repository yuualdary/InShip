package Users

import (
	"InShip/models"
	"errors"
	"math/rand"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)


type Service interface {
	RegisterUser(input RegisterInput) (models.Users, error)
	SaveAvatar(ID int, filelocatiion string) (models.Users,error)
	LoginUser(input LoginInput)(models.Users, error)
	GetUserById(ID int) (models.Users, error)
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
	
	GetSplit := strings.Replace(User.Name," ","",-1)
	GetChar := GetSplit[0:3]
	GetRandNum:= rand.Intn(10000)
	User.Initial = GetChar+ strconv.Itoa(GetRandNum)

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

	NewUser, err := s.repository.SaveUser(User)

	if err != nil {

		return NewUser, err
	}
	return NewUser, nil

}

func(s *service)LoginUser(input LoginInput)(models.Users, error){

	email := input.Email
	password := input.Password

	//get data user

	GetDataUser, err := s.repository.FindUserEmail(email)

	if err !=nil{

		return GetDataUser,err
	}

	if GetDataUser.ID == 0{

		return GetDataUser, errors.New("Email Not Found")
	}

	//hash password

	err = bcrypt.CompareHashAndPassword([]byte(GetDataUser.Password),[]byte(password))

	if err != nil{

		return GetDataUser, err

	}

	return GetDataUser,nil
}
func(s *service)SaveAvatar(ID int, filelocatiion string) (models.Users,error){

	
	//check id

	CheckUserID, err := s.repository.FindUserById(ID)

	if err != nil{

		return CheckUserID,err
	}

	CheckUserID.Profile_photo = filelocatiion

	Update,err := s.repository.SaveUser(CheckUserID)

	if err !=nil{
		return Update,err
	}

	return Update, nil
	

}
func (s *service)GetUserById(ID int) (models.Users, error){

	user, err := s.repository.FindUserById(ID)

	if err != nil{
		return user, err
	}

	if user.ID == 0{
		return user, errors.New("No user found")
	}
	return user, nil
}