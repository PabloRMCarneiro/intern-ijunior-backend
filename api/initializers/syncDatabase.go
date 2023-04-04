package initializers

import (
	"jwt-gin/api/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
