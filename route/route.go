package route

import (
	"jazza/auth"

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

	router := gin.Default()
	api := router.Group("/api/v1")

	/// INFO : USER ENDPOINT [CREATE, READ, UPDATE, DELETE]
	NewUserRoute(a.db, api).UserRouteMap()

	/// INFO : AUTH ENDPOINT [LOGIN, REGISTER]
	NewAuthRoute(a.db, api, middleWareService).AuthRouteMap()

	router.Run()
}
