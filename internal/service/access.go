package service

import (
	"context"
	"simpson/internal/service/model"

	"gorm.io/gorm"
)

type accessService struct {
	gormDB *gorm.DB
}
type AccessService interface {
	AddAccess(ctx context.Context, access model.Access) error
	GetAccessByID(ctx context.Context, tx *gorm.DB, id uint) (model.Access, error)
}

func NewAccessService(
	db *gorm.DB,
) AccessService {
	return &accessService{
		gormDB: db,
	}
}

func (r *accessService) AddAccess(ctx context.Context, access model.Access) error {
	err := r.gormDB.Create(&access).Error
	return err
}

func (r *accessService) GetAccessByID(ctx context.Context, tx *gorm.DB, id uint) (model.Access, error) {
	var (
		access model.Access
		err    error
	)
	db := tx
	if tx == nil {
		db = r.gormDB
	}
	err = db.Table(access.Table()).Where("id = ?", id).First(&access).Error
	if err != nil {
		return access, err
	}
	return access, nil
}
