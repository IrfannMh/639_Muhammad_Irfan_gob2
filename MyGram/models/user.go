package models

import (
	"MyGram/helpers"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint
	Username    string `gorm:"not null" json:"username" form:"username" valid:"required"`
	Email       string `gorm:"not null" json:"email" form:"email" valid:"required,email"`
	Password    string `gorm:"not null" json:"password"form:"password" valid:"required"`
	Age         int    `gorm:"not null" json:"age" form:"age" valid:"required"`
	SocialMedia SocialMedia
	Photo       []Photo
	Comment     []Comment
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// _, errCreate := govalidator.ValidateStruct(u)

	// if errCreate != nil {
	// 	err = errCreate
	// 	return
	// }

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
