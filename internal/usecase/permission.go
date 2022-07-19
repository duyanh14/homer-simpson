package usecase

import (
	"context"
	"simpson/config"
	"simpson/internal/common"
	"simpson/internal/dto"
	"simpson/internal/helper/logger"
	"simpson/internal/service"
	"simpson/internal/service/model"
)

type permissionUsecase struct {
	config            *config.Config
	permissionService service.PermissionService
}

type PermissionUsecase interface {
	AddPermission(ctx context.Context, req dto.AddPermissionReqDTO) error
}

func NewPermissionUsecase(
	config *config.Config,
	permissionService service.PermissionService,
) PermissionUsecase {
	return &permissionUsecase{
		config:            config,
		permissionService: permissionService,
	}
}

func (u *permissionUsecase) AddPermission(ctx context.Context, req dto.AddPermissionReqDTO) error {
	var (
		log = logger.GetLogger()
	)
	userID := ctx.Value("user_id").(uint)
	err := u.permissionService.AddPermission(ctx, model.Permission{
		Name:        req.Name,
		Alias:       req.Alias,
		Code:        req.Code,
		Description: req.Description,
		CreatedBy:   userID,
	})
	if err != nil {
		log.Errorf("add permission, error while call database error %v", err)
		return common.ErrDatabase
	}
	return nil
}
