package models

import (
	"gofiber/src/configs"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name  string  `json:"name"`
	Email string `json:"email"`
	Password string     `json:"password"`
}

func GetAllUser() ([]*Users, int64){
	var results []*Users
	var count int64
	configs.DB.Find(&results).Count(&count)
	return results, count
}

func GetUserById(id int) *Users{
	var user Users
	configs.DB.First(&user, "id = ?", id)
	return &user
}

func GetUserByEmail(email string) *Users{
	var user Users
	configs.DB.First(&user, "email = ?",email)
	return &user
}

func PostUser(newUser *Users) error{
	results:=configs.DB.Create(&newUser)
	return results.Error
}
func UpdateUser(id int, user *Users) error{
	results:= configs.DB.Model(&Users{}).Where("id = ?", id).Updates(user)
	return results.Error
}

func DeleteUser(id int) error{
	results:= configs.DB.Delete(&Users{}, "id = ?", id)
	return results.Error
}