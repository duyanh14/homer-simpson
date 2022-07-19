package helper

import (
	"net/http"
	"simpson/internal/common"
	"simpson/internal/usecase"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	tokenParttern = "^Bearer (\\S*)"
)

func ignoreAuthen(arr []string, method string) bool {
	for _, item := range arr {
		if strings.Contains(method, item) {
			return true
		}
	}
	return false
}

func AuthenticationJwt(jwt usecase.JwtUsecase, listIgnore []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ignoreAuthen(listIgnore, ctx.Request.RequestURI) {
			ctx.Next()
			return
		}

		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, ResponseData{
				StatusCode: common.Failed,
				Message:    "Authorization not found",
			})
			ctx.Abort()
			return
		}
		splitToken := strings.Split(token, " ")
		if len(splitToken) != 2 {
			ctx.JSON(http.StatusUnauthorized, ResponseData{
				StatusCode: common.Failed,
				Message:    "Authorization not found",
			})
			ctx.Abort()
			return
		}
		claim, err := jwt.VerifyToken(ctx, splitToken[1])
		if err != nil {
			messageStr := ""
			if err == common.ErrTokenExpired {
				messageStr = err.Error()
			}
			if err == common.ErrTokenInvalid {
				messageStr = err.Error()
			}
			ctx.JSON(http.StatusUnauthorized, ResponseData{
				StatusCode: common.Failed,
				Message:    messageStr,
			})
			ctx.Abort()
			return
		}
		ctx.Set("user_id", claim.UserID)
		ctx.Set("user_name", claim.Username)
		ctx.Next()
	}
}
