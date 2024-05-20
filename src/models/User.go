package models

import (
	_ "database/sql/driver"
	"gofiber/src/configs"

	"gorm.io/gorm"
)

type Role string

const (
	CUSTOMER Role = "customer"
	SELLER   Role = "seller"
)

type User struct {
	gorm.Model
	Name         string       `json:"name" validate:"required,min=3"`
	Email        string       `json:"email" validate:"required,email"`
	Password     string       `json:"password" validate:"required"`
	Phone_number string       `json:"phone_number" validate:"min=10,max=13"`
	Role         string       `json:"role" validate:"required"`
	Address      []APIAddress `json:"address"`
}

type APIAddress struct {
	gorm.Model
	Label      string `json:"label"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Address    string `json:"address" `
	PostalCode string `json:"postal_code"`
	City       string `json:"city"`
	Primary    bool   `json:"primary" gorm:"default:0"`
	UserID     int    `json:"user_id"`
}

func GetAllUser(sort, name string, limit, offset int) []*User {
	var results []*User
	name = "%" + name + "%"
	configs.DB.Preload("Address", func(db *gorm.DB) *gorm.DB {
		var items []*APIAddress
		return db.Model(&Address{}).Find(&items)
	}).Order(sort).Limit(limit).Offset(offset).Where("name ILIKE ?", name).Find(&results)
	return results
}

func GetDetailUser(id interface{}) *User {
	var user User
	configs.DB.Preload("Address", func(db *gorm.DB) *gorm.DB {
		var items []*APIAddress
		return db.Model(&Address{}).Find(&items)
	}).First(&user, "id = ?", id)
	return &user
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	results := configs.DB.Preload("Address", func(db *gorm.DB) *gorm.DB {
		var items []*APIAddress
		return db.Model(&Address{}).Find(&items)
	}).First(&user, "email = ?", email)
	return &user, results.Error
}

func PostUser(newUser *User) error {
	results := configs.DB.Create(&newUser)
	return results.Error
}
func UpdateUser(id int, user *User) error {
	results := configs.DB.Model(&User{}).Where("id = ?", id).Updates(user)
	return results.Error
}

func DeleteUser(id int) error {
	results := configs.DB.Delete(&User{}, "id = ?", id)
	return results.Error
}

