package usecase

import (
	"context"
	"simpson/internal/repository"
)

type userUsecase struct {
	userRepo repository.UserRepo
}

type UserUsecase interface {
	Register(ctx context.Context) error
}

func NewUserUsecase(
	userRepo repository.UserRepo,
) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) Register(ctx context.Context) error {
	err := u.userRepo.Register(ctx)
	if err != nil {
		return err
	}
	return nil
}
