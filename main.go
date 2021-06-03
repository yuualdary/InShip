package main

import (
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
	UsersService := Users.NewService(UserRepository)
	AuthService := auth.NewService()


	UsersHandler := handler.NewUserHandler(UsersService,AuthService)


	router.Static("/images","./images")

	v1:= router.Group("/api/v1")
	{
		v1.POST("/users/register",UsersHandler.RegisterUser)
		v1.POST("/users/login",UsersHandler.LoginUser)
		v1.POST("/users/uploadavatar",middleware.AuthMiddleware(AuthService,UsersService),UsersHandler.SaveAvatar)
	}

	// GetCurrentDate := time.Now().Local()

	// fmt.Println(GetCurrentDate)


	router.Run(":8000")

}
