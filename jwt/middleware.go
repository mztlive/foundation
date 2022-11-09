package jwt

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mztlive/foundation/response"
)

// JwtAuthMiddleware Token认证中间件
func JwtAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusForbidden, response.ForbiddenRequest())
			c.Abort()
			return
		}
		claim, err := DeToken(token)
		if err != nil {
			c.JSON(http.StatusForbidden, response.ForbiddenRequest())
			c.Abort()
			return
		}
		c.Set("userID", claim.UserID)
		c.Next()
	}
}
