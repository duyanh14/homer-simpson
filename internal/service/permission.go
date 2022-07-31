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
	// get list permission of userID
	GetPermissionsUserId(ctx context.Context, userID uint) ([]model.Permission, error)
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

func (r *permissionService) GetPermissionsUserId(ctx context.Context, userID uint) ([]model.Permission, error) {
	var (
		pers []model.Permission
		err  error
	)
	sql := `
		select permissions.id, permissions.name, permissions.alias, permissions.code, permissions.description 
		from permissions 
		left join role_permissions on permissions.id = role_permissions.permission_id
		left join roles on role_permissions.role_id = roles.id
		left join user_roles on roles.id = user_roles.user_id
		where user_roles.user_id = 3 and permissions.deleted_at is null;
	`
	r.gormDB.Exec(sql)
	return pers, err
}
