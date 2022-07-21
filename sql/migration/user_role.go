package migration

import (
	"simpson/internal/service/model"

	"gorm.io/gorm"
)

func CreateTableUserRole(db *gorm.DB) error {
	return db.Migrator().CreateTable(&model.UserRole{})
}
