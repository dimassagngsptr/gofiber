package controllers

import (
	"gofiber/src/helpers"
	"gofiber/src/models"
	"math"
	"strconv"
	"strings"

	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetAllProducts(c *fiber.Ctx) error {
	keyword := c.Query("search")
	pageOld := c.Query("page")
	limitOld := c.Query("limit")
	page,_ := strconv.Atoi(pageOld)
	limit,_ :=strconv.Atoi(limitOld)
	sort := c.Query("sort")
	sortBy := c.Query("orderBy")
	if page == 0{
		page = 1
	}
	if limit == 0{
		limit = 5
	}
	offset := (page - 1) *limit
	if sort == "" {
		sort = "ASC"
	}
	if sortBy == "" {
		sortBy = "name"
	}
	sort = sortBy + " " + strings.ToLower(sort)
	products := models.SelectAllProduct(sort,keyword,limit,offset)
	count := helpers.CountData("products")
	totalPage := math.Ceil(float64(count) / float64(limit))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Successfully retrieved all products",
		"data":products,
		"totalData":count,
		"totalPage":totalPage,
		"limit":limit,
		"page":page,
	})
}

func GetDetailProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	foundProduct := models.SelectProductById(id)
	if foundProduct == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}
	return c.JSON(foundProduct)
}

func CreateProduct(c *fiber.Ctx) error {
	var newProduct models.Product
	if err := c.BodyParser(&newProduct); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
		return err
	}
	models.PostProduct(&newProduct)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var updatedProduct models.Product
	if err := c.BodyParser(&updatedProduct); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
		return err
	}

	err := models.UpdateProduct(id, &updatedProduct)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Failed to update product with ID %d", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Product with ID %d updated successfully", id),
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	err := models.DeleteProduct(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Failed to delete product with ID %d", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Product with ID %d deleted successfully", id),
	})
}