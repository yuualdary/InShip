package Company

import (
	"InShip/models"

	"gorm.io/gorm"
)


type Repository interface {

	SaveCompanies(companies models.Companies)(models.Companies,error)
	UpdateCompanies(companies models.Companies)(models.Companies,error)
	GetDetailCompany(CompanyID int)(models.Companies,error)

}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {

	return &repository{db}
}




func (r *repository) SaveCompanies(companies models.Companies)(models.Companies,error){

	err:= r.db.Create(&companies).Error

	if err != nil{

		return companies,err
	}

	return companies, nil
}
func (r *repository) UpdateCompanies(companies models.Companies)(models.Companies,error){


	err:= r.db.Save(&companies).Error

	if err != nil{

		return companies,err
	}

	return companies, nil
}

func(r *repository) GetDetailCompany(CompanyID int)(models.Companies,error){

	var DetailCompany models.Companies
	
	err:= r.db.Where("id = ? ", CompanyID).Find(&DetailCompany).Error
 
	if err != nil{
		return DetailCompany,err
	}

	return DetailCompany,nil
}

