package models

import (
	"gofiber/src/configs"

	"gorm.io/gorm"
)

type UserAddress struct {
	gorm.Model
	Label      string `json:"label"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	PostalCode string `json:"postal_code"`
	City       string `json:"city"`
	Primary    bool   `json:"primary" gorm:"default:0"`
}

func GetAllUserAddress() []*UserAddress{
	var results []*UserAddress
	configs.DB.Find(&results)
	return results
}

func GetUserAddress(id int) *UserAddress {
	var results UserAddress
	configs.DB.First(&results,"id = ?",id)
	return &results
}

func GetUserAddressByNameAndAddress(name string, address string) *UserAddress{
	var results UserAddress
	configs.DB.Model(&UserAddress{}).Where("name = ? AND address = ?", name, address).First(&results)
	return &results
}

func PostUserAddress(newUserAddress *UserAddress) error{
	results := configs.DB.Create(&newUserAddress)
	return results.Error
}

func UpdateUserAddress(id int, newUserAddress *UserAddress) error{
	results:= configs.DB.Model(&UserAddress{}).Where("id = ?", id).Updates(newUserAddress)
	return results.Error
}

func DeleteUserAddress(id int) error{
	results:=configs.DB.Delete(&UserAddress{}, "id = ?", id)
	return results.Error
}
