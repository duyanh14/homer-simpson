package model

import "gorm.io/gorm"

type RolePermission struct {
	gorm.Model
	Description  string `gorm:"column:description" bson:"description"`
	CreatedBy    uint   `gorm:"column:created_by" bson:"created_by"`
	PermissionID uint   `gorm:"column:permission_id" bson:"permission_id"`
	RoleID       uint   `gorm:"column:role_id" bson:"role_id"`
}

func (u RolePermission) Table() string {
	return "role_permissions"
}
