package main

import (
	"InShip/Company"
	"InShip/Users"
	"InShip/auth"
	"InShip/config"
	"InShip/handler"
	"InShip/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main(){

	router := gin.Default()
	router.Use(cors.Default())

	config.ConnectDatabase()


	UserRepository:= Users.NewRepository(config.DB)
	CampaignRepository :=Company.NewRepository(config.DB)

	UsersService := Users.NewService(UserRepository)
	CompanyService := Company.NewService(CampaignRepository,UserRepository)
	AuthService := auth.NewService()



	UsersHandler := handler.NewUserHandler(UsersService,AuthService)
	CompanyHandler := handler.NewCompanyHandler(CompanyService)

	router.Static("/images","./images")

	v1:= router.Group("/api/v1")
	{
		v1.POST("/users/register",UsersHandler.RegisterUser)
		v1.POST("/users/login",UsersHandler.LoginUser)
		v1.POST("/users/uploadavatar",middleware.AuthMiddleware(AuthService,UsersService),UsersHandler.SaveAvatar)
		v1.POST("/users/otpcheck",middleware.AuthMiddleware(AuthService,UsersService),UsersHandler.CheckOtp)
		v1.POST("/users/resendotp",middleware.AuthMiddleware(AuthService,UsersService),UsersHandler.ResendOTP)
		v1.POST("/company",middleware.AuthMiddleware(AuthService,UsersService),CompanyHandler.CreateCompany)
		v1.POST("/company/:id",CompanyHandler.DetailCompany)
		v1.PUT("/company/:id",middleware.AuthMiddleware(AuthService,UsersService),CompanyHandler.UpdateCompany)


	}


	// GetCurrentDate := time.Now().Local()

	// fmt.Println(GetCurrentDate)


	router.Run(":8000")

}
