package usecase

import (
	"context"
	"simpson/internal/common"
	"simpson/internal/dto"
	"simpson/internal/helper/logger"
	"simpson/internal/service"
	"simpson/internal/service/model"
)

type partnerUsecase struct {
	partnerService service.PartnerService
}

type PartnerUsecase interface {
	AddPartner(ctx context.Context, req dto.PartnerDTO) error
}

func NewPartnerUsecase(
	partnerService service.PartnerService,
) PartnerUsecase {
	return &partnerUsecase{
		partnerService: partnerService,
	}
}

func (u *partnerUsecase) AddPartner(ctx context.Context, req dto.PartnerDTO) error {
	var (
		log = logger.GetLogger()
	)
	err := u.partnerService.AddPartner(ctx, model.Partner{})
	if err != nil {
		log.Errorf("error while call database error %v", err)
		return common.ErrDatabase
	}
	return nil
}
