package service

import (
	"context"
	"simpson/internal/service/model"

	"gorm.io/gorm"
)

type partnerService struct {
	gormDB *gorm.DB
}
type PartnerService interface {
	AddPartner(ctx context.Context, partner model.Partner) error
}

func NewPartnerService(
	db *gorm.DB,
) PartnerService {
	return &partnerService{
		gormDB: db,
	}
}

func (r *partnerService) AddPartner(ctx context.Context, partner model.Partner) error {
	return r.gormDB.Create(&partner).Error
}
