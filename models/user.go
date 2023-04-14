package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// User represents the model for an user
type User struct {
	GormModel
	Username    string        `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Username is required!"`
	Email       string        `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required!,email~Invalid format email!"`
	Password    string        `gorm:"not null" json:"-" form:"password" valid:"required~Password is required!,minstringlength(6)~Password minimum length must be 6 characters!"`
	Age         int           `gorm:"not null" json:"age" form:"age" valid:"required~Age is required!,range(9|99)~Age minimum value is 9"`
	Photo       []Photo       `json:"photos,omitempty"`
	Comment     []Comment     `json:"comments,omitempty"`
	Socialmedia []Socialmedia `json:"medsos,omitempty"`
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
