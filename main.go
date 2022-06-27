package main

import (
	"jazza/config"
	"jazza/route"
	"log"
)

func main() {
	database, err := config.NewAppConfig().InitialDataBase()
	if err != nil {
		log.Fatal(err)
		return
	}
	appRoute := route.NewAppRouteConfig(database)
	appRoute.DeclareAppRoute()

}
