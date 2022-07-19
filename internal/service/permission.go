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
	// GetPermissionByPermissionname(ctx context.Context, permissionname string) (model.Permission, error)
	// GetPermissionByPhone(ctx context.Context, phone string) (model.Permission, error)
	// GetPermissionByEmail(ctx context.Context, email string) (model.Permission, error)
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

// func (r *permissionService) GetPermissionByPermissionname(ctx context.Context, permissionname string) (model.Permission, error) {
// 	var (
// 		permission model.Permission
// 		err  error
// 	)
// 	err = r.gormDB.Table(permission.Table()).Where("permissionname = ?", permissionname).First(&permission).Error
// 	if err != nil {
// 		return permission, err
// 	}
// 	return permission, nil
// }

// func (r *permissionService) GetPermissionByEmail(ctx context.Context, email string) (model.Permission, error) {
// 	var (
// 		permission model.Permission
// 		err  error
// 	)
// 	err = r.gormDB.Table(permission.Table()).Where("email = ?", email).First(&permission).Error
// 	if err != nil {
// 		return permission, err
// 	}
// 	return permission, nil
// }

// func (r *permissionService) GetPermissionByPhone(ctx context.Context, phone string) (model.Permission, error) {
// 	var (
// 		permission model.Permission
// 		err  error
// 	)
// 	err = r.gormDB.Table(permission.Table()).Where("phone = ?", phone).First(&permission).Error
// 	if err != nil {
// 		return permission, err
// 	}
// 	return permission, nil
// }
