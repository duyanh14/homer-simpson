package migration

import (
	"simpson/internal/service/model"

	"gorm.io/gorm"
)

func CreateTableAccess(db *gorm.DB) error {
	return db.Migrator().CreateTable(&model.Access{})
}
