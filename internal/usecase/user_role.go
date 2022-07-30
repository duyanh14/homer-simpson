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

type userRoleUsecase struct {
	config          *config.Config
	userRoleService service.UserRoleService
	commonService   service.CommonService
	userService     service.UserService
	roleService     service.RoleService
}

type UserRoleUsecase interface {
	AddUserRole(ctx context.Context, req dto.AddUserRoleReqDTO) error
}

func NewUserRoleUsecase(
	config *config.Config,
	userRoleService service.UserRoleService,
	commonService service.CommonService,
	userService service.UserService,
	roleService service.RoleService,
) UserRoleUsecase {
	return &userRoleUsecase{
		config:          config,
		userRoleService: userRoleService,
		commonService:   commonService,
		userService:     userService,
		roleService:     roleService,
	}
}

func (u *userRoleUsecase) AddUserRole(ctx context.Context, req dto.AddUserRoleReqDTO) error {
	var (
		log = logger.GetLogger()
		err error
	)

	if len(req.RoleIds) == 0 {
		return errors.New("list role id not found")
	}

	if req.UserID == 0 {
		return errors.New("user id not found")
	}

	tx := u.commonService.GetDatabase(ctx).Begin()

	defer func() {
		if r := recover(); r != nil {
			log.Error("add user role, recover err %s", r)
			if err = tx.Rollback().Error; err != nil {
				log.Error("add user role, error when rollback transaction, err %s", err)
			}
		}
	}()

	_, err = u.userService.GetUserByID(ctx, tx, req.UserID)
	if err != nil {
		log.Error("add user role, get user detail %d, err %s", req.UserID, err)
		if err == gorm.ErrRecordNotFound {
			return common.ErrUserNotFound
		}
		return common.ErrDatabase
	}
	createdBy := ctx.Value("user_id").(uint)

	for _, roleID := range req.RoleIds {

		_, err = u.roleService.GetRoleByID(ctx, tx, roleID)
		if err != nil {
			log.Error("add user role, get role detail %d, err %s", roleID, err)
			if err == gorm.ErrRecordNotFound {
				return common.ErrRoleNotFound
			}
			return common.ErrDatabase
		}

		err = u.userRoleService.AddUserRole(ctx, tx, model.UserRole{
			Description: req.Description,
			RoleID:      roleID,
			UserID:      req.UserID,
			CreatedBy:   createdBy,
		})
		if err != nil {
			log.Error("add user role, error while add user %d , role  %d, err %s", req.UserID, roleID, err)
			return common.ErrDatabase
		}
	}
	if err != nil {
		log.Error("add user role, error when create for each user role, err %s", err)
		if err = tx.Rollback().Error; err != nil {
			log.Error("add user role, error when rollback transaction, err %s", err)
		}
		return common.ErrDatabase
	}
	err = tx.Commit().Error
	if err != nil {
		log.Error("add user role, commit transaction err %v", err)
		return common.ErrCommon
	}
	return nil

}
