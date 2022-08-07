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

func (role Permission) ColumnUpdate() map[string]interface{} {
	fieldUpdate := make(map[string]interface{})
	if role.Alias != "" {
		fieldUpdate["alias"] = role.Alias
	}
	if role.Name != "" {
		fieldUpdate["name"] = role.Name
	}
	if role.Code != "" {
		fieldUpdate["code"] = role.Code
	}
	if role.Description != "" {
		fieldUpdate["description"] = role.Description
	}
	if role.CreatedBy != 0 {
		fieldUpdate["created_by"] = role.CreatedBy
	}
	return fieldUpdate
}
