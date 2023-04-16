package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Comment represents the model for an comment
type Comment struct {
	GormModel
	UserId  uint
	PhotoId uint   `gorm:"foreignKey:Photo.ID" json:"photo_id"`
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Username is required!"`
	Photo   *Photo `gorm:"foreignKey:PhotoId"`
	User    *User  `gorm:"foreignKey:UserId"`
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
