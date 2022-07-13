package repository

import (
	"context"
	"simpson/internal/repository/model"

	"gorm.io/gorm"
)

type partnerRepo struct {
	gormDB *gorm.DB
}
type PartnerRepo interface {
	AddPartner(ctx context.Context, partner model.Partner) error
}

func NewPartnerRepo(
	db *gorm.DB,
) PartnerRepo {
	return &partnerRepo{
		gormDB: db,
	}
}

func (r *partnerRepo) AddPartner(ctx context.Context, partner model.Partner) error {
	return r.gormDB.Create(&partner).Error
}
