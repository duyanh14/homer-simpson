package helper

import (
	"fmt"
	"simpson/internal/helper/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SetRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get("request-id")
		if requestID == "" {
			requestID = uuid.NewString()
		}
		fmt.Println("requestID", requestID)
		logger.GetLogger().SetLogginID(requestID)
		c.Set("request-id", requestID)
		c.Writer.Header().Set("request-id", requestID)
	}
}
