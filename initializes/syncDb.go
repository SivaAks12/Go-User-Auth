package initializes

import (
	"example.com/go-jwt/models"
)

func SyncDataBase() {
	DB.AutoMigrate(&models.User{})
}
