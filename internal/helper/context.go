package helper

import (
	"net/http"
	"simpson/internal/common"

	"github.com/gin-gonic/gin"
)

type ContextGin struct {
	*gin.Context
}
type responseData struct {
	StatusCode string      `json:"status"`
	Message    string      `json:"message,omitempty"`
	Code       string      `json:"code,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

type HandlerFunc func(ctx *ContextGin)

func WithContext(handler HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler(&ContextGin{
			ctx,
		})
	}
}

func (c *ContextGin) BadRequest(err error) {
	resp := responseData{
		StatusCode: common.Failed,
		Message:    common.ErrorMessage(err),
	}
	c.responseJson(http.StatusBadRequest, resp)
}

func (c *ContextGin) BadLogic(err error) {
	resp := responseData{
		StatusCode: common.Failed,
		Message:    common.ErrorMessage(err),
		Code:       common.ErrorCode(err),
	}
	// code 400 or 200
	c.responseJson(http.StatusBadRequest, resp)
}

func (c *ContextGin) OKResponse(data interface{}) {
	resp := responseData{
		StatusCode: common.OK,
	}
	if data != nil {
		resp.Data = data
	}
	c.responseJson(http.StatusOK, resp)
}

func (c *ContextGin) responseJson(code int, data interface{}) {
	c.JSON(code, data)
	if code != http.StatusOK {
		c.Abort()
	}
}
