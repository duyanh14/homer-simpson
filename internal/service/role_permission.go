package service

import (
	"context"
	"fmt"
	"simpson/internal/service/model"

	"gorm.io/gorm"
)

type rolePermissionService struct {
	gormDB *gorm.DB
}

type RolePermissionService interface {
	AddRolePermission(ctx context.Context, tx *gorm.DB, rolePer model.RolePermission) error
	GetPermissionsByRoleIDs(ctx context.Context, roleIDs string) ([]model.Permission, error)
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
