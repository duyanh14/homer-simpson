package model

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	Description string `gorm:"column:description" bson:"description"`
	CreatedBy   uint   `gorm:"column:created_by" bson:"created_by"`
	UserID      uint   `gorm:"column:user_id;not null;" bson:"user_id"`
	RoleID      uint   `gorm:"column:role_id;not null" bson:"role_id"`
}

func (u UserRole) Table() string {
	return "user_roles"
}
