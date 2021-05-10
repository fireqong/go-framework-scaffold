package admin

import (
	"fmt"
	"main/kernel"
	"main/middleware"
	"main/module/admin/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	adminRouter := router.Group("/admin")

	adminRouter.Use(middleware.RedisSession)

	adminRouter.POST("/login", controller.DoLogin)
	adminRouter.POST("/register", controller.DoRegister)

	authRouter := router.Group("/admin")
	authRouter.Use(middleware.RedisSession, middleware.Auth)

	authRouter.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%#v", struct {
				name string
				age  int
			}{name: "church", age: 20}),
		})
	})

	authRouter.GET("/test", func(context *gin.Context) {
		kernel.Session.Set("is_login", "2")
	})
}
