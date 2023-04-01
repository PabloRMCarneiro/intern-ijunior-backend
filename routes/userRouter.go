package routes

import (
	"jwt-gin/controllers"
	"jwt-gin/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(incomingRoutes *gin.Engine, db *gorm.DB) {

	userController := &controllers.UserController{
		UserRepository: models.NewUserRepository(db),
	}

	incomingRoutes.POST("/users/signup", userController.Signup)
	incomingRoutes.POST("/users/login", userController.Login)
	incomingRoutes.POST("/users/logout", userController.Logout)

	incomingRoutes.GET("/users", userController.GetUsers)
	incomingRoutes.GET("/users/:id", userController.GetUser)
	incomingRoutes.GET("/users/loggedin", userController.GetLoggedInUser)

	incomingRoutes.PUT("/users/update/:id", userController.UpdateUser)

	incomingRoutes.DELETE("/users/delete/:id", userController.DeleteUser)
}
