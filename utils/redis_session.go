package utils

import (
	"os"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
)

func RedisSession() redis.Store {
	store, err := redis.NewStore(10,"tcp", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PASSWORD"), []byte(os.Getenv("SESSION_SECRET_KEY")))
	if err != nil {
		panic(err)
	}

	store.Options(sessions.Options{
		MaxAge: int((24 *time.Hour).Seconds()),
		Path: "/",
		HttpOnly: true,
		Secure: true,
	})

	return store
}