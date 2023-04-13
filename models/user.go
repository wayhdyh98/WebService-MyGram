package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Username is required!"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required!,email~Invalid format email!"`
	Password string `gorm:"not null" json:"-" form:"password" valid:"required~Password is required!,minstringlength(6)~Password minimum length must be 6 characters!"`
	Age      int    `gorm:"not null" json:"age" form:"age" valid:"required~Age is required!,min=9~Age minimum value is 9"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errC := govalidator.ValidateStruct(u)

	if errC != nil {
		err = errC
		return
	}

	err = nil
	return
}
