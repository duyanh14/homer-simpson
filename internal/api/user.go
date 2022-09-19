package api

import (
	"simpson/internal/dto"
	"simpson/internal/helper"
	"simpson/internal/helper/logger"
	"simpson/internal/helper/queue"
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

func (h *userRouter) login() gin.HandlerFunc {
	return helper.WithContext(func(ctx *helper.ContextGin) {
		var (
			log  = logger.GetLogger()
			req  = dto.UserLoginReqDTO{}
			resp = dto.UserLoginRespDTO{}
		)
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			log.Error("error while bind json %v", err)
			ctx.BadRequest(err)
			return
		}
		resp, err = h.userUsecase.Login(ctx, req)
		if err != nil {
			log.Error("error user login %w", err)
			ctx.BadLogic(err)
			return
		}
		ctx.OKResponse(resp)
	})
}

func (h *userRouter) verifyToken() gin.HandlerFunc {
	return helper.WithContext(func(ctx *helper.ContextGin) {
		var (
			req = dto.UserVerifyDTO{}
			log = logger.GetLogger()
		)
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			log.Error("error while bind json %v", err)
			ctx.BadRequest(err)
			return
		}
		err = h.userUsecase.Verify(ctx, req)
		if err != nil {
			log.Error("error user register %w", err)
			ctx.BadLogic(err)
			return
		}
		ctx.OKResponse(nil)
	})
}

// get list permission by userID(use jwt)
func (h *userRouter) listPermission() gin.HandlerFunc {
	return helper.WithContext(func(ctx *helper.ContextGin) {
		log := logger.GetLogger()
		pers, err := h.userUsecase.GetPermissions(ctx)
		if err != nil {
			log.Error("get list permission, error %w", err)
			ctx.BadLogic(err)
			return
		}
		ctx.OKResponse(pers)
	})
}

func (h *userRouter) checkAccess() gin.HandlerFunc {
	return helper.WithContext(func(ctx *helper.ContextGin) {
		var (
			req = dto.CheckAccessReqDTO{}
			log = logger.GetLogger()
		)
		// err := ctx.Query("code")
		// if err != nil {
		// 	log.Error("check access, error while bind json %v", err)
		// 	ctx.BadRequest(err)
		// 	return
		// }
		req.PermissionCode = ctx.Query("code")
		resp, err := h.userUsecase.CheckAccess(ctx, req)
		if err != nil {
			log.Error("check access, error %w", err)
			ctx.BadLogic(err)
			return
		}
		log.Info("checking access success", resp.IsAccess)
		ctx.OKResponse(resp)
	})

}

func (h *userRouter) userInfo() gin.HandlerFunc {
	return helper.WithContext(func(ctx *helper.ContextGin) {
		var (
			req = dto.UserInfoReqDTO{}
			log = logger.GetLogger()
		)
		resp, err := h.userUsecase.UserInfo(ctx, req)
		if err != nil {
			log.Error("get user info, error %w", err)
			ctx.BadLogic(err)
			return
		}
		ctx.OKResponse(resp)
	})
}

////

type Loc struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type InC struct {
	ID       int    `json:"id"`
	CodeName string `json:"codeName"`
	Loc      Loc    `json:"loc"`
	OffcerID int    `json:"offcerId"`
}

type Officers struct {
	ID        int    `json:"id"`
	Loc       Loc    `json:"loc"`
	BadgeName string `json:"badgeName"`
}

type Data struct {
	In       []InC         `json:"incidents"`
	Officers []Officers    `json:"officers"`
	Lisst    queue.Dequeue `json:"list"`
}

func (h *userRouter) locationMap() gin.HandlerFunc {
	return helper.WithContext(func(ctx *helper.ContextGin) {

		inc := []InC{}
		inc = append(inc, InC{
			ID:       1,
			CodeName: "dsf",
			Loc: Loc{
				X: 1,
				Y: 1,
			},
			OffcerID: 1,
		})

		off := []Officers{}
		off = append(off, Officers{
			ID: 1,
			Loc: Loc{
				X: 1,
				Y: 1,
			},
			BadgeName: "sdfds",
		})

		resp := Data{
			In:       inc,
			Officers: off,
		}

		ctx.OKResponse(resp)
	})
}
