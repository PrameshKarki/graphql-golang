package auth

import (
	"context"

	"github.com/gin-gonic/gin"
)

type contextKey string

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		UID := "TEST"
		ctx := context.WithValue(c.Request.Context(), contextKey("user_id"), UID)
		c.Request = c.Request.WithContext(ctx)
		c.Set("user_id", UID)
		c.Next()
	}
}
