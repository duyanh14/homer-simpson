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
