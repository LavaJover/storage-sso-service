package models

import "gorm.io/gorm"

type User struct{
	gorm.Model
	Email string `gorm:"column:email;type:varchar;size:255;unique"`
	Password string `gorm:"column:password;type:varchar;size:255"`
	IsActive bool `gorm:"column:is_active"`
}