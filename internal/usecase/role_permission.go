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
