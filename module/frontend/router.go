package frontend

import (
	"github.com/gin-gonic/gin"
	"main/module/frontend/controller"
)

func InitRouter(router *gin.Engine) {
	frontendRouter := router.Group("/")
	frontendRouter.GET("/", controller.Index)
}
