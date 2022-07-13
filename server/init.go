package server

import (
	"context"
	roleHttps "golang-course/service/domain/role/delivery/http"
	"golang-course/service/domain/role/usecase"

	"golang-course/service/repository"
)

type domains struct {
	role usecase.IRoleUsecase
}

func (s *Server) initDomains(
	ctx context.Context,
	repo repository.IRepo,
) *domains {

	role := usecase.NewRoleUsecase(repo)
	return &domains{
		role: role,
	}
}

func (s *Server) initRouters(domains *domains) {
	router := s.router.Group("v1/")
	roleHttps.NewRoleHandler(domains.role).RoleAPIRouter(router)
	// router.GET("add", roleHttps.AddRole)
}
