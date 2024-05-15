package models

import (
	"gofiber/src/configs"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name"`
	Image string `json:"image"`
	Product []Product `json:"product"`

}

type APIProduct struct{
	gorm.Model
	Name string `json:"name"`
	Price float64 `json:"price"`
	Descriptions string `json:"descriptions"`
	Image string `json:"image"`
	Stock int `json:"stock"`
	CategoryID uint `json:"category_id"`
}

func GetAllCategories() ([]*Category, int64){
	var categories []*Category
	var count int64
	configs.DB.Preload("Product", func(db *gorm.DB) *gorm.DB {
		var items []*APIProduct
		return db.Model(&Product{}).Find(&items)
	}).Find(&categories).Count(&count)
	return categories,count
}

func GetCategoryById(id int) *Category{
	var category Category
	configs.DB.Preload("Product",func(db *gorm.DB)*gorm.DB{
		var results []*APIProduct
		return db.Model(&Product{}).Find(&results)
	}).First(&category,"id = ?",id)
	return &category
}

func GetCategoryByName(name string) *Category{
	var category Category
	configs.DB.First(&category,"name = ?", name)
	return &category
}

func PostCategory(newCategory *Category) error{
	results:= configs.DB.Create(&newCategory)
	return results.Error
}

func UpdateCategory(id int, newCategory *Category) error{
	results:= configs.DB.Model(&Category{}).Where("id = ?", id).Updates(newCategory)
	return results.Error
}

func DeleteCategory(id int) error{
	results:= configs.DB.Delete(&Category{}, "id = ?", id)
	return results.Error
}