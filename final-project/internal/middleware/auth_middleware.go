package middleware

import (
	"final-project/internal/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")

		if authorization == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Unauthorized",
				"data":    nil,
			})
			c.Abort()
			return
		}

		tokenString := strings.Split(authorization, "Bearer ")
		if len(tokenString) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Unauthorized",
				"data":    nil,
			})
			c.Abort()
			return
		}

		token := tokenString[1]
		verifyToken, err := helper.VerifyJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Unauthorized",
				"data":    nil,
			})
			c.Abort()
			return
		}

		c.Set("user", verifyToken)
		c.Next()
	}
}
