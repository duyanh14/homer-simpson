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
	// management
	ListRole(ctx context.Context) ([]model.Role, error)
	DeleteRole(ctx context.Context, roleID uint) (int64, error)
	UploadRole(ctx context.Context, role model.Role) (int64, error)
}

func NewRoleService(
	db *gorm.DB,
) RoleService {
	return &roleService{
		gormDB: db,
	}
}

func (r *roleService) ListRole(ctx context.Context) ([]model.Role, error) {
	var (
		roles = []model.Role{}
		err   error
	)
	err = r.gormDB.Table(model.Role{}.Table()).Find(&roles).Error
	return roles, err
}

func (r *roleService) DeleteRole(ctx context.Context, roleID uint) (int64, error) {
	role := model.Role{}
	deleteDB := r.gormDB.Table(role.Table()).Where("id = ?", roleID).Delete(&role)
	if err := deleteDB.Error; err != nil {
		return 0, err
	}
	return deleteDB.RowsAffected, nil
}

func (r *roleService) UploadRole(ctx context.Context, role model.Role) (int64, error) {
	updateDB := r.gormDB.Table(role.Table()).Where("id = ?", role.ID).Updates(role.ColumnUpdate())
	if err := updateDB.Error; err != nil {
		return 0, err
	}
	return updateDB.RowsAffected, nil
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
