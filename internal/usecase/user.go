package usecase

import (
	"context"
	"errors"
	"simpson/internal/common"
	"simpson/internal/dto"
	"simpson/internal/helper/logger"
	"simpson/internal/service"
	"simpson/internal/service/model"
	"simpson/internal/usecase/validation"
	"unicode"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userUsecase struct {
	userService service.UserService
}

type UserUsecase interface {
	Register(ctx context.Context, req dto.UserDTO) error
}

func NewUserUsecase(
	userService service.UserService,
) UserUsecase {
	return &userUsecase{
		userService: userService,
	}
}

func (u *userUsecase) Register(ctx context.Context, req dto.UserDTO) error {
	var (
		log = logger.GetLogger()
		err error
	)

	if err = u.validatorPw(req.Password); err != nil {
		log.Error("passwrd not security, err %s", err)
		return err
	}

	if req.Username == "" {
		return errors.New("username is required")
	}

	// checking user name exists by username
	_, err = u.userService.GetUserByUsername(ctx, req.Username)
	if err != gorm.ErrRecordNotFound {
		return errors.New("username is exist")
	}

	if req.Email != "" {
		if err = validation.ValidationEmail(req.Email); err != nil {
			log.Error("email err %s", err)
			return err
		}
		// checking user name exists by email
		_, err = u.userService.GetUserByEmail(ctx, req.Email)
		if err != gorm.ErrRecordNotFound {
			return errors.New("email is exist")
		}
	}

	if req.Phone != "" {
		if err = validation.ValidationPhone(req.Phone); err != nil {
			log.Error("phone err %s", err)
			return err
		}
		// checking user name exists by phone
		_, err = u.userService.GetUserByPhone(ctx, req.Phone)
		if err != gorm.ErrRecordNotFound {
			return errors.New("phone is exist")
		}
	}

	pass, err := hashPw(req.Password)
	if err != nil {
		log.Errorf("error while hass password error %v", err)
		return errors.New("hash password failed")
	}

	err = u.userService.Register(ctx, model.User{
		Username: req.Username,
		Phone:    req.Phone,
		Email:    req.Email,
		Password: pass,
	})
	if err != nil {
		log.Errorf("error while call database user register error %v", err)
		return common.ErrDatabase
	}
	return nil
}

/*
 * Password rules:
 * at least 7 letters
 * at max 20 letters
 * at least 1 number
 * at least 1 upper case
 * at least 1 special character
 */
func (u *userUsecase) validatorPw(pw string) error {
	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(pw) < 7 {
		return errors.New("password least 7 letters")
	}
	if len(pw) > 20 {
		return errors.New("password max 20 letters")
	}
	for _, charStr := range pw {
		switch {
		case unicode.IsUpper(charStr):
			hasUpper = true
		case unicode.IsLower(charStr):
			hasLower = true
		case unicode.IsNumber(charStr):
			hasNumber = true
		case unicode.IsPunct(charStr) || unicode.IsSymbol(charStr):
			hasSpecial = true
		}
	}
	if !hasUpper {
		return errors.New("password least 1 upper case")
	}
	if !hasNumber {
		return errors.New("password least 1 number")
	}
	if !hasLower {
		return errors.New("password least 1 lower case")
	}
	if !hasSpecial {
		return errors.New("password least 1 special character")
	}

	return nil
}

func hashPw(pw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), 14)
	return string(bytes), err
}

func checkPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
