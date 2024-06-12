package main

import (
	"gofiber/src/configs"
	"gofiber/src/helpers"
	"gofiber/src/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/joho/godotenv"
)
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return "0.0.0.0:" + port

}

func main() {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	app := fiber.New()

	app.Use(helmet.New())
	// app.Static("/public", "./src/public")
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "*",
	// 	AllowMethods: "GET POST PUT PATCH DELETE",
	// 	AllowHeaders: "*",
	// 	ExposeHeaders: "Content-Length",
	// }))

	configs.InitDB()
	helpers.Migration()
	routes.Router(app)

	app.Listen(getPort())
}
