package usecase

import (
	"context"
	"errors"
	"simpson/internal/repository"
)

type Usecase struct {
	UserUsecase UserUsecase
}

func InitUsecase(ctx context.Context, repo repository.Repository) (*Usecase, error) {
	if repo == nil {
		return nil, errors.New("repo empty in usecase")
	}
	return &Usecase{
		UserUsecase: NewUserUsecase(repo.NewUserRepo()),
	}, nil
}
