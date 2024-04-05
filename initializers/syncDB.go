package initializers

import "github.com/satyam-jha-16/jwt-auth/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}
