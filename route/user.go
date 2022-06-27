package route

import (
	"jazza/handler"
	"jazza/repository"
	"jazza/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userRoute struct {
	db  *gorm.DB
	api *gin.RouterGroup
}

func NewUserRoute(db *gorm.DB, api *gin.RouterGroup) *userRoute {
	return &userRoute{db, api}
}

func (r *userRoute) UserRouteMap() {
	userRepository := repository.NewUserRepository(r.db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	r.api.POST("/users", userHandler.CreateUser)
	r.api.GET("/users", userHandler.GetUsers)
	r.api.GET("/users/:id", userHandler.GetUser)
	r.api.PUT("/users/:id", userHandler.UpdateUser)
	r.api.DELETE("/users/:id", userHandler.DeleteUser)
}
