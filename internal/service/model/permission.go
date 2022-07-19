package model

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Name        string `gorm:"column:name" bson:"name"`
	Alias       string `gorm:"column:alias" bson:"alias"`
	Code        string `gorm:"column:code;unique" bson:"code"`
	Description string `gorm:"column:description" bson:"description"`
	CreatedBy   uint   `gorm:"column:created_by" bson:"created_by"`
}

func (u Permission) Table() string {
	return "permissions"
}
