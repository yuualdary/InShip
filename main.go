package main

import (
	"InShip/Users"
	"InShip/config"
	"InShip/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main(){

	router := gin.Default()
	router.Use(cors.Default())

	config.ConnectDatabase()


	UserRepository:= Users.NewRepository(config.DB)
	UsersService := Users.NewService(UserRepository)
	


	UsersHandler := handler.NewUserHandler(UsersService)


	router.Static("/images","./images")

	v1:= router.Group("/api/v1")
	{
		v1.POST("/users/register",UsersHandler.RegisterUser)

	}

	// GetCurrentDate := time.Now().Local()

	// fmt.Println(GetCurrentDate)


	router.Run(":8000")

}
