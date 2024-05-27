package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SetSessionExample(ctx *gin.Context) {
	session := sessions.Default(ctx)

	session.Set("user_id", uuid.New())
	session.Set("user_email", "example@nettasec.com")
	err := session.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":"Failed to save session",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":"Successfully logged in",
		"user_id":session.Get("user_id"),
		"user_email":session.Get("user_email"),
	})
}

func ClearSessionExample(ctx *gin.Context) {
	session := sessions.Default(ctx)

	session.Clear()
	session.Options(sessions.Options{MaxAge: -1})
	session.Save()

	ctx.SetCookie("connect.sid","",-1,"/","localhost",true,true)
	ctx.Redirect(http.StatusTemporaryRedirect, "/api/session/auth")
}

func AuthSessionExample(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":"If you are seeing this message, you are authenticated.",
	})
}