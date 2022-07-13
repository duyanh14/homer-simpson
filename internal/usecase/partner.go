package usecase

import (
	"context"
	"simpson/internal/dto"
	"simpson/internal/repository"
	"simpson/internal/repository/model"
)

type partnerUsecase struct {
	partnerRepo repository.PartnerRepo
}

type PartnerUsecase interface {
	AddPartner(ctx context.Context, req dto.PartnerDTO) error
}

func NewPartnerUsecase(
	partnerRepo repository.PartnerRepo,
) PartnerUsecase {
	return &partnerUsecase{
		partnerRepo: partnerRepo,
	}
}

func (u *partnerUsecase) AddPartner(ctx context.Context, req dto.PartnerDTO) error {
	// if err := validator.AddPartnerValidator(req); err != nil {
	// 	return err
	// }
	err := u.partnerRepo.AddPartner(ctx, model.Partner{})
	if err != nil {
		// TODO logging
		//	return common.DatabaseError
	}
	return nil
}
