package migration

import (
	"simpson/internal/service/model"

	"gorm.io/gorm"
)

func CreateTableRole(db *gorm.DB) error {
	return db.Migrator().CreateTable(&model.Role{})
}
