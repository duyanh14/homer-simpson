package repository

import (
	"context"

	"gorm.io/gorm"
)

type userRepo struct {
	gormDB *gorm.DB
}
type UserRepo interface {
	Register(ctx context.Context) error
}

func NewUserRepo(
	db *gorm.DB,
) UserRepo {
	return &userRepo{
		gormDB: db,
	}
}

func (r *userRepo) Register(ctx context.Context) error {
	return nil
}
