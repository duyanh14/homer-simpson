package db

import (
	"fmt"
	"simpson/config"
	"simpson/sql/migration"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres(config config.Postgres) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.Host, config.Username, config.Password, config.Database, config.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	zap.S().Debug("connect to postgress successful")
	if config.Migrate {
		Migration(db)
	}
	return db, err
}
func Migration(db *gorm.DB) {
	err := migration.CreateTableUser(db)
	if err != nil {
		zap.S().Error("migrator create table user err %s", err)
	}
	err = migration.CreateTablePermission(db)
	if err != nil {
		zap.S().Error("migrator create table permission err %s", err)
	}
	err = migration.CreateTableRole(db)
	if err != nil {
		zap.S().Error("migrator create table role err %s", err)
	}

	err = migration.CreateTableUserRole(db)
	if err != nil {
		zap.S().Error("migrator create table user role err %s", err)
	}
}
