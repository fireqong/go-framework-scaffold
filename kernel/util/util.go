package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, message gin.H) {
	c.JSON(http.StatusOK, message)
}

func ResponseErr(c *gin.Context, err error) {
	Response(c, gin.H{
		"err_message": err.Error(),
	})
}
