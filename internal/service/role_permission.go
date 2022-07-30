package service

import (
	"context"
	"simpson/internal/service/model"

	"gorm.io/gorm"
)

type rolePermissionService struct {
	gormDB *gorm.DB
}

type RolePermissionService interface {
	AddRolePermission(ctx context.Context, tx *gorm.DB, rolePer model.RolePermission) error
}

func NewRolePermissionService(
	db *gorm.DB,
) RolePermissionService {
	return &rolePermissionService{
		gormDB: db,
	}
}

func (r *rolePermissionService) AddRolePermission(ctx context.Context, tx *gorm.DB, rolePermission model.RolePermission) error {
	db := tx
	if tx == nil {
		db = r.gormDB
	}
	err := db.Create(&rolePermission).Error
	return err
}
