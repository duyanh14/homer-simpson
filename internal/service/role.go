package service

import (
	"context"
	"simpson/internal/service/model"

	"gorm.io/gorm"
)

type roleService struct {
	gormDB *gorm.DB
}
type RoleService interface {
	AddRole(ctx context.Context, role model.Role) error
	GetRoleByID(ctx context.Context, tx *gorm.DB, id uint) (model.Role, error)
}

func NewRoleService(
	db *gorm.DB,
) RoleService {
	return &roleService{
		gormDB: db,
	}
}

func (r *roleService) AddRole(ctx context.Context, role model.Role) error {
	err := r.gormDB.Create(&role).Error
	return err
}

func (r *roleService) GetRoleByID(ctx context.Context, tx *gorm.DB, id uint) (model.Role, error) {
	var (
		role model.Role
		err  error
	)
	db := tx
	if tx == nil {
		db = r.gormDB
	}
	err = db.Table(role.Table()).Where("id = ?", id).First(&role).Error
	if err != nil {
		return role, err
	}
	return role, nil
}
