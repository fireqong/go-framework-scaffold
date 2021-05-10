package initialization

import (
	"main/module/api"
	"main/module/frontend"

	"github.com/gin-gonic/gin"
)

func App() *gin.Engine {
	app := gin.Default()

	//admin.InitRouter(app)
	frontend.InitRouter(app)
	api.InitRouter(app)

	//app.StaticFS("/public", gin.Dir("public", false))

	return app
}
