package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiAuth(c *gin.Context) {
	if c.GetHeader("Authorization") == "353cb65361ed3bcea5a0043e15518914" {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusOK, map[string]string{
			"err_message": "非法调用",
		})
	}
}
