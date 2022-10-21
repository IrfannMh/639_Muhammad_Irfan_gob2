package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required"`
	UserID         uint
	User           User `valid:"-"`
}

func (sm *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(sm)

	if errCreate != nil {
		err = errCreate
		return
	}
	return
}
