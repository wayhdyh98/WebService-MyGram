package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title    string `gorm:"not null" json:"title" form:"title" valid:"required~Photo title is required!"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Password is required!"`
	UserId   uint
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errC := govalidator.ValidateStruct(p)

	if errC != nil {
		err = errC
		return
	}

	err = nil
	return
}
