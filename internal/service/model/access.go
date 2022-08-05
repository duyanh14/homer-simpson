package model

import "gorm.io/gorm"

type Access struct {
	gorm.Model
	Type        string `gorm:"column:type" bson:"type"`
	Description string `gorm:"column:description" bson:"description"`
	CreatedBy   uint   `gorm:"column:created_by" bson:"created_by"`
}

func (u Access) Table() string {
	return "accesses"
}
