package usecase

import (
	"context"
	"errors"
	"simpson/internal/service"
)

type Usecase struct {
	UserUsecase    UserUsecase
	PartnerUsecase PartnerUsecase
}

func InitUsecase(ctx context.Context, repo service.Service) (*Usecase, error) {
	if repo == nil {
		return nil, errors.New("repo empty in usecase")
	}
	return &Usecase{
		UserUsecase:    NewUserUsecase(repo.NewUserService()),
		PartnerUsecase: NewPartnerUsecase(repo.NewPartnerService()),
	}, nil
}
