package middleware

import (
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		sessionUserID := session.Get("user_id")
		if sessionUserID == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Autharization failed",
			})
			ctx.Abort()
			
		}
		session.Set("Expires", time.Now().Add(24*time.Hour))
		session.Save()	
	}
}
