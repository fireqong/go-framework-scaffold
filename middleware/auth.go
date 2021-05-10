package middleware

import (
	"main/kernel"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(ctx *gin.Context) {
	if kernel.Session.Has("is_login") && kernel.Session.Get("is_login") == "1" {
		ctx.Next()
	} else {
		ctx.AbortWithStatusJSON(http.StatusForbidden, map[string]string{
			"message": "please login first",
		})
	}
}
