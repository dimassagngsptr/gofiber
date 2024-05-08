package controllers

import (
	"fmt"
	"gofiber/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllUser(c *fiber.Ctx) error {
	users := models.GetAllUser()
	return c.JSON(users)
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

func CreateUser(c *fiber.Ctx) error {
	var newUser models.Users
	if err := c.BodyParser(&newUser); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
		return err
	}
	isExist := models.GetUserByEmail(newUser.Email)
	if isExist.Email != "" {
		err:=c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"User email already exist",
		})
		return err
	}

	models.PostUser(&newUser)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}

func UpdateUser(c *fiber.Ctx) error{
	id, _ := strconv.Atoi(c.Params("id"))
	var updateUser models.Users
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