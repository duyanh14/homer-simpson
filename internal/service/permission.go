package service

import (
	"context"
	"simpson/internal/service/model"

	"gorm.io/gorm"
)

type permissionService struct {
	gormDB *gorm.DB
}
type PermissionService interface {
	AddPermission(ctx context.Context, permission model.Permission) error
	GetPermissionByID(ctx context.Context, tx *gorm.DB, id uint) (model.Permission, error)
}

func NewPermissionService(
	db *gorm.DB,
) PermissionService {
	return &permissionService{
		gormDB: db,
	}
}

func (r *permissionService) AddPermission(ctx context.Context, permission model.Permission) error {
	err := r.gormDB.Create(&permission).Error
	return err
}

func (r *permissionService) GetPermissionByID(ctx context.Context, tx *gorm.DB, id uint) (model.Permission, error) {
	var (
		per model.Permission
		err error
	)
	db := tx
	if tx == nil {
		db = r.gormDB
	}
	err = db.Table(per.Table()).Where("id = ?", id).First(&per).Error
	if err != nil {
		return per, err
	}
	return per, nil
}
