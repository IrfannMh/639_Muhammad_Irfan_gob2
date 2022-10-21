package models

import (
	"MyGram/helpers"
	"errors"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint
	Username string `gorm:"not null" json:"username" form:"username" valid:"required"`
	Email    string `gorm:"not null" json:"email" form:"email" valid:"required,email"`
	Password string `gorm:"not null" json:"password"form:"password" valid:"required"`
	Age      int    `gorm:"not null" json:"age" form:"age" valid:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Age < 8 {
		return errors.New("your age must be at least 8")
	}
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	var user User

	errInputUsername := tx.Model(&User{}).Where("username = ?", u.Username).First(&user).Error
	errInputEmail := tx.Model(&User{}).Where("email = ?", u.Email).First(&user).Error
	if errInputUsername == nil {
		return errors.New("username already exist")
	} else if errInputEmail == nil {
		return errors.New("email already exist")
	}

	if len(u.Password) <= 6 {
		return errors.New("password length must be at least 6 character")
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
