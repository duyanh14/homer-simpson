package service

import (
	"context"
	"simpson/internal/service/model"

	"gorm.io/gorm"
)

type rolePermissionPermissionService struct {
	gormDB *gorm.DB
}
type RolePermissionService interface {
	AddRolePermission(ctx context.Context, tx *gorm.DB, rolePermissionPermission model.RolePermission) error
}

func NewRolePermissionService(
	db *gorm.DB,
) RolePermissionService {
	return &rolePermissionPermissionService{
		gormDB: db,
	}
}

func (r *rolePermissionPermissionService) AddRolePermission(ctx context.Context, tx *gorm.DB, rolePermissionPermission model.RolePermission) error {
	db := tx
	if tx == nil {
		db = r.gormDB
	}
	err := db.Create(&rolePermissionPermission).Error
	return err
}
