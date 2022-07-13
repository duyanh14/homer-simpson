package api

import (
	"net/http"
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
	return func(ctx *gin.Context) {
		err := h.userUsecase.Register(ctx)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, nil)
	}
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
