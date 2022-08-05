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
	NewRoleService() RoleService
	NewCommonService() CommonService
	NewUserRoleService() UserRoleService
	NewRolePermisiosnService() RolePermissionService
	NewAccessService() AccessService
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
func (r *service) BuildTransaction(ctx context.Context) *gorm.DB {
	return r.gorm
}

func (r *service) NewRolePermisiosnService() RolePermissionService {
	return NewRolePermissionService(r.gorm)
}

func (r *service) NewUserRoleService() UserRoleService {
	return NewUserRoleService(r.gorm)
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

func (r *service) NewRoleService() RoleService {
	return NewRoleService(r.gorm)
}

func (r *service) NewAccessService() AccessService {
	return NewAccessService(r.gorm)
}

func (r *service) NewCommonService() CommonService {
	return NewCommonService(r.gorm)
}

// common server
type CommonService interface {
	GetDatabaseTx(ctx context.Context) *gorm.DB
}

type commonService struct {
	gormDB *gorm.DB
}

func (r *commonService) GetDatabaseTx(ctx context.Context) *gorm.DB {
	return r.gormDB.Begin()
}

func NewCommonService(
	db *gorm.DB,
) CommonService {
	return &commonService{
		gormDB: db,
	}
}
