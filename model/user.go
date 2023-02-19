package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PasswordDigest string //加密之后的密码
}

func (user *User) SetPassword(password string) error {
	passwordDigest, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err == nil {
		user.PasswordDigest = string(passwordDigest)
	}
	return err
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
