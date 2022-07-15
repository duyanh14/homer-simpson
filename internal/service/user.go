package service

import (
	"context"
	"simpson/internal/service/model"

	"gorm.io/gorm"
)

type userService struct {
	gormDB *gorm.DB
}
type UserService interface {
	Register(ctx context.Context, user model.User) error
	GetUserByUsername(ctx context.Context, username string) (model.User, error)
	GetUserByPhone(ctx context.Context, phone string) (model.User, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
}

func NewUserService(
	db *gorm.DB,
) UserService {
	return &userService{
		gormDB: db,
	}
}

func (r *userService) Register(ctx context.Context, user model.User) error {
	err := r.gormDB.Create(&user).Error
	return err
}

func (r *userService) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	var (
		user model.User
		err  error
	)
	err = r.gormDB.Table(user.Table()).Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userService) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	var (
		user model.User
		err  error
	)
	err = r.gormDB.Table(user.Table()).Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userService) GetUserByPhone(ctx context.Context, phone string) (model.User, error) {
	var (
		user model.User
		err  error
	)
	err = r.gormDB.Table(user.Table()).Where("phone = ?", phone).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
