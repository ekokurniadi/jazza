package handler

import (
	"jazza/auth"
	"jazza/formatter"
	"jazza/helper"
	"jazza/input"
	"jazza/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	service     service.AuthUserService
	authService auth.Service
}

func NewAuthHandler(service service.AuthUserService, authService auth.Service) *authHandler {
	return &authHandler{service, authService}
}

func (h *authHandler) Login(c *gin.Context) {
	var loginInput input.AuthLoginInput
	err := c.ShouldBindJSON(&loginInput)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response :=
			helper.ApiResponse("Login Gagal, silahkan coba kembali", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loggedUser, err := h.service.FindUserByPhoneNumber(loginInput)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response :=
			helper.ApiResponse("Login Gagal, silahkan coba kembali", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedUser.ID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response :=
			helper.ApiResponse("Login Gagal, silahkan coba kembali", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := formatter.FormatAuth(loggedUser, token)
	response :=
		helper.ApiResponse("Login success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
