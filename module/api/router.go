package api

import (
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"main/middleware"
)

func InitRouter(router *gin.Engine) {
	apiRouter := router.Group("api")
	apiRouter.Use(location.Default(), middleware.ApiAuth)
}
