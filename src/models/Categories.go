package models

import (
	"gofiber/src/configs"

	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	Name string `json:"name"`

}

func GetAllCategories() ([]*Categories, int64){
	var categories []*Categories
	var count int64
	configs.DB.Find(&categories).Count(&count)
	return categories,count
}

func GetCategoryById(id int) *Categories{
	var category Categories
	configs.DB.First(&category,"id = ?",id)
	return &category
}

func GetCategoryByName(name string) *Categories{
	var category Categories
	configs.DB.First(&category,"name = ?", name)
	return &category
}

func PostCategory(newCategory *Categories) error{
	results:= configs.DB.Create(&newCategory)
	return results.Error
}

func UpdateCategory(id int, newCategory *Categories) error{
	results:= configs.DB.Model(&Categories{}).Where("id = ?", id).Updates(newCategory)
	return results.Error
}

func DeleteCategory(id int) error{
	results:= configs.DB.Delete(&Categories{}, "id = ?", id)
	return results.Error
}