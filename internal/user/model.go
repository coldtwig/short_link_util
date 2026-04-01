package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"index"`
	Password string
	Name     string
}

func NewUser() *User {
	return &User{
		Email:    "",
		Password: "",
		Name:     "",
	}
}
