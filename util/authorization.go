package util

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthByTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		expectedToken := os.Getenv("API_TOKEN")
		actualToken := c.GetHeader("X-Api-Token")

		// 空チェック
		if actualToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		// 前後の空白を除去して比較
		if strings.TrimSpace(actualToken) != strings.TrimSpace(expectedToken) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		c.Next()
	}
}
