package models

import (
	"gofiber/src/configs"

	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	Label      string `json:"label"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	PostalCode string `json:"postal_code"`
	City       string `json:"city"`
	Primary    bool   `json:"primary" gorm:"default:0"`
	UserID int `json:"user_id"`
	User User `gorm:"foreignKey:UserID"`
}

func GetAllAddress() ([]*Address, int64){
	var results []*Address
	var count int64
	configs.DB.Preload("User").Find(&results).Count(&count)
	return results, count
}

func GetAddress(id int) *Address {
	var results Address
	configs.DB.Preload("User").First(&results,"id = ?",id)
	return &results
}

func GetAddressByNameAndAddress(name string, address string) *Address{
	var results Address
	configs.DB.Model(&Address{}).Where("name = ? AND address = ?", name, address).First(&results)
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
