package api

import (
	"fmt"
	"simpson/internal/helper"

	"github.com/gin-gonic/gin"
)

type healthRouter struct {
}

func NewHealthHandler() healthRouter {
	return healthRouter{}
}

func (h *healthRouter) liveness() gin.HandlerFunc {
	return helper.WithContext(func(ctx *helper.ContextGin) {
		fmt.Println("liveness")
	})
}

func (h *healthRouter) readiness() gin.HandlerFunc {
	return helper.WithContext(func(ctx *helper.ContextGin) {
		fmt.Println("readiness")
	})
}
