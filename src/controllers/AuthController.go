package controllers

import (
	"fmt"
	"gofiber/src/helpers"
	"gofiber/src/models"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *fiber.Ctx) error{
	var input map[string]interface{}
	if err := c.BodyParser(&input); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Failed to parse request body",
		})
	}
	input = helpers.XssMiddleware(input)
	var newUser models.User
	mapstructure.Decode(input, &newUser)


	errors := helpers.ValidateStruct(newUser)
	if len(errors) > 0{
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}
	user, _:= models.GetUserByEmail(newUser.Email)
	if user.Email != ""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":fmt.Sprintf("User with email %v already exist", newUser.Email),
		})
	}
	_, err := helpers.ValidatePassword(newUser.Password)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"message":err.Error(),
		})
	}
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password),bcrypt.DefaultCost)
	newUser.Password = string(hashPassword)
	if err := models.PostUser(&newUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Failed to create new user",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":"Register successfully",
	})
}

func LoginUser(c *fiber.Ctx) error{
	var input models.User

	if err := c.BodyParser(&input); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode":fiber.StatusBadRequest,
			"message":"Failed to parse request body",
		})	
	}
	user, err := models.GetUserByEmail(input.Email)
	if err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"statusCode":fiber.StatusNotFound,
			"message":"Email unregistered",
		})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(input.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"statusCode":fiber.StatusUnauthorized,
			"message":"Incorrect password",
		})
	}
	token, err := helpers.GenerateToken(os.Getenv("JWT_KEY"), user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusCode":fiber.StatusInternalServerError,
			"message":"Failed to generate Token",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusCode":fiber.StatusOK,
		"message":"Login successful",
		"token":token,
		"data":user,
	})
}