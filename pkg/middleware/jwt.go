package middleware

import (
	"log"
	"net/http"
	"strings"
	"time"
	"web3-music-platform/pkg/utils"

	"github.com/gin-gonic/gin"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		token = strings.Split(token, " ")[1]
		if token == "" {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "token not found",
			})
			return
		}
		claims, err := utils.ParseToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "parse token fail" + err.Error(),
			})
			c.Abort()
			return
		}

		if time.Now().Unix() > claims.ExpiresAt {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Permission expired, please log in again",
			})
			c.Abort()
			return
		}
		log.Print("claims==>", claims.Address)
		c.Set("user_addr", claims.Address)
		userAddr, exists := c.Get("user_addr")
		log.Print("user_addr==>", userAddr, "exists==>", exists)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": "set user info fail" + err.Error(),
			})
			return
		}
		c.Next()
	}
}
