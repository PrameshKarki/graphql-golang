package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/PrameshKarki/event-management-golang/utils"
	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization != "" {
			parts := strings.Split(authorization, "Bearer")
			token := strings.Trim(parts[1], " ")
			data, err := utils.VerifyToken(token, utils.GoDotEnv("TOKEN_SECRET"))
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "message": "Unauthorized!"})
				c.Abort()
			} else {
				ctx := context.WithValue(c.Request.Context(), "user", data)
				c.Request = c.Request.WithContext(ctx)

			}
		}
		c.Next()
	}
}
