package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"main/kernel"
	"main/sessions"
	"time"
)

const (
	SessionName    = "GOSESSID"
	CookieLifeTime = 3600
	CookiePath     = "/"
	CookieDomain   = "127.0.0.1"
	CookieHttpOnly = false
	CookieSecurity = false
)

func RedisSession(ctx *gin.Context) {
	cookie, err := ctx.Request.Cookie(SessionName)

	handler := &sessions.RedisSessionHandler{
		Client:  kernel.Redis,
		Context: context.Background(),
	}

	if err != nil {
		//if session id not set yet, generate a new session id.
		kernel.Session = sessions.New(handler, CookieLifeTime*time.Second)
		ctx.SetCookie(SessionName, kernel.Session.SessionId, CookieLifeTime, CookiePath, CookieDomain, CookieSecurity, CookieHttpOnly)
	} else {
		//if session id set already,
		kernel.Session = sessions.NewWithSessionId(handler, cookie.Value, CookieLifeTime*time.Second)
	}

	ctx.Next()
}
