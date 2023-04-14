package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId  uint
	PhotoId uint   `gorm:"foreignKey:Photo.ID" json:"photo_id"`
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Username is required!"`
	Photo   *Photo `gorm:"foreignKey:PhotoId" json:"photo"`
	User    *User  `gorm:"foreignKey:UserId" json:"user"`
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
