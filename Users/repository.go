package Users

import (
	"InShip/models"

	"gorm.io/gorm"
)


type Repository interface {
	RegisterUser(user models.Users) (models.Users, error)
	FindUserEmail(email string) (models.Users, error)

}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository{

	return &repository{db}
}

func(r *repository) RegisterUser(user models.Users) (models.Users, error){

	err:=r.db.Save(&user).Error

	if err !=nil{

		return user,err
	}

	return user,nil


}

func(r *repository)FindUserEmail(email string) (models.Users, error){

	var User models.Users
	err:= r.db.Where("email = ?",email).Find(&User).Error

	if err != nil {
		return User, err
	}

	return User,nil

}
