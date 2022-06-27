package route

import (
	"jazza/auth"
	"jazza/handler"
	"jazza/repository"
	"jazza/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type appRouteConfig struct {
	db *gorm.DB
}

func NewAppRouteConfig(db *gorm.DB) *appRouteConfig {
	return &appRouteConfig{db}
}

func (a *appRouteConfig) DeclareAppRoute() {

	middleWareService := auth.NewService()

	userRepository := repository.NewUserRepository(a.db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	authRepository := repository.NewAuthRepository(a.db)
	authServices := service.NewAuthUserService(authRepository)
	authHandler := handler.NewAuthHandler(authServices, middleWareService)

	router := gin.Default()
	api := router.Group("/api/v1")

	/// INFO : USER ENDPOINT [CREATE, READ, UPDATE, DELETE]
	api.POST("/users", userHandler.CreateUser)
	api.GET("/users", userHandler.GetUsers)
	api.GET("/users/:id", userHandler.GetUser)
	api.PUT("/users", userHandler.UpdateUser)
	api.DELETE("/users", userHandler.DeleteUser)

	/// INFO : AUTH ENDPOINT [LOGIN, REGISTER]
	api.POST("/auth/login", authHandler.Login)

	router.Run()
}
