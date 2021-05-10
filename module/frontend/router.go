package frontend

import (
	"main/module/frontend/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	frontendRouter := router.Group("/")
	frontendRouter.GET("/", controller.Index)
}
