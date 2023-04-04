package routes

import (
	"jwt-gin/api/controllers"
	"jwt-gin/api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ContractRoutes(incomingRoutes *gin.Engine, db *gorm.DB) {

	contractController := &controllers.ContractController{
		ContractRepository: models.NewContractRepository(db),
	}

	incomingRoutes.GET("/contract/:id", contractController.GetContract)
	incomingRoutes.GET("/contract", contractController.GetContracts)
	incomingRoutes.POST("/contract/create", contractController.CreateContracts)
	incomingRoutes.PUT("/contract/update/:id", contractController.UpdateContracts)
	incomingRoutes.DELETE("/contract/delete/:id", contractController.DeleteContracts)
}
