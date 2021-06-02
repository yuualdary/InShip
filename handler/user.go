package handler

import (
	"InShip/Users"
	"InShip/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type UserHandler struct {
	UserService Users.Service
}

func NewUserHandler(UserService Users.Service) *UserHandler{
	return &UserHandler{UserService}
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

