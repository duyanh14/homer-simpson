package service

import (
	"context"
	"simpson/internal/service/model"

	"gorm.io/gorm"
)

type userUserRoleService struct {
	gormDB *gorm.DB
}
type UserRoleService interface {
	AddUserRole(ctx context.Context, tx *gorm.DB, userUserRole model.UserRole) error
}

func NewUserRoleService(
	db *gorm.DB,
) UserRoleService {
	return &userUserRoleService{
		gormDB: db,
	}
}

func (r *userUserRoleService) AddUserRole(ctx context.Context, tx *gorm.DB, userUserRole model.UserRole) error {
	db := tx
	if tx == nil {
		db = r.gormDB
	}
	err := db.Create(&userUserRole).Error
	return err
}
