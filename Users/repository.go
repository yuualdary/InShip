package Users

import (
	"InShip/models"
	"fmt"

	"gorm.io/gorm"
)


type Repository interface {
	SaveUser(user models.Users) (models.Users, error)
	FindUserEmail(email string) (models.Users, error)
	FindUserById(ID int)(models.Users,error)
	SaveOTP(otp models.Otps) (models.Otps,error)
	GetUserOtp(UserID int) (models.Otps,error)
	UpdateUser(users models.Users) (models.Users,error)
	UpdateOTP(otp models.Otps) (models.Otps,error)


}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository{

	return &repository{db}
}

func(r *repository) SaveUser(user models.Users) (models.Users, error){

	err:=r.db.Create(&user).Error

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
func(r *repository)FindUserById(ID int)(models.Users,error){

	var User models.Users

	err:= r.db.Where("id = ?", ID).Find(&User).Error

	if err !=nil{
		return User,err
	}

	return User,err

}

// func (r *repository)FindUserWithCompany(ID int) (models.Users,error){

// 	var company models.Users

// 	err := r.db.Preload("company_id = ?",GetUser.ID )
// }




func(r *repository)	SaveOTP(otp models.Otps) (models.Otps,error){

	err:= r.db.Create(&otp).Error

	if err != nil{
		
		return otp,err
	}
	
	return otp,nil
}
func(r *repository)	UpdateOTP(otp models.Otps) (models.Otps,error){

	err:= r.db.Save(&otp).Error

	if err != nil{
		
		return otp,err
	}
	
	return otp,nil
}

func(r *repository)	UpdateUser(users models.Users) (models.Users,error){

	err:= r.db.Save(&users).Error

	if err != nil{
		
		return users,err
	}
	
	return users,nil
}



func (r *repository)GetUserOtp(UserID int) (models.Otps,error){
	
	var Otp models.Otps
	err:= r.db.Where("users_id = ?", UserID).Find(&Otp).Error
	fmt.Println(err)
	if err !=nil{
		return Otp,err
	}

	return Otp,err
}
