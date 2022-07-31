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
	GetRolesByUserID(ctx context.Context, userID uint) ([]model.Role, error)
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

func (r *userUserRoleService) GetRolesByUserID(ctx context.Context, userID uint) ([]model.Role, error) {
	var (
		roles []model.Role
		err   error
	)
	sql := `
		select roles.id, roles.name, roles.code, roles.alias from roles 
		join user_roles on roles.id = user_roles.role_id
		where roles.deleted_at is null 
		and user_roles.user_id = ?
	`
	err = r.gormDB.Raw(sql, userID).Find(&roles).Error
	return roles, err
}
