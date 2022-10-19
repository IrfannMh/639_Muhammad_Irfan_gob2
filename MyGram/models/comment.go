package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Message string `gorm:"not null" json:"message" form:"message" valid:"required"`
	UserID  uint   `gorm:"not null"`
	PhotoID uint   `gorm:"not null" json:"photo_id" form:"photo_id"`
	User    *User
	Photo   *Photo
}
