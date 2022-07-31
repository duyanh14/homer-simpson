package usecase

import (
	"context"
	"errors"
	"simpson/config"
	"simpson/internal/service"
)

type Usecase struct {
	UserUsecase           UserUsecase
	PartnerUsecase        PartnerUsecase
	JwtUsecase            JwtUsecase
	PermissionUsecase     PermissionUsecase
	RoleUsecase           RoleUsecase
	UserRoleUsecase       UserRoleUsecase
	RolePermissionUsecase RolePermissionUsecase
}

func InitUsecase(ctx context.Context, repo service.Service, cfg *config.Config) (*Usecase, error) {
	if repo == nil {
		return nil, errors.New("repo empty in usecase")
	}
	pri, pub, sign, err := ParseKey(cfg)
	if err != nil {
		return nil, err
	}
	jwtUsecase := NewJwtUsecase(cfg, pri, pub, sign)
	roleUsecase := NewRoleUsecase(cfg, repo.NewRoleService())
	userUsecase := NewUserUsecase(cfg, repo.NewUserService(), jwtUsecase, repo.NewPermissionService(), repo.NewRolePermisiosnService(), repo.NewUserRoleService())
	permissionUsecase := NewPermissionUsecase(cfg, repo.NewPermissionService(), repo.NewRolePermisiosnService(), repo.NewUserRoleService())
	userRoleUsecase := NewUserRoleUsecase(cfg, repo.NewUserRoleService(), repo.NewCommonService(), repo.NewUserService(), repo.NewRoleService())
	rolePerUsecase := NewRolePermissionUsecase(cfg, repo.NewRolePermisiosnService(), repo.NewCommonService(), repo.NewPermissionService(), repo.NewRoleService())

	return &Usecase{
		UserUsecase:           userUsecase,
		PartnerUsecase:        NewPartnerUsecase(repo.NewPartnerService()),
		JwtUsecase:            jwtUsecase,
		PermissionUsecase:     permissionUsecase,
		RoleUsecase:           roleUsecase,
		UserRoleUsecase:       userRoleUsecase,
		RolePermissionUsecase: rolePerUsecase,
	}, nil
}
