package service

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Service interface {
	NewUserService() UserService
	NewPartnerService() PartnerService
	NewPermissionService() PermissionService
}
type service struct {
	gorm *gorm.DB
}

func InitService(ctx context.Context, db *gorm.DB) (Service, error) {
	if db == nil {
		return nil, errors.New("database init nil")
	}
	return &service{
		gorm: db,
	}, nil
}

func (r *service) NewUserService() UserService {
	return NewUserService(r.gorm)
}

func (r *service) NewPartnerService() PartnerService {
	return NewPartnerService(r.gorm)
}

func (r *service) NewPermissionService() PermissionService {
	return NewPermissionService(r.gorm)
}
