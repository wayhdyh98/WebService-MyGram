package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Socialmedia represents the model for an socialmedia
type Socialmedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~Name is required!"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url" valid:"required~Social media is required!"`
	UserId         uint
	User           *User `gorm:"foreignKey:UserId" json:"user"`
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
