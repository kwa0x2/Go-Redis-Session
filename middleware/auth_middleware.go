package middleware

import (
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the default session for the current context
		session := sessions.Default(ctx)
		// Get the user ID from the session
		sessionUserID := session.Get("user_id")
		// If the user ID is nil (meaning no user is logged in), return an unauthorized status and message
		if sessionUserID == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization failed",
			})
			// Abort the request processing since the user is not authorized
			ctx.Abort()
		}
		// Set the expiration time for the session to 24 hours from now
		session.Set("Expires", time.Now().Add(24*time.Hour))
		// Save the session
		session.Save()
	}
}
