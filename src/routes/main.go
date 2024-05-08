package routes

import (
	"gofiber/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	// products
	app.Get("/products", controllers.GetAllProducts)
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
	app.Get("/address", controllers.GetAllUserAddresses)
	app.Get("/address/:id", controllers.GetUserAddress)
	app.Post("/address", controllers.CreateUserAddress)
	app.Put("/address/:id", controllers.UpdateUserAddress)
	app.Delete("/address/:id", controllers.DeleteUserAddress)
	//categories
	app.Get("/category", controllers.GetAllCategories)
	app.Get("/category/:id",controllers.GetCategoryById)
	app.Post("/category", controllers.CreateCategory)
	app.Put("/category/:id", controllers.UpdateCategory)
	app.Delete("/category/:id", controllers.DeleteCategory)
	}