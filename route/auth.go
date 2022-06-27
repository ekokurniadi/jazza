package route

import (
	"jazza/auth"
	"jazza/handler"
	"jazza/repository"
	"jazza/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type authRoute struct {
	db          *gorm.DB
	api         *gin.RouterGroup
	authService auth.Service
}

func NewAuthRoute(db *gorm.DB, api *gin.RouterGroup, authService auth.Service) *authRoute {
	return &authRoute{db, api, authService}
}

func (r *authRoute) AuthRouteMap() {
	authRepository := repository.NewAuthRepository(r.db)
	authServices := service.NewAuthUserService(authRepository)
	authHandler := handler.NewAuthHandler(authServices, r.authService)

	r.api.POST("/auth/login", authHandler.Login)
}
