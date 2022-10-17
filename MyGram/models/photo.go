package models

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title    string `gorm:"not null" json:"title" form:"title" valid:"required"`
	Caption  string `json:"caption" form:"caption"`
	PhotoURL string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required"`
	UserID   uint
	Comment  []Comment
}
