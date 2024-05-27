package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/go-redis-session/controller"
	"github.com/kwa0x2/go-redis-session/middleware"
	"github.com/kwa0x2/go-redis-session/utils"
)

func main() {
	utils.LoadEnv()
	router := gin.New()
	store := utils.RedisSession()
	// The name of your cookie is "connect.sid"
	router.Use(sessions.Sessions("connect.sid", store))

	sessionRoutes:=router.Group("/api/session")
	sessionRoutes.GET("set", controller.SetSessionExample)
	sessionRoutes.GET("clear", controller.ClearSessionExample)
	sessionRoutes.GET("auth",middleware.SessionAuthMiddleware(), controller.AuthSessionExample)

	router.Run(":9000")
}