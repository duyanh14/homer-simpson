package ginwrapper

import "github.com/gin-gonic/gin"

type HandlerFunc func(ctx *Context)

// Context wrapper of gin.Context
type Context struct {
	*gin.Context
}

func WithContext(handler HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		wrappedContext := &Context{
			ctx,
		}
		handler(wrappedContext)
	}
}

type Response struct {
	Status        string `json:"status"`
	ReasonCode    uint8  `json:"reason_code"`
	ReasonMessage string `json:"reason_message"`
}

func (r Response) ResponseOK() Response {
	return Response{
		Status:        "OK",
		ReasonCode:    200,
		ReasonMessage: "",
	}
}
