package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Socialmedia struct {
	gorm.Model
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~Name is required!"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url" valid:"required~Social media is required!"`
	UserId         uint
}

func (s *Socialmedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errC := govalidator.ValidateStruct(s)

	if errC != nil {
		err = errC
		return
	}

	err = nil
	return
}
