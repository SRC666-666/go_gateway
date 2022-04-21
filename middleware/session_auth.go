package middleware

import (
	"errors"

	"github.com/e421083458/gateway_demo/public"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if AdminInfo, ok := session.Get(public.AdminSessionInfoKey).(string); !ok || AdminInfo == "" {
			ResponseError(c, InternalErrorCode, errors.New("user not login"))
			c.Abort()
			return
		}
		c.Next()
	}
}
