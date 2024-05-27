package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// SetSessionExample sets a user ID and email in the session and saves it.
func SetSessionExample(ctx *gin.Context) {
	session := sessions.Default(ctx)

	// Generate a new UUID for the user and set it in the session
	session.Set("user_id", uuid.New().String())
	session.Set("user_email", "example@nettasec.com")
	err := session.Save()
	if err != nil {
		// Return an internal server error if the session could not be saved
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":err.Error(),
		})
		return
	}

	// Return a success message with the user ID and email from the session
	ctx.JSON(http.StatusOK, gin.H{
		"message":"Successfully logged in",
		"user_id":session.Get("user_id"),
		"user_email":session.Get("user_email"),
	})
}

// ClearSessionExample clears the session and deletes the session cookie.
func ClearSessionExample(ctx *gin.Context) {
	session := sessions.Default(ctx)

	// Clear all session data and set the session's max age to -1 (delete it)
	session.Clear()
	session.Options(sessions.Options{MaxAge: -1})
	session.Save()

	// Delete the session cookie
	ctx.SetCookie("connect.sid","",-1,"/","localhost",true,true)

	// Return a success message indicating the session has been cleared
	ctx.JSON(http.StatusOK, gin.H{
		"message":"Successfully cleared. Please go to the authentication endpoint for testing.",
	})
}

// AuthSessionExample returns a message indicating the user is authenticated.
func AuthSessionExample(ctx *gin.Context) {
	// If the request reaches this handler, it means the user has passed through the authentication middleware.
	
	// Return a success message indicating the user is authenticated.
	ctx.JSON(http.StatusOK, gin.H{
		"message": "If you are seeing this message, you are authenticated.",
	})
}
