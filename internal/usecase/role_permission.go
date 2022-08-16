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
	"strconv"

	"gorm.io/gorm"
)

type rolePermissionUsecase struct {
	config                *config.Config
	rolePermissionService service.RolePermissionService
	commonService         service.CommonService
	permissionService     service.PermissionService
	roleService           service.RoleService
}

type RolePermissionUsecase interface {
	AddRolePermission(ctx context.Context, req dto.AddRolePermissionReqDTO) error
	GetListPermissionOfRole(ctx context.Context, req dto.GetListPermissionOfRole) ([]dto.Permission, error)
	GetListRoleOfPermission(ctx context.Context, req dto.GetListRoleOfPermission) ([]dto.Role, error)
}

func NewRolePermissionUsecase(
	config *config.Config,
	rolePermissionService service.RolePermissionService,
	commonService service.CommonService,
	permissionService service.PermissionService,
	roleService service.RoleService,
) RolePermissionUsecase {
	return &rolePermissionUsecase{
		config:                config,
		rolePermissionService: rolePermissionService,
		commonService:         commonService,
		permissionService:     permissionService,
		roleService:           roleService,
	}
}

func (u *rolePermissionUsecase) AddRolePermission(ctx context.Context, req dto.AddRolePermissionReqDTO) error {
	var (
		log = logger.GetLogger()
		err error
	)

	if req.PermissionID == 0 {
		return errors.New("permission id not found")
	}

	if req.RoleId == 0 {
		return errors.New("role id not found")
	}

	tx := u.commonService.GetDatabaseTx(ctx)

	defer func() {
		if r := recover(); r != nil {
			log.Error("add role permission, recover err %s", r)
			if err = tx.Rollback().Error; err != nil {
				log.Error("add role permission, error when rollback transaction, err %s", err)
			}
		}
		if err != nil {
			log.Error("add role permission, error when create for each role permission, err %s", err)
			if err = tx.Rollback().Error; err != nil {
				log.Error("add role permission, error when rollback transaction, err %s", err)
			}
		}
	}()

	_, err = u.roleService.GetRoleByID(ctx, tx, req.RoleId)
	if err != nil {
		log.Error("add role permission, get role detail %d, err %s", req.RoleId, err)
		if err == gorm.ErrRecordNotFound {
			return common.ErrRoleNotFound
		}
		return common.ErrDatabase
	}

	createdBy := ctx.Value("user_id").(uint)

	_, err = u.permissionService.GetPermissionByID(ctx, tx, req.PermissionID)
	if err != nil {
		log.Error("add role permission, get permisison detail %d, err %s", req.PermissionID, err)
		if err == gorm.ErrRecordNotFound {
			return common.ErrPermisisonNotFound
		}
		return common.ErrDatabase
	}
	err = u.rolePermissionService.AddRolePermission(ctx, tx, model.RolePermission{
		Description:  req.Description,
		RoleID:       req.RoleId,
		CreatedBy:    createdBy,
		PermissionID: req.PermissionID,
	})
	if err != nil {
		log.Error("add role permission, error while add permission %d , role  %d, err %s", req.PermissionID, req.RoleId, err)
		return common.ErrDatabase
	}

	err = tx.Commit().Error
	if err != nil {
		log.Error("add role permission, commit transaction err %v", err)
		return common.ErrCommon
	}
	return nil

}

func (u *rolePermissionUsecase) GetListPermissionOfRole(ctx context.Context,
	req dto.GetListPermissionOfRole) (listPermission []dto.Permission, err error) {
	log := logger.GetLogger()

	if req.RoleId == "" {
		return listPermission, common.ErrRoleIdRequire
	}
	roleID, err := (strconv.Atoi(req.RoleId))
	if err != nil {
		log.Error("get list permission of role, role id need number, err %v", err)
		return listPermission, common.ErrRoleIdNeedNumber
	}

	modelPermissions, err := u.rolePermissionService.GetPermissionsByRoleID(ctx, uint(roleID))
	if err != nil {
		log.Error("add list permission of role %s, error while call database err %s", req.RoleId, err)
		return nil, common.ErrDatabase
	}
	listPermission = make([]dto.Permission, len(modelPermissions))
	for i, item := range modelPermissions {
		listPermission[i] = dto.Permission{
			Name:        item.Name,
			Code:        item.Code,
			Description: item.Description,
			Alias:       item.Alias,
			CreatedAt:   item.CreatedAt,
			DeletedAt:   item.DeletedAt.Time,
			UpdatedAt:   item.UpdatedAt,
			ID:          item.ID,
			CreatedBy:   item.CreatedBy,
		}
	}
	return listPermission, nil

}

func (u *rolePermissionUsecase) GetListRoleOfPermission(ctx context.Context,
	req dto.GetListRoleOfPermission) (listRole []dto.Role, err error) {
	log := logger.GetLogger()

	if req.PermissionID == "" {
		return listRole, common.ErrRoleIdRequire
	}
	perID, err := (strconv.Atoi(req.PermissionID))
	if err != nil {
		log.Error("get list permission of role, role id need number, err %v", err)
		return listRole, common.ErrRoleIdNeedNumber
	}

	modelRoles, err := u.rolePermissionService.GetRolesByPermissionID(ctx, uint(perID))
	if err != nil {
		log.Error("add list role of permission %s, error while call database err %s", req.PermissionID, err)
		return nil, common.ErrDatabase
	}
	listRole = make([]dto.Role, len(modelRoles))
	for i, item := range modelRoles {
		listRole[i] = dto.Role{
			Name:        item.Name,
			Code:        item.Code,
			Description: item.Description,
			Alias:       item.Alias,
			CreatedAt:   item.CreatedAt,
			DeletedAt:   item.DeletedAt.Time,
			UpdatedAt:   item.UpdatedAt,
			ID:          item.ID,
			CreatedBy:   item.CreatedBy,
		}
	}
	return listRole, nil

}
