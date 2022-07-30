package migration

import (
	"simpson/internal/service/model"

	"gorm.io/gorm"
)

func CreateTableRolePermission(db *gorm.DB) error {
	return db.Migrator().CreateTable(&model.RolePermission{})
}
