package controllers

import (
	"fmt"
	"gofiber/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllUserAddresses(c *fiber.Ctx) error {
	address := models.GetAllUserAddress()
	return c.JSON(address)
}

func GetUserAddress(c *fiber.Ctx) error{
	id, _ := strconv.Atoi(c.Params("id"))
	foundUserAddress := models.GetUserAddress(id)
	if foundUserAddress.ID == 0{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message":"Address not found",
		})
	}
	return c.JSON(foundUserAddress)
}

func CreateUserAddress(c *fiber.Ctx) error {
	var newUserAddress models.UserAddress
	if err := c.BodyParser(&newUserAddress); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"invalid request body",
		})
		return err
	}
	addressExist := models.GetUserAddressByNameAndAddress(newUserAddress.Name, newUserAddress.Address)
	if addressExist.Address != "" || addressExist.Name != ""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Address already exists",
		})
	}
	models.PostUserAddress(&newUserAddress)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":"success created new address",
	})
}

func UpdateUserAddress(c *fiber.Ctx) error{
	var newUserAddress models.UserAddress
	id, _ := strconv.Atoi(c.Params("id"))
	if err := c.BodyParser(&newUserAddress); err != nil{
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Invalid request body",
		})
		return err
	}
	
	err:=models.UpdateUserAddress(id, &newUserAddress)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Failed to update user address with ID %d", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Successfully updated user address with ID %d",id),
	})
}

func DeleteUserAddress(c *fiber.Ctx) error{
	id, _ := strconv.Atoi(c.Params("id"))
	err := models.DeleteUserAddress(id)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Failed to delete user address with ID %d", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Successfully deleted user address with ID %d", id),
	})
}