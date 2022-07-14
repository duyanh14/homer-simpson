package usecase

import (
	"context"
	"simpson/internal/service"
)

type userUsecase struct {
	userService service.UserService
}

type UserUsecase interface {
	Register(ctx context.Context) error
}

func NewUserUsecase(
	userService service.UserService,
) UserUsecase {
	return &userUsecase{
		userService: userService,
	}
}

func (u *userUsecase) Register(ctx context.Context) error {
	err := u.userService.Register(ctx)
	if err != nil {
		return err
	}
	return nil
}
