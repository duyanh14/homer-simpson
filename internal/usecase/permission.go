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

type permissionUsecase struct {
	config                *config.Config
	permissionService     service.PermissionService
	rolePermissionService service.RolePermissionService
	userRoleService       service.UserRoleService
}

type PermissionUsecase interface {
	AddPermission(ctx context.Context, req dto.AddPermissionReqDTO) error
	GetPermissions(ctx context.Context) ([]string, error)

	ListPermission(ctx context.Context, req dto.ListPermissionReqDTO) ([]dto.Permission, error)
	UpdateRole(ctx context.Context, req dto.UpdatePermissionReqDTO) error
	DeletePermission(ctx context.Context, req dto.DeletePermissionReqDTO) error
}

func NewPermissionUsecase(
	config *config.Config,
	permissionService service.PermissionService,
	rolePermissionService service.RolePermissionService,
	userRoleService service.UserRoleService,
) PermissionUsecase {
	return &permissionUsecase{
		config:                config,
		permissionService:     permissionService,
		rolePermissionService: rolePermissionService,
		userRoleService:       userRoleService,
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

	roles, err := u.userRoleService.GetRolesByUserID(ctx, userID)
	if err != nil {
		log.Errorf("get list role of userID %d, error  %v", userID, err)
		return pers, common.ErrDatabase
	}
	if len(roles) == 0 {
		log.Errorf("get list permission, roles of userID %d not found ", userID)
		return pers, common.ErrPermisisonNotFound
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
		var perr *pgconn.PgError
		if errors.As(err, &perr) && perr.Code == common.DuplicateKeyValue {
			return common.ErrPermissionCodeIsExists
		}
		return common.ErrDatabase
	}
	return nil
}

func (u *permissionUsecase) ListPermission(ctx context.Context, req dto.ListPermissionReqDTO) ([]dto.Permission, error) {
	var (
		resp = []dto.Permission{}
		err  error
		log  = logger.GetLogger()
	)
	listPer, err := u.permissionService.ListPermission(ctx)
	if err != nil {
		log.Errorf("get list permission, error while call database error %v", err)
		return resp, common.ErrDatabase
	}
	resp = make([]dto.Permission, len(listPer))
	for i, item := range listPer {
		roletemp := dto.Permission{
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

func (u *permissionUsecase) UpdateRole(ctx context.Context, req dto.UpdatePermissionReqDTO) error {
	var (
		err         error
		log         = logger.GetLogger()
		rowaffected int64
	)
	// check role code is required
	req.Code = strings.TrimSpace(req.Code)
	if req.Code == "" {
		return common.ErrPermissionCodeRequire
	}
	if req.PermissionID == 0 {
		return common.ErrPermissionIdRequire
	}
	userID := ctx.Value("user_id").(uint)
	rowaffected, err = u.permissionService.UploadPermission(ctx, model.Permission{
		Model: gorm.Model{
			ID: req.PermissionID,
		},
		Name:        req.Name,
		Alias:       req.Alias,
		Code:        req.Code,
		Description: req.Description,
		CreatedBy:   userID,
	})
	if err != nil {
		log.Errorf("update permission, error while call database error %v", err)
		var perr *pgconn.PgError
		if errors.As(err, &perr) && perr.Code == common.DuplicateKeyValue {
			return common.ErrPermissionCodeIsExists
		}
		return common.ErrDatabase
	}
	if rowaffected == 0 {
		return common.ErrRecordNotFound
	}
	return nil
}

func (u *permissionUsecase) DeletePermission(ctx context.Context, req dto.DeletePermissionReqDTO) error {
	var (
		err         error
		log         = logger.GetLogger()
		rowaffected int64
	)
	if req.PermissionID == 0 {
		return common.ErrPermissionIdRequire
	}

	rowaffected, err = u.permissionService.DeletePermission(ctx, req.PermissionID)
	if err != nil {
		log.Errorf("delete permission, error while call database error %v", err)
		return common.ErrDatabase
	}
	if rowaffected == 0 {
		return common.ErrRecordNotFound
	}
	return nil
}
