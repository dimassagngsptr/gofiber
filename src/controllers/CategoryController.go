package controllers

import (
	"fmt"
	"gofiber/src/helpers"
	"gofiber/src/models"
	"math"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetAllCategory(c *fiber.Ctx) error {
	keyword := c.Query("search")
	sort := c.Query("sort")
	sortBy := c.Query("orderBy")
	pageOld := c.Query("page")
	limitOld := c.Query("limit")
	page, _ := strconv.Atoi(pageOld)
	if page == 0{
		page =1
	}
	limit, _ := strconv.Atoi(limitOld)
	if limit == 0{
		limit = 5
	}
	offset := (page -1) * limit
	if sort == ""{
		sort = "ASC"
	}
	if sortBy == ""{
		sortBy ="name"
	}
	sort = sortBy + " " + strings.ToLower(sort)
	categories := models.GetAllCategories(sort,keyword,limit,offset)
	count := helpers.CountData("category")
	totalPage := math.Ceil(float64(count)/float64(limit))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Successfully retrieved all categories",
		"data":categories,
		"totalPage":totalPage,
		"totalData":count,
		"limit":limit,
		"page":page,
	})
}

func GetCategoryById(c *fiber.Ctx) error{
	id, _ := strconv.Atoi(c.Params("id"))
	foundCategory:= models.GetCategoryById(id)
	if foundCategory.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message":"Category not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Successfully found category",
		"data":foundCategory,

	})
}

func CreateCategory(c *fiber.Ctx) error{
	var newCategory models.Category
	if err:= c.BodyParser(&newCategory); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	models.PostCategory(&newCategory)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Category has been created",
	})
}

func UpdateCategory(c *fiber.Ctx) error {
	var newCategory models.Category
	id, _ := strconv.Atoi(c.Params("id"))
	if err := c.BodyParser(&newCategory); err !=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Invalid request body",
		})
	}
	err :=models.UpdateCategory(id, &newCategory)
		if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Failed to update category with ID %d", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Success Update category with ID %d", id),
	})
}

func DeleteCategory(c *fiber.Ctx) error{
	id, _ := strconv.Atoi(c.Params("id"))
	err := models.DeleteCategory(id)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":fmt.Sprintf("Failed to delete category with ID %d", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":fmt.Sprintf("Category with ID %d deleted successfully", id),
	})
}