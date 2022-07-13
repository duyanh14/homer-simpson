package repository

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository interface {
	NewUserRepo() UserRepo
	NewPartnerRepo() PartnerRepo
}
type repository struct {
	gorm *gorm.DB
}

func InitRepository(ctx context.Context, db *gorm.DB) (Repository, error) {
	if db == nil {
		return nil, errors.New("database init nil")
	}
	return &repository{
		gorm: db,
	}, nil
}

func (r *repository) NewUserRepo() UserRepo {
	zap.S().Debug("repository new user repositofy successful")
	return NewUserRepo(r.gorm)
}

func (r *repository) NewPartnerRepo() PartnerRepo {
	return NewPartnerRepo(r.gorm)
}
