package models

import (
	"crudapi/pkgs/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email"`
	Age     int    `json:"age"`
	Phone   string `json:"phone"`
	Country string `json:"country"`
}

func init() {
	config.InitDB()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("ID = ?", Id).Find(&getUser)
	return &getUser, db
}

func GetUserByEmail(Email string) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("Email = ?", Email).Find(&getUser)
	return &getUser, db
}

func DeleteUser(Id int64) User {
	var user User
	db.Where("ID = ?", Id).Delete(user)
	return user
}
