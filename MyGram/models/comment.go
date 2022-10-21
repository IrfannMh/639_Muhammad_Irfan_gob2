package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Message string `gorm:"not null" json:"message" form:"message" valid:"required"`
	UserID  uint   `gorm:"not null"`
	PhotoID uint   `gorm:"not null" json:"photo_id" form:"photo_id"`
	User    User   `valid:"-"`
	Photo   Photo  `valid:"-"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}
	return
}
