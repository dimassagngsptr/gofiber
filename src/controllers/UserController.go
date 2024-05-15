package controllers

import (
	"fmt"
	_"gofiber/src/helpers"
	"gofiber/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_"github.com/mitchellh/mapstructure"
)

func GetAllUser(c *fiber.Ctx) error {
	users, count := models.GetAllUser()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Successfully retrieved all users",
		"data":users,
		"count":count,
	})
}
func GetDetailUser(c *fiber.Ctx) error{
	id, _ := strconv.Atoi(c.Params("id"))
	foundUser := models.GetUserById(id)
	if foundUser.ID == 0{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message":"User not found",
		})
	}
	return c.JSON(foundUser)
}

// func CreateUser(c *fiber.Ctx) error {
// 	var user map[string]interface{}
// 	if err := c.BodyParser(&user); err != nil {
// 		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Invalid request body",
// 		})
// 		return err
// 	}


// 	user = helpers.XssMiddleware(user)
// 	var newUser models.User
// 	mapstructure.Decode(user, &newUser)


// 	errors := helpers.ValidateStruct(newUser)
// 	if len(errors) > 0{
// 		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
// 	}
	
// 	userExist, _ := models.GetUserByEmail(newUser.Email)
// 	if userExist.Email != "" {
// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message":fmt.Sprintf("User with email %v already exist", newUser.Email),
// 		})
// 	}

// 	models.PostUser(&newUser)
// 	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
// 		"message": "User created successfully",
// 	})
// }

func UpdateUser(c *fiber.Ctx) error{
	id, _ := strconv.Atoi(c.Params("id"))
	var updateUser models.User
	if err := c.BodyParser(&updateUser); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Invalud request body",
		})
		return err
	}
	err := models.UpdateUser(id, &updateUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Failed to update user with ID %d", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":fmt.Sprintf("User with ID %d updated successfully", id),
	})
}

func DeleteUser(c *fiber.Ctx) error{
	id, _ := strconv.Atoi(c.Params("id"))
	err := models.DeleteUser(id)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":fmt.Sprintf("Failed to delete user with ID %d", id),
		})
	}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":fmt.Sprintf("Success to delete user with ID %d", id),
		})
}