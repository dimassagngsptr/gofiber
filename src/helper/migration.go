package helper

import (
	"gofiber/src/configs"
	"gofiber/src/models"
)

func Migration() {
	configs.DB.AutoMigrate(&models.Product{})
	configs.DB.AutoMigrate(&models.Users{})
	configs.DB.AutoMigrate(&models.UserAddress{})
	configs.DB.AutoMigrate(&models.Categories{})
}