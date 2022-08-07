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
	"strings"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type roleUsecase struct {
	config      *config.Config
	roleService service.RoleService
}

type RoleUsecase interface {
	AddRole(ctx context.Context, req dto.AddRoleReqDTO) error
	ListRole(ctx context.Context, req dto.ListRoleReqDTO) ([]dto.Role, error)
	UpdateRole(ctx context.Context, req dto.UpdateRoleReqDTO) error
	DeleteRole(ctx context.Context, req dto.DeleteRoleReqDTO) error
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
	req.Code = strings.TrimSpace(req.Code)
	if req.Code == "" {
		return common.ErrRoleCodeRequire
	}
	// todo checking code exists

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

func (u *roleUsecase) ListRole(ctx context.Context, req dto.ListRoleReqDTO) ([]dto.Role, error) {
	var (
		resp = []dto.Role{}
		err  error
		log  = logger.GetLogger()
	)
	rolesModel, err := u.roleService.ListRole(ctx)
	if err != nil {
		log.Errorf("get list role, error while call database error %v", err)
		return resp, common.ErrDatabase
	}
	resp = make([]dto.Role, len(rolesModel))
	for i, item := range rolesModel {
		roletemp := dto.Role{
			ID:          item.ID,
			Name:        item.Name,
			Code:        item.Code,
			Alias:       item.Alias,
			Description: item.Description,
			CreatedBy:   item.CreatedBy,
			UpdatedAt:   item.UpdatedAt,
			CreatedAt:   item.CreatedAt,
		}
		resp[i] = roletemp
	}
	return resp, err
}

func (u *roleUsecase) UpdateRole(ctx context.Context, req dto.UpdateRoleReqDTO) error {
	var (
		err         error
		log         = logger.GetLogger()
		rowaffected int64
	)
	// check role code is required
	req.Code = strings.TrimSpace(req.Code)
	if req.Code == "" {
		return common.ErrRoleCodeRequire
	}
	if req.RoleID == 0 {
		return common.ErrRoleIdRequire
	}
	userID := ctx.Value("user_id").(uint)
	rowaffected, err = u.roleService.UploadRole(ctx, model.Role{
		Model: gorm.Model{
			ID: req.RoleID,
		},
		Name:        req.Name,
		Alias:       req.Alias,
		Code:        req.Code,
		Description: req.Description,
		CreatedBy:   userID,
	})

	if err != nil {
		log.Errorf("update role, error while call database error %v", err)
		var perr *pgconn.PgError
		if errors.As(err, &perr) && perr.Code == common.DuplicateKeyValue {
			return common.ErrRoleCodeIsExists
		}
		return common.ErrDatabase
	}
	if rowaffected == 0 {
		return common.ErrRecordNotFound
	}
	return nil
}

type ErrorC struct {
	Code string
}

func (u *roleUsecase) DeleteRole(ctx context.Context, req dto.DeleteRoleReqDTO) error {
	var (
		err         error
		log         = logger.GetLogger()
		rowaffected int64
	)
	if req.RoleID == 0 {
		return common.ErrRoleIdRequire
	}

	rowaffected, err = u.roleService.DeleteRole(ctx, req.RoleID)
	if err != nil {
		log.Errorf("delete role, error while call database error %v", err)
		return common.ErrDatabase
	}
	if rowaffected == 0 {
		return common.ErrRecordNotFound
	}
	return nil
}

//
