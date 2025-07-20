package middleware

import (
	"regexp"
	"strings"

	"github.com/cybercoder/restbill/pkg/utils"
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
		sub, err := utils.StringToUint(parts[0])
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "The user id is not unsigned int."})
			return
		}

		c.Set("sub", sub)
		c.Set("email", parts[1])
	}
}
