package service

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	NewUserService() UserService
	NewPartnerService() PartnerService
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
	zap.S().Debug("service new user repositofy successful")
	return NewUserService(r.gorm)
}

func (r *service) NewPartnerService() PartnerService {
	return NewPartnerService(r.gorm)
}
