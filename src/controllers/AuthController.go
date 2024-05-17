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
	err := helpers.ValidatePassword(newUser.Password)
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
	jwtKey := os.Getenv("JWT_KEY")
	payload := map[string]interface{}{
        "ID": user.ID,
				"role":user.Role,
    }
	token, err := helpers.GenerateToken(jwtKey, payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusCode":fiber.StatusInternalServerError,
			"message":"Failed to generate Token",
		})
	}
	refreshToken, err := helpers.GenerateRefreshToken(jwtKey, payload)
	if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Could not generate refresh token",
        })
    }
		item := map[string]string{
        "Email":        user.Email,
        "Token":        token,
        "RefreshToken": refreshToken,
    }
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusCode":fiber.StatusOK,
		"message":"Login successful",
		"data":item,
	})
}
func RefreshToken(c *fiber.Ctx) error {
    var input struct {
        RefreshToken string `json:"refreshToken"`
    }
    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Failed to parse request body",
        })
    }

    jwtKey := os.Getenv("JWT_KEY")
    token, err := helpers.GenerateToken(jwtKey, map[string]interface{}{"refreshToken": input.RefreshToken})
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Could not generate access token",
        })
    }

    refreshToken, err := helpers.GenerateRefreshToken(jwtKey, map[string]interface{}{"refreshToken": input.RefreshToken})
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Could not generate refresh token",
        })
    }
    item := map[string]string{
        "Token":        token,
        "RefreshToken": refreshToken,
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "Refresh successfully",
        "data":    item,
    })
}