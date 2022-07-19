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

type roleUsecase struct {
	config      *config.Config
	roleService service.RoleService
}

type RoleUsecase interface {
	AddRole(ctx context.Context, req dto.AddRoleReqDTO) error
}

func NewRoleUsecase(
	config *config.Config,
	roleService service.RoleService,
) RoleUsecase {
	return &roleUsecase{
		config:      config,
		roleService: roleService,
	}
}

func (u *roleUsecase) AddRole(ctx context.Context, req dto.AddRoleReqDTO) error {
	var (
		log = logger.GetLogger()
	)
	userID := ctx.Value("user_id").(uint)
	err := u.roleService.AddRole(ctx, model.Role{
		Name:        req.Name,
		Alias:       req.Alias,
		Code:        req.Code,
		Description: req.Description,
		CreatedBy:   userID,
	})
	if err != nil {
		log.Errorf("add role, error while call database error %v", err)
		return common.ErrDatabase
	}
	return nil
}
