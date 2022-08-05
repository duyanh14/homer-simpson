package usecase

import (
	"context"
	"simpson/config"
	"simpson/internal/common"
	"simpson/internal/dto"
	"simpson/internal/helper/logger"
	"simpson/internal/service"
	"simpson/internal/service/model"
	"strings"
)

type accessUsecase struct {
	config        *config.Config
	accessService service.AccessService
}

type AccessUsecase interface {
	AddAccess(ctx context.Context, req dto.AddAccessReqDTO) error
}

func NewAccessUsecase(
	config *config.Config,
	accessService service.AccessService,
) AccessUsecase {
	return &accessUsecase{
		config:        config,
		accessService: accessService,
	}
}

func (u *accessUsecase) AddAccess(ctx context.Context, req dto.AddAccessReqDTO) error {
	var (
		log = logger.GetLogger()
	)
	req.Type = strings.TrimSpace(req.Type)
	if req.Type == "" {
		return common.ErrAccessTypeRequire
	}
	// todo checking code exists

	userID := ctx.Value("user_id").(uint)
	err := u.accessService.AddAccess(ctx, model.Access{
		Description: req.Description,
		Type:        req.Type,
		CreatedBy:   userID,
	})
	if err != nil {
		log.Errorf("add access, error while call database error %v", err)
		return common.ErrDatabase
	}
	return nil
}
