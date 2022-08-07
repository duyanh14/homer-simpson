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
	GetPermissioByCode(ctx context.Context, code string) (model.Permission, error)
	GetPermissionByID(ctx context.Context, tx *gorm.DB, id uint) (model.Permission, error)
	// management
	ListPermission(ctx context.Context) ([]model.Permission, error)
	DeletePermission(ctx context.Context, permissionID uint) (int64, error)
	UploadPermission(ctx context.Context, role model.Permission) (int64, error)
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

func (r *permissionService) GetPermissioByCode(ctx context.Context, code string) (model.Permission, error) {
	var (
		per model.Permission
		err error
	)
	db := r.gormDB
	err = db.Table(per.Table()).Where("code = ?", code).First(&per).Error
	if err != nil {
		return per, err
	}
	return per, err
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

func (r *permissionService) ListPermission(ctx context.Context) ([]model.Permission, error) {
	var (
		permissions = []model.Permission{}
		err         error
	)
	err = r.gormDB.Table(model.Permission{}.Table()).Find(&permissions).Error
	return permissions, err
}

func (r *permissionService) DeletePermission(ctx context.Context, permissionID uint) (int64, error) {
	permission := model.Permission{}
	deleteDB := r.gormDB.Table(permission.Table()).Where("id = ?", permissionID).Delete(&permission)
	if err := deleteDB.Error; err != nil {
		return 0, err
	}
	return deleteDB.RowsAffected, nil
}

func (r *permissionService) UploadPermission(ctx context.Context, permission model.Permission) (int64, error) {
	updateDB := r.gormDB.Table(permission.Table()).Where("id = ?", permission.ID).Updates(permission.ColumnUpdate())
	if err := updateDB.Error; err != nil {
		return 0, err
	}
	return updateDB.RowsAffected, nil
}
