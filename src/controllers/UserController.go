package controllers

import (
	"fmt"
	"gofiber/src/helpers"
	"gofiber/src/middlewares"
	"gofiber/src/models"
	"math"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUser(c *fiber.Ctx) error {
	keyword := c.Query("search")
	sort := c.Query("sort")
	sortBy := c.Query("orderBy")
	pageOld := c.Query("page")
	limitOld := c.Query("limit")
	page, _ := strconv.Atoi(pageOld)
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(limitOld)
	if limit == 0 {
		limit = 5
	}
	offset := (page - 1) * limit
	if sort == "" {
		sort = "ASC"
	}
	if sortBy == "" {
		sortBy = "name"
	}
	sort = sortBy + " " + strings.ToLower(sort)
	users := models.GetAllUser(sort, keyword, limit, offset)
	fmt.Println("user", users)
	count := helpers.CountData("users")
	totalPage := math.Ceil(float64(count) / float64(limit))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "Successfully retrieved all users",
		"data":      users,
		"count":     count,
		"totalPage": totalPage,
		"limit":     limit,
		"page":      page,
	})
}
func GetDetailUser(c *fiber.Ctx) error {
	claims := middlewares.GetUserClaims(c)
	if claims == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Tidak dapat mengakses",
		})
	}
	userId := claims["ID"]
	foundUser := models.GetDetailUser(userId)
	if foundUser.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    foundUser,
		"message": "Successfully found",
	})
}

func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var input map[string]interface{}
	if err := c.BodyParser(&input); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
		return err
	}
	input = helpers.XssMiddleware(input)
	var updateUser models.User
	mapstructure.Decode(input, &updateUser)

	errors := helpers.ValidateStruct(updateUser)
	if len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}
	user, _ := models.GetUserByEmail(updateUser.Email)
	if input["email"] != user.Email {
		if user.Email != "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": fmt.Sprintf("User with email %v already exist", updateUser.Email),
			})
		}
	}
	if input["password"] != "" {
		err := helpers.ValidatePassword(updateUser.Password)
		if err != nil {
			return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(updateUser.Password), bcrypt.DefaultCost)
		updateUser.Password = string(hashPassword)
	}

	err := models.UpdateUser(id, &updateUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Failed to update user with ID %d", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("User with ID %d updated successfully", id),
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	err := models.DeleteUser(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Failed to delete user with ID %d", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Success to delete user with ID %d", id),
	})
}
