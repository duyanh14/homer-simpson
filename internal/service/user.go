package service

import (
	"context"

	"gorm.io/gorm"
)

type userService struct {
	gormDB *gorm.DB
}
type UserService interface {
	Register(ctx context.Context) error
}

func NewUserService(
	db *gorm.DB,
) UserService {
	return &userService{
		gormDB: db,
	}
}

func (r *userService) Register(ctx context.Context) error {
	return nil
}
