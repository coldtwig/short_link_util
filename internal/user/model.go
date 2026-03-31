package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" validate:"required,email" gorm:"uniqueIndex"`
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
