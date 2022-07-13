package db

import (
	"fmt"
	"simpson/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres(config *config.Postgres) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.Host, config.Username, config.Password, config.Database, config.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
