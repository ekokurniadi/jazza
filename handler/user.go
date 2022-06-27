package handler

import (
	"net/http"
	"strconv"

	"jazza/formatter"
	"jazza/helper"
	"jazza/input"
	"jazza/service"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *userHandler {
	return &userHandler{service}
}
func (h *userHandler) GetUser(c *gin.Context) {
	var input input.InputIDUser
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get User", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	userDetail, err := h.service.UserServiceGetByID(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get User", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Detail User", http.StatusOK, "success", formatter.FormatUser(userDetail))
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) GetUsers(c *gin.Context) {
	users, err := h.service.UserServiceGetAll()
	if err != nil {
		response := helper.ApiResponse("Failed to get Users", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("List of Users", http.StatusOK, "success", formatter.FormatUsers(users))
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var input input.UserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Create User failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newUser, err := h.service.UserServiceCreate(input)
	if err != nil {
		response := helper.ApiResponse("Create User failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Create User", http.StatusOK, "success", formatter.FormatUser(newUser))
	c.JSON(http.StatusOK, response)
}
func (h *userHandler) UpdateUser(c *gin.Context) {
	var inputID input.InputIDUser
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Users", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var inputData input.UserInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Update User failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	updatedUser, err := h.service.UserServiceUpdate(inputID, inputData)
	if err != nil {
		response := helper.ApiResponse("Failed to get Users", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Update User", http.StatusOK, "success", formatter.FormatUser(updatedUser))
	c.JSON(http.StatusOK, response)
}
func (h *userHandler) DeleteUser(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	var inputID input.InputIDUser
	inputID.ID = id
	_, err := h.service.UserServiceGetByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Users", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.service.UserServiceDeleteByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Users", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Delete User", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
