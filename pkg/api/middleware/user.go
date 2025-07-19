package middleware

import (
	"regexp"
	"strings"

	"github.com/cybercoder/restbill/pkg/logger"
	"github.com/gin-gonic/gin"
)

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userHeader := c.GetHeader("X-Remote-User")
		if userHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		userHeader = regexp.MustCompile(`.*#`).ReplaceAllString(userHeader, "")
		parts := strings.Split(userHeader, "/")
		logger.Infof("sub %s email %s", parts[0], parts[1])
		c.Set("sub", parts[0])
		c.Set("email", parts[1])
	}
}
