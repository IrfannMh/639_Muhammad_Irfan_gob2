package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Message string `gorm:"not null" json:"message" form:"message" valid:"required"`
	UserID  uint
	PhotoID uint
}
