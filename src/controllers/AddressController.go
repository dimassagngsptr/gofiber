package controllers

import (
	"fmt"
	"gofiber/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllAddresses(c *fiber.Ctx) error {
	address, count := models.GetAllAddress()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully retrieved all user address",
		"data":address,
		"count":count,
	})
}

func GetDetailAddress(c *fiber.Ctx) error{
	id, _ := strconv.Atoi(c.Params("id"))
	foundAddress := models.GetAddress(id)
	if foundAddress.ID == 0{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message":"Address not found",
		})
	}
	return c.JSON(foundAddress)
}

func CreateAddress(c *fiber.Ctx) error {
	var newAddress models.Address
	if err := c.BodyParser(&newAddress); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"invalid request body",
		})
		return err
	}
	addressExist := models.GetAddressByNameAndAddress(newAddress.Name, newAddress.Address)
	if addressExist.Address != "" || addressExist.Name != ""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Address already exists",
		})
	}
	models.PostAddress(&newAddress)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":"success created new address",
	})
}

func UpdateAddress(c *fiber.Ctx) error{
	var newAddress models.Address
	id, _ := strconv.Atoi(c.Params("id"))
	if err := c.BodyParser(&newAddress); err != nil{
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Invalid request body",
		})
		return err
	}
	
	err:=models.UpdateAddress(id, &newAddress)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Failed to update user address with ID %d", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Successfully updated user address with ID %d",id),
	})
}

func DeleteAddress(c *fiber.Ctx) error{
	id, _ := strconv.Atoi(c.Params("id"))
	err := models.DeleteAddress(id)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Failed to delete user address with ID %d", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Successfully deleted user address with ID %d", id),
	})
}