package api

import (
	"net/http"
	"simpson/internal/dto"
	"simpson/internal/helper"
	"simpson/internal/helper/logger"
	"simpson/internal/usecase"

	"github.com/gin-gonic/gin"
)

type userRouter struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(
	userUsecase usecase.UserUsecase,
) userRouter {
	return userRouter{
		userUsecase: userUsecase,
	}
}

func (h *userRouter) register() gin.HandlerFunc {
	return helper.WithContext(func(ctx *helper.ContextGin) {
		var (
			req = dto.UserDTO{}
			log = logger.GetLogger()
		)
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			log.Error("error while bind json %v", err)
			ctx.BadRequest(err)
			return
		}
		err = h.userUsecase.Register(ctx, req)
		if err != nil {
			log.Error("error user register %w", err)
			ctx.BadLogic(err)
			return
		}
		ctx.OKResponse(nil)
	})
}

type abc struct {
	Name string `json:"name"`
}

func (h *userRouter) login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, abc{
			Name: "ducnp",
		})

	}
}
