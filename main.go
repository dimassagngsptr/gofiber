package main

import (
	"gofiber/src/configs"
	"gofiber/src/helper"
	"gofiber/src/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	app:=fiber.New()

	configs.InitDB()
	helper.Migration()
	routes.Router(app)
	

	app.Listen(":3000")
}