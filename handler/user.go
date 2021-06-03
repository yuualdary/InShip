package handler

import (
	"InShip/Users"
	"InShip/auth"
	"InShip/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type UserHandler struct {
	UserService Users.Service
	AuthService auth.Service
}

func NewUserHandler(UserService Users.Service, AuthService auth.Service) *UserHandler{
	return &UserHandler{UserService,AuthService}
}


func(h *UserHandler)RegisterUser(c *gin.Context){

	var input Users.RegisterInput

	err := c.ShouldBindJSON(&input)
	//form validation

	// DateValidator:=helper.DateValidator(input.BOD) 

	// if DateValidator != ""{

	// 	response := helper.APIResponse("Fail Register Data", http.StatusBadRequest,"errors",DateValidator)
	// 	c.JSON(http.StatusUnprocessableEntity,response)
	// 	return
	// }

	if err != nil{

		errors := helper.FormatValidationError(err)
		
		ErrorMessage := gin.H{
			"errors" : errors,
		}
		response := helper.APIResponse("Fail Register Data", http.StatusBadRequest,"errors",ErrorMessage)
		// fmt.Println(response)

		c.JSON(http.StatusUnprocessableEntity,response)
		return
	}

	
	//

	NewUser, err := h.UserService.RegisterUser(input)

	if err != nil{
		// fmt.Println(err)
		
		ErrorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Fail Register Data", http.StatusBadRequest,"errors",ErrorMessage)
		fmt.Println(response)

		c.JSON(http.StatusBadRequest,response)
		return
	}


	response := helper.APIResponse("Account Has Been Registered", http.StatusOK,"success",NewUser)
	c.JSON(http.StatusOK,response)
}

func (h *UserHandler)SaveAvatar(c *gin.Context){

}

func (h *UserHandler)LoginUser(c *gin.Context){

	var input Users.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil{
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error" : errors,
		}
		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	NewLogin,err := h.UserService.LoginUser(input)

	if err != nil{
		ErrorMessage := gin.H{
			"error" : err.Error(),
		}	
		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.AuthService.GenerateToken(NewLogin.ID)

	if err != nil {
		ErrorMessage := gin.H{
			"error" : err.Error(),
		}	
		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := Users.FormatUser(NewLogin,token)
	response := helper.APIResponse("success Login", http.StatusOK,"success",formatter)
	c.JSON(http.StatusOK, response)

}

