package migration

import (
	"simpson/internal/service/model"

	"gorm.io/gorm"
)

func CreateTableUser(db *gorm.DB) error {
	return db.Migrator().CreateTable(&model.User{})
}
