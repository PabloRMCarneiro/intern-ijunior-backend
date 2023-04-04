package main

import (
	"jwt-gin/api/initializers"
	"jwt-gin/api/middleware"
	"jwt-gin/api/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEndVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router, initializers.DB)
	routes.ContractRoutes(router, initializers.DB)
	routes.ProjectRoutes(router, initializers.DB)
	router.Use(middleware.RequireAuth())

	// create a table "x" in the database

	//initializers.DB.AutoMigrate(&models.Project{})

	router.Run()
}
