package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	StringableModel

	Username string `form:"username" gorm:"type:varchar(255)" binding:"required" json:"username"`
	Password string `form:"password" gorm:"type:varchar(255)" binding:"required" json:"password"`
}
