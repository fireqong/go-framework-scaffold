package api

import (
	"main/middleware"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	apiRouter := router.Group("api")
	apiRouter.Use(location.Default(), middleware.ApiAuth)
}
