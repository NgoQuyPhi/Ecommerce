package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserId   int    `json:"user_id"`
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) bool {
	if user.Password == providedPassword {
		return true
	} else {
		return false
	}
}
