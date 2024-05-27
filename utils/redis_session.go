package utils

import (
	"os"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
)

func RedisSession() redis.Store {
	// Create a new Redis session store
	store, err := redis.NewStore(10, "tcp", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PASSWORD"), []byte(os.Getenv("SESSION_SECRET_KEY")))
	if err != nil {
		// Panic if there is an error creating the store
		panic(err)
	}

	// Configure the session options
	store.Options(sessions.Options{
		// Set the maximum age of the session to 24 hours
		MaxAge:   int((24 * time.Hour).Seconds()),
		// Set the path for the session cookie to "/"
		Path:     "/",
		// Set the HttpOnly flag to true to prevent client-side JavaScript access to the cookie
		HttpOnly: true,
		// Set the Secure flag to true to ensure the cookie is only sent over HTTPS
		Secure:   true,
	})

	// Return the configured store
	return store
}