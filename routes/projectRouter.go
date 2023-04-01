package routes

import (
	"jwt-gin/controllers"
	"jwt-gin/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProjectRoutes(incomingRoutes *gin.Engine, db *gorm.DB) {

	projectController := &controllers.ProjectController{
		ProjectRepository: models.NewProjectRepository(db),
	}

	incomingRoutes.GET("/project/:id", projectController.GetProject)
	incomingRoutes.GET("/project/", projectController.GetProjects)
	incomingRoutes.POST("/project/create", projectController.CreateProject)
	incomingRoutes.PUT("/project/update/:id", projectController.UpdateProject)
	incomingRoutes.PUT("/project/update/:id/:userId", projectController.UpdateUserInProject)
	incomingRoutes.DELETE("/project/delete/:id", projectController.DeleteProject)
	incomingRoutes.PUT("/project/delete/:id", projectController.RemoveUserInProject)
}
