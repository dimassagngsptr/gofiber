package routes

import (
	"gofiber/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	// products
	app.Get("/product", controllers.GetAllProducts)
	app.Get("/product/:id", controllers.GetDetailProduct)
	app.Post("/product", controllers.CreateProduct)
	app.Put("/product/:id", controllers.UpdateProduct)
	app.Delete("/product/:id", controllers.DeleteProduct)
	// users
	app.Get("/users", controllers.GetAllUser)
	app.Get("/user/:id", controllers.GetDetailUser)
	app.Post("/user", controllers.CreateUser)
	app.Put("/user/:id", controllers.UpdateUser)
	app.Delete("/user/:id", controllers.DeleteUser)
	// addresses
	app.Get("/address", controllers.GetAllAddresses)
	app.Get("/address/:id", controllers.GetDetailAddress)
	app.Post("/address", controllers.CreateAddress)
	app.Put("/address/:id", controllers.UpdateAddress)
	app.Delete("/address/:id", controllers.DeleteAddress)
	//categories
	app.Get("/category", controllers.GetAllCategory)
	app.Get("/category/:id",controllers.GetCategoryById)
	app.Post("/category", controllers.CreateCategory)
	app.Put("/category/:id", controllers.UpdateCategory)
	app.Delete("/category/:id", controllers.DeleteCategory)
	// auth
	app.Post("/auth/register", controllers.RegisterUser)
	app.Post("/auth/login", controllers.LoginUser)

	}