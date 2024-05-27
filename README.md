[![Go](https://github.com/kwa0x2/Go-Redis-Session/actions/workflows/go.yml/badge.svg)](https://github.com/kwa0x2/Go-Redis-Session/actions/workflows/go.yml)


# Go Redis Session API Example

This project includes Docker installation for Redis and uses gin-contrib/sessions for Redis session management. The project contains functionality for setting session, clearing session, and authentication middleware.

## Installation:

**1. Start the services using Docker Compose:**

```bash
docker-compose up -d
```

**2. Install the Go dependencies:**

```bash
go mod download
```

**3. Start the project with Air for hot reloading:**

```bash
air
```

## Examples:

### Set Session

```go
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
```

### Clear Session

```go
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

```

### Auth Middleware

```go
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
```
