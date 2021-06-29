package Company

import (
	"InShip/Users"
	"InShip/models"
	"errors"
)


type Service interface {
	CreateCompany(UserID int, input CompanyInput) (models.Companies, error)
	UpdateCompany(UserID int,input CompanyInput, CompanyID CompanyDetailInput)(models.Companies, error)
	DetailCompany(CompanyID int)(models.Companies, error)
	CheckOwner(UserID int, CompanyID CompanyDetailInput)(models.Companies, error)
}

type service struct{
	repository Repository
	UserRepository Users.Repository
}

func NewService(repository Repository, UserRepository Users.Repository) *service{
	return &service{repository, UserRepository}
}

func (s *service)CreateCompany(UserID int, input CompanyInput) (models.Companies, error){

	CheckUser, err := s.UserRepository.FindUserById(UserID)
	if err != nil {

		return CheckUser.Companies, err
	}
	if CheckUser.CompanyID != 0{

		return CheckUser.Companies,errors.New("you already have company")
	}

	companies := models.Companies{}
	companies.CompanyName = input.CompanyName
	companies.CompanyAddress = input.CompanyAddress
	companies.CompanyProfile = input.CompanyProfile

	SaveCompany, err := s.repository.SaveCompanies(companies)

	if err != nil{
		return SaveCompany,err
	}

	CheckUser.CompanyID = int(SaveCompany.ID)

	UpdUser,err := s.UserRepository.UpdateUser(CheckUser)

	if err != nil{
		return UpdUser.Companies,err
	}
	return SaveCompany,nil

}

func (s *service)UpdateCompany(UserID int,input CompanyInput, CompanyID CompanyDetailInput)(models.Companies, error){


	companies,err := s.CheckOwner(UserID,CompanyID)
	
	if err != nil {

		return companies, err
	}

	
	companies.CompanyName = input.CompanyName
	companies.CompanyAddress = input.CompanyAddress
	companies.CompanyProfile = input.CompanyProfile

	SaveCompany, err := s.repository.UpdateCompanies(companies)

	if err != nil{
		return SaveCompany,err
	}

	
	return SaveCompany,nil

}

func (s *service)CheckOwner(UserID int, CompanyID CompanyDetailInput)(models.Companies, error){

	CheckUser,err := s.UserRepository.FindUserById(UserID)

	if err!=nil{
		return CheckUser.Companies,err
	}

	CheckCompany,err:=s.repository.GetDetailCompany(CompanyID.ID)
	
	if err!=nil{
		return CheckCompany,err
	}

	if uint(CheckUser.CompanyID) != CheckCompany.ID{

		return CheckCompany,errors.New("you are not the owner")

	}

	return CheckCompany,nil

}

func (s *service)DetailCompany(CompanyID int)(models.Companies, error){

	DetailCompany,err := s.repository.GetDetailCompany(CompanyID)

	if err !=nil{
		return DetailCompany,err
	}
	return DetailCompany,nil
}

