package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Photo represents the model for an photo
type Photo struct {
	GormModel
	Title    string `gorm:"not null" json:"title" form:"title" valid:"required~Photo title is required!"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Password is required!"`
	UserId   uint
	User     *User `gorm:"foreignKey:UserId" json:"user"`
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
