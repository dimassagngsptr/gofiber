package controllers

import (
	"fmt"
	_ "gofiber/src/helpers"
	"gofiber/src/models"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mitchellh/mapstructure"
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
	sort := c.Query("sort")
	sortBy := c.Query("orderBy")
	if sort == "" {
		sort = "ASC"
	}
	if sortBy == "" {
		sortBy ="name"
	}
	sort = sortBy + " " + strings.ToLower(sort)
	foundUser := models.GetUserById(sort,id)
	if foundUser.ID == 0{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message":"User not found",
		})
	}
	return c.JSON(foundUser)
}

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