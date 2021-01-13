package kernel

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"main/sessions"
)

var (
	DB      *gorm.DB
	App     *gin.Engine
	Session *sessions.Session
	Redis   *redis.Client
)
