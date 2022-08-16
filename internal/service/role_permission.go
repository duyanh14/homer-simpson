package service

import (
	"context"
	"fmt"
	"simpson/internal/service/model"

	"gorm.io/gorm"
)

type rolePermissionService struct {
	gormDB  *gorm.DB
	isDebug bool
}

type RolePermissionService interface {
	AddRolePermission(ctx context.Context, tx *gorm.DB, rolePer model.RolePermission) error
	GetPermissionsByRoleIDs(ctx context.Context, roleIDs string) ([]model.Permission, error)
	GetPermissionsByRoleID(ctx context.Context, roleID uint) ([]model.Permission, error)
	GetRolesByPermissionID(ctx context.Context, permissionID uint) ([]model.Role, error)
}

func NewRolePermissionService(
	db *gorm.DB,
	isDebug bool,
) RolePermissionService {
	return &rolePermissionService{
		gormDB:  db,
		isDebug: isDebug,
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

func (r *rolePermissionService) GetPermissionsByRoleIDs(ctx context.Context, roleIDs string) ([]model.Permission, error) {
	var (
		pers []model.Permission
		err  error
	)

	preSql := `
		select permissions.id, permissions.name, permissions.code, permissions.alias from permissions 
		join role_permissions on permissions.id = role_permissions.permission_id
		where permissions.deleted_at is null 
		and role_permissions.role_id in 
	`
	sql := fmt.Sprintf("%s (%s)", preSql, roleIDs)
	err = r.gormDB.Raw(sql).Find(&pers).Error
	// db := r.gormDB.Table(model.Role{}.Table()).Joins(model.UserRole{}.Table()).Where("user_roles.user_id = ?", userID).Find(&roles)
	return pers, err
}

func (r *rolePermissionService) GetPermissionsByRoleID(ctx context.Context, roleID uint) ([]model.Permission, error) {
	var (
		pers []model.Permission
		err  error
	)
	db := r.gormDB
	db = AppendSql(db, r.isDebug, GetAll)
	preSql := `
		select permissions.id, permissions.name, permissions.code, permissions.alias from permissions 
		join role_permissions on permissions.id = role_permissions.permission_id
		where permissions.deleted_at is null 
		and role_permissions.role_id = ? 
	`
	err = db.Raw(preSql, roleID).Find(&pers).Error
	return pers, err
}

func (r *rolePermissionService) GetRolesByPermissionID(ctx context.Context, permissionID uint) ([]model.Role, error) {
	var (
		role []model.Role
		err  error
	)
	db := r.gormDB
	db = AppendSql(db, r.isDebug, GetAll)
	preSql := `
		select roles.id, roles.name, roles.code, roles.alias, roles.created_at, roles.updated_at,roles.description,roles.deleted_at,roles.created_by
		from roles 
		join role_permissions on roles.id = role_permissions.role_id
		and role_permissions.permission_id = ? 
	`
	err = db.Raw(preSql, permissionID).Find(&role).Error
	return role, err
}
