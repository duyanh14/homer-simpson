package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"simpson/config"
	"simpson/internal/api"
	"simpson/internal/helper"
	"simpson/internal/helper/cache"
	"simpson/internal/helper/db"
	"simpson/internal/helper/logger"
	"simpson/internal/service"
	"simpson/internal/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//Server ...
type Server struct {
	httpServer *http.Server
	router     *gin.Engine
	cfg        *config.Config
}

// NewServer construct server
func NewServer(ctx context.Context) (*Server, error) {
	router := gin.New()
	s := &Server{
		router: router,
	}

	return s, nil
}

func (s *Server) initCors(ctx context.Context) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{
		"*",
		"Origin",
		"Content-Length",
		"Content-Type",
		"Authorization",
	}
	s.router.Use(cors.New(corsConfig))
}

// Init server
func (s *Server) Init(ctx context.Context) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	logger.Newlogger(cfg.Logger)

	// init database postgres
	dbPostgres, err := db.InitPostgres(cfg.Postgres)
	if err != nil {
		return err
	}

	// init cache redis
	cacheI, err := cache.NewRedisInstance(ctx, cfg.Redis)
	if err != nil {
		return err
	}
	fmt.Println(cacheI)

	s.cfg = cfg

	s.initCors(ctx)

	repo, err := service.InitService(ctx, dbPostgres)
	if err != nil {
		return err
	}

	usecase, err := usecase.InitUsecase(ctx, repo, cfg)
	if err != nil {
		return err
	}

	err = s.Router(usecase)
	if err != nil {
		return err
	}
	zap.L().Info("start server ok")
	return nil
}

//ListenHTTP ...
func (s *Server) ListenHTTP() error {

	address := fmt.Sprintf(":%d", s.cfg.HTTPAddress)

	s.httpServer = &http.Server{
		Handler: s.router,
		Addr:    address,
	}

	zap.S().Info("Start server at port %d", s.cfg.HTTPAddress)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Router(usecase *usecase.Usecase) error {
	if usecase == nil {
		return errors.New("router user nil")
	}
	router := s.router.Group("v1")

	router.Use(helper.AuthenticationJwt(usecase.JwtUsecase, s.cfg.IgnoreAuthen))
	// validatorIn := validation.InitValidator()
	//
	userRouter := api.NewUserHandler(usecase.UserUsecase)
	userRouter.UserRouter(router)

	roleRouter := api.NewRoleHandler()
	roleRouter.RoleRouter(router)

	permissionRouter := api.NewPermissionHandler(usecase.PermissionUsecase)
	permissionRouter.PermissionRouter(router)

	partnerRouter := api.NewPartnerHandler(usecase.PartnerUsecase, nil)
	partnerRouter.PartnerRouter(router)

	return nil
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			zap.S().Errorf("Recover when start project err:%s", err)
			os.Exit(0)
		}
	}()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	s, err := NewServer(ctx)
	if err != nil {
		panic(err)
	}
	err = s.Init(ctx)
	if err != nil {
		panic(err)
	}
	zap.S().Debug("Start project ok at port %s", s.cfg.HTTPAddress)
	if err := s.ListenHTTP(); err != nil {
		panic(err)
	}

}
