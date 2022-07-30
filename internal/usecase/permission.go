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

type permissionUsecase struct {
	config            *config.Config
	permissionService service.PermissionService
}

type PermissionUsecase interface {
	AddPermission(ctx context.Context, req dto.AddPermissionReqDTO) error
	GetPermissions(ctx context.Context) ([]string, error)
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

func (u *permissionUsecase) GetPermissions(ctx context.Context) ([]string, error) {
	var (
		pers []string
		err  error
		log  = logger.GetLogger()
	)
	userID := ctx.Value("user_id").(uint)
	if userID == 0 {
		return pers, common.ErrUserIDNotFoundInJwt
	}
	persModel, err := u.permissionService.GetPermissionsUserId(ctx, userID)
	if err != nil {
		log.Errorf("get list permission of userID %d, error  %v", userID, err)
	}
	pers = make([]string, len(persModel))
	for i, permission := range persModel {
		pers[i] = permission.Code
	}
	return pers, err
}

func (u *permissionUsecase) AddPermission(ctx context.Context, req dto.AddPermissionReqDTO) error {
	var (
		log = logger.GetLogger()
	)
	req.Code = strings.TrimSpace(req.Code)
	if req.Code == "" {
		return common.ErrPermissionCodeRequire
	}
	// todo checking code exists
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
