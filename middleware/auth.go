package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/new-aspect/nexus-api/model"
	"strings"
)

// 思路，确认token是否合法、如果合法则继续，并将用户id写入context里，不合法则停止
// 确认是有用用户制定渠道，如果有也将渠道写入context里面
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, "-")
		key := parts[0]

		token, err := model.ValidateUseToken(key)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid Authorization token " + err.Error()})
			c.Abort()
			return
		}
		c.Set("id", token.UserId)

		if len(parts) > 1 {
			c.Set("channelId", parts[1])
		}

		c.Next()
	}
}
