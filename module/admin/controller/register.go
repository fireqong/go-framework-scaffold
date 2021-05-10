package controller

import (
	"main/kernel"
	"main/kernel/util"
	"main/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func DoRegister(c *gin.Context) {
	admin := model.Admin{}
	if err := c.ShouldBind(&admin); err != nil {
		util.Response(c, gin.H{
			"err_msg": "Username or Password cannot be null.",
		})
		return
	}

	res := kernel.DB.Where(map[string]interface{}{"username": admin.Username}).First(&admin)

	if res.RowsAffected != 0 {
		util.Response(c, gin.H{
			"err_msg": "Name of User has been exists already.",
		})
		return
	}

	hashValue, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(err.Error())
	}

	admin.Password = string(hashValue)
	kernel.DB.Create(&admin)

	util.Response(c, gin.H{
		"ok_msg": "Register success.",
	})
}
