package models

import (
	"gofiber/src/configs"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string  `json:"name" validate:"required,min=3"`
	Email string `json:"email" validate:"required"`
	Password string     `json:"password" validate:"required"`
	Phone_number string `json:"phone_number"`
	Address []Address `json:"address"`

}

func GetAllUser() ([]*User, int64){
	var results []*User
	var count int64
	configs.DB.Preload("Address").Find(&results).Count(&count)
	return results, count
}

func GetUserById(id int) *User{
	var user User
	configs.DB.Preload("Address").First(&user, "id = ?", id)
	return &user
}

func GetUserByEmail(email string) (*User, error){
	var user User
	results:=configs.DB.Preload("Address").First(&user, "email = ?",email)
	return &user,results.Error 
}

func PostUser(newUser *User) error{
	results:=configs.DB.Create(&newUser)
	return results.Error
}
func UpdateUser(id int, user *User) error{
	results:= configs.DB.Model(&User{}).Where("id = ?", id).Updates(user)
	return results.Error
}

func DeleteUser(id int) error{
	results:= configs.DB.Delete(&User{}, "id = ?", id)
	return results.Error
}