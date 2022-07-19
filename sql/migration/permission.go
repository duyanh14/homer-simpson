package migration

import (
	"simpson/internal/service/model"

	"gorm.io/gorm"
)

func CreateTablePermission(db *gorm.DB) error {
	return db.Migrator().CreateTable(&model.Permission{})
}
