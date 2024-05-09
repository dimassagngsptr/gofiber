package models

import (
	"gofiber/src/configs"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name string `json:"name"`
	Price float64 `json:"price"`
	Descriptions string `json:"descriptions"`
	Image string `json:"image"`
	Stock int `json:"stock"`
	CategoryID int `json:"category_id"`
	Category Category `gorm:"foreignKey:CategoryID"`
}
func SelectAllProduct() ([]*Product, int64) {
	var items []*Product
	var count int64
	configs.DB.Preload("Category").Find(&items).Count(&count)
	return items,count
}

func SelectProductById(id int) *Product {
	var item Product
	configs.DB.Preload("Category").First(&item, "id = ?", id)
	return &item
}

func PostProduct(newProduct *Product) error {
	result := configs.DB.Create(&newProduct)
	return result.Error
}

func UpdateProduct(id int, item *Product) error {
	result := configs.DB.Model(&Product{}).Where("id = ?", id).Updates(item)
	return result.Error
}

func DeleteProduct(id int) error {
	result := configs.DB.Delete(&Product{}, "id = ?", id)
	return result.Error
}