package handler

import (
	"InShip/Company"
	"InShip/helper"
	"InShip/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	CompanyService Company.Service
}

func NewCompanyHandler(CompanyService Company.Service)*CompanyHandler{
	return &CompanyHandler{CompanyService }
}

func(h *CompanyHandler)CreateCompany(c *gin.Context){

	var input Company.CompanyInput
	
	err := c.ShouldBindJSON(&input)

	if err != nil{
		errors := helper.FormatValidationError(err)
		
		ErrorMessage := gin.H{
			"errors" : errors,
		}
		response := helper.APIResponse("Fail Get Data ", http.StatusBadRequest,"errors",ErrorMessage)
		// fmt.Println(response)

		c.JSON(http.StatusUnprocessableEntity,response)
		return
	}
	CurrentUser :=c.MustGet("CurrentUser").(models.Users)//get current user
	//input.User = CurrentUser
	NewCompany, err := h.CompanyService.CreateCompany(int(CurrentUser.ID),input)

	
	if err != nil{
		// fmt.Println(err)
		
		ErrorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Fail Save Data", http.StatusBadRequest,"errors",ErrorMessage)
		fmt.Println(response)

		c.JSON(http.StatusBadRequest,response)
		return
	}

	formatter := Company.DetailCompanyFunc(NewCompany)
	response := helper.APIResponse("Detail Company Data", http.StatusOK, "success",formatter)
	c.JSON(http.StatusOK,response)

}


func(h *CompanyHandler)UpdateCompany(c *gin.Context){

	var CompanyID Company.CompanyDetailInput
	

	err:= c.ShouldBindUri(&CompanyID)

	if err != nil{
		response := helper.APIResponse("Fail Get Bind Data", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input Company.CompanyInput
	err = c.ShouldBindJSON(&input)

	if err != nil{		
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error" : errors,
		}
		response := helper.APIResponse("Fail Get Form Company Data", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	CurrentUser :=c.MustGet("CurrentUser").(models.Users)//get current user
	//input.User = CurrentUser
	NewCompany, err := h.CompanyService.UpdateCompany(int(CurrentUser.ID),input,CompanyID)

	
	if err != nil{
		// fmt.Println(err)
		
		ErrorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Fail Save Data", http.StatusBadRequest,"errors",ErrorMessage)
		fmt.Println(response)

		c.JSON(http.StatusBadRequest,response)
		return
	}

	formatter := Company.DetailCompanyFunc(NewCompany)
	response := helper.APIResponse("Detail Company Data", http.StatusOK, "success",formatter)
	c.JSON(http.StatusOK,response)

}
func(h *CompanyHandler)DetailCompany(c *gin.Context){

	var input Company.CompanyDetailInput

	err:= c.ShouldBindUri(&input)

	if err != nil{
		response := helper.APIResponse("Fail Get Bind Data", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	GetDetail,err := h.CompanyService.DetailCompany(input.ID)

	if err != nil{
		response := helper.APIResponse("Fail Get Detail Company Data", http.StatusNotFound, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := Company.DetailCompanyFunc(GetDetail)
	response := helper.APIResponse("Detail Company Data", http.StatusOK, "success",formatter)
	c.JSON(http.StatusBadRequest, response)
}