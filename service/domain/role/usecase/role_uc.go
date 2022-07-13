package usecase

import (
	"context"
	"golang-course/service/models"
	"golang-course/service/models/dto"
	"golang-course/service/repository"
	"strconv"
)

type IRoleUsecase interface {
	GetListRole(ctx context.Context) (listRole []*dto.RoleResponseDTO, err error)
	AddRole(ctx context.Context, req dto.RoleRequestDTO) error
	DeleteRole(ctx context.Context, req dto.RoleDeleteRequestDTO) error
	UpdateRole(ctx context.Context, req dto.RoleUpdateRequestDTO) error
}

type roleUsecase struct {
	repo repository.IRepo
}

func NewRoleUsecase(
	repo repository.IRepo,

) IRoleUsecase {
	return &roleUsecase{
		repo: repo,
	}
}

func (c *roleUsecase) GetListRole(ctx context.Context) (listRole []*dto.RoleResponseDTO, err error) {
	list, err := c.repo.NewRoleRepo().GetRole(ctx)
	if err != nil {
		return listRole, err
	}
	listRole = make([]*dto.RoleResponseDTO, len(list))
	for i, role := range list {
		listRole[i] = &dto.RoleResponseDTO{
			RoleID:          role.RoleID,
			RoleName:        role.RoleName,
			RoleAlias:       role.RoleAlias.String,
			RoleDescription: role.RoleDescription.String,
			CreatedBy:       role.CreatedBy.String,
			CreatedDate:     role.CreatedDate,
			IsActive:        role.IsActive.Bool,
		}
	}
	return listRole, nil
}

func (c *roleUsecase) AddRole(ctx context.Context, req dto.RoleRequestDTO) error {
	err := c.repo.NewRoleRepo().AddRole(ctx, models.AddRole{
		RoleName:        req.RoleName,
		RoleAlias:       req.RoleAlias,
		RoleDescription: req.RoleDescription,
		CreatedBy:       req.CreatedBy,
	})
	return err
}

func (c *roleUsecase) DeleteRole(ctx context.Context, req dto.RoleDeleteRequestDTO) error {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		return err
	}
	return c.repo.NewRoleRepo().DeleteRole(ctx, id)
}

func (c *roleUsecase) UpdateRole(ctx context.Context, req dto.RoleUpdateRequestDTO) error {
	return c.repo.NewRoleRepo().UpdateRole(ctx, models.UpdateRole{
		RoleName:        req.RoleName,
		RoleAlias:       req.RoleAlias,
		RoleDescription: req.RoleDescription,
		RoleID:          req.ID,
	})
}
