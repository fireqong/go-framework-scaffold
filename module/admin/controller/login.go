package controller

import (
	"main/kernel"
	"main/kernel/util"
	"main/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func DoLogin(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")

	if username == "" || password == "" {
		util.Response(c, gin.H{
			"err_msg": "Username or Password cannot be empty.",
		})
		return
	}

	admin := model.Admin{}
	kernel.DB.Where("username = ?", username).First(&admin)

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		util.Response(c, gin.H{
			"err_msg": "Username or Password incorrect.",
		})
		return
	}

	kernel.Session.Set("is_login", "1")

	util.Response(c, gin.H{
		"ok_msg": "Auth success.",
	})

}
