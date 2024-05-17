package routes

import (
	"gofiber/src/controllers"
	"gofiber/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	api:= app.Group("/v1")
	auth := api.Group("/auth")
	// products
	// api.Get("/products",  controllers.GetAllProducts)
	api.Get("/products", middlewares.JwtMiddleware(),middlewares.ValidateSellerRole(), controllers.GetAllProducts)
	api.Get("/product/:id", controllers.GetDetailProduct)
	api.Post("/product", controllers.CreateProduct)
	api.Put("/product/:id", controllers.UpdateProduct)
	api.Delete("/product/:id", controllers.DeleteProduct)
	// users
	api.Get("/users", controllers.GetAllUser)
	api.Get("/user",middlewares.JwtMiddleware(), controllers.GetDetailUser)
	api.Post("/refreshToken", controllers.RefreshToken)
	api.Put("/user/:id", controllers.UpdateUser)
	api.Delete("/user/:id", controllers.DeleteUser)
	// addresses
	api.Get("/address", controllers.GetAllAddresses)
	api.Get("/address/:id", controllers.GetDetailAddress)
	api.Post("/address", controllers.CreateAddress)
	api.Put("/address/:id", controllers.UpdateAddress)
	api.Delete("/address/:id", controllers.DeleteAddress)
	//categories
	api.Get("/categories", controllers.GetAllCategory)
	api.Get("/category/:id",controllers.GetCategoryById)
	api.Post("/category", controllers.CreateCategory)
	api.Put("/category/:id", controllers.UpdateCategory)
	api.Delete("/category/:id", controllers.DeleteCategory)
	// auth
	auth.Post("/register", controllers.RegisterUser)
	auth.Post("/login", controllers.LoginUser)
	
	}