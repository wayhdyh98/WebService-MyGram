package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId  uint
	PhotoId uint
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Username is required!"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errC := govalidator.ValidateStruct(c)

	if errC != nil {
		err = errC
		return
	}

	err = nil
	return
}
