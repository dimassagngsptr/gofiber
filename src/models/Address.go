package models

import (
	"gofiber/src/configs"

	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	Label      string `json:"label" validate:"required,min=3"`
	Name       string `json:"name" validate:"required,min=5"`
	Phone      string `json:"phone" validate:"required"`
	Address    string `json:"address" validate:"required,min=20"`
	PostalCode string `json:"postal_code" validate:"required,min=5"`
	City       string `json:"city" validate:"required"`
	Primary    bool   `json:"primary" gorm:"default:0"`
	UserID int `json:"user_id"`
	User User `gorm:"foreignKey:UserID"`
}

func GetAllAddress() []*Address{
	var results []*Address
	configs.DB.Preload("User").Find(&results)
	return results
}

func GetAddress(id int) *Address {
	var results Address
	configs.DB.Preload("User").First(&results,"id = ?",id)
	return &results
}

func GetAddressByNameAndAddress(name string, address string, id uint) *Address{
	var results Address
	configs.DB.Model(&Address{}).Where("name = ? AND address = ? AND ID = ?", name, address, id).First(&results)
	return &results
}

func PostAddress(newAddress *Address) error{
	results := configs.DB.Create(&newAddress)
	return results.Error
}

func UpdateAddress(id int, newAddress *Address) error{
	results:= configs.DB.Model(&Address{}).Where("id = ?", id).Updates(newAddress)
	return results.Error
}

func DeleteAddress(id int) error{
	results:=configs.DB.Delete(&Address{}, "id = ?", id)
	return results.Error
}
