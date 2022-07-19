package usecase

import (
	"context"
	"errors"
	"simpson/config"
	"simpson/internal/common"
	"simpson/internal/dto"
	"simpson/internal/helper/logger"
	"simpson/internal/service"
	"simpson/internal/service/model"
	"simpson/internal/usecase/validation"

	"gorm.io/gorm"
)

type userUsecase struct {
	config      *config.Config
	userService service.UserService
	jwtUsecase  JwtUsecase
}

type UserUsecase interface {
	Register(ctx context.Context, req dto.UserDTO) error
	Verify(ctx context.Context, req dto.UserVerifyDTO) error
	Login(ctx context.Context, req dto.UserLoginReqDTO) (dto.UserLoginRespDTO, error)
}

func NewUserUsecase(
	config *config.Config,
	userService service.UserService,
	jwtUsecase JwtUsecase,
) UserUsecase {
	return &userUsecase{
		config:      config,
		userService: userService,
		jwtUsecase:  jwtUsecase,
	}
}

func (u *userUsecase) Register(ctx context.Context, req dto.UserDTO) error {
	var (
		log = logger.GetLogger()
		err error
	)

	if err = validation.ValidatorPw(req.Password); err != nil {
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

	pass, err := validation.HashPw(req.Password)
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

func (u *userUsecase) Login(ctx context.Context, req dto.UserLoginReqDTO) (dto.UserLoginRespDTO, error) {
	var (
		log  = logger.GetLogger()
		err  error
		resp = dto.UserLoginRespDTO{}
	)
	if req.Username == "" {
		return resp, errors.New("username is required")
	}
	if req.Password == "" {
		return resp, errors.New("username is required")
	}
	userDetail, err := u.userService.GetUserByUsername(ctx, req.Username)
	if err != nil {
		log.Error("get user username %s detail err %s", req.Username, err)
		if err == gorm.ErrRecordNotFound {
			return resp, errors.New("username not found")
		}
		return resp, common.ErrDatabase
	}
	err = validation.CheckPasswordHash(userDetail.Password, req.Password)
	if err != nil {
		log.Error("check password hash of username %s,err %s", req.Username, err)
		return resp, errors.New("password invalid")
	}
	resp.Jwt, err = u.jwtUsecase.GeneratorToken(ctx, dto.JwtReq{
		UserID:   userDetail.ID,
		Username: userDetail.Username,
		Email:    userDetail.Email,
		Phone:    userDetail.Phone,
	})
	if err != nil {
		log.Error("generator token jwt of username %s, err %s", req.Username, err)
		return resp, common.ErrCommon
	}
	return resp, err
}

func (u *userUsecase) Verify(ctx context.Context, req dto.UserVerifyDTO) error {
	log := logger.GetLogger()

	if req.Jwt == "" {

		return errors.New("token is required")
	}
	_, err := u.jwtUsecase.VerifyToken(ctx, req.Jwt)
	if err != nil {
		log.Error("verify jwt err %s", err)
		return err
	}
	return nil
}
