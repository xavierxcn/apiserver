package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xavierxcn/apiserver/internal/serve/handler"
	"github.com/xavierxcn/apiserver/internal/serve/pkg/errno"
	"github.com/xavierxcn/apiserver/pkg/token"
)

// AuthMiddleware 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
