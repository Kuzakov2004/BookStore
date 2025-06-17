package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AdminAuthRequired is a middleware that checks if the user has a valid session.
// It should be used on routes that require authentication.
// If no valid session exists, it aborts the request with 401 Unauthorized.
func AdminAuthRequired(c *gin.Context) {
	// Get the session from the request context
	session := sessions.Default(c)

	// Try to get the user from the session
	if user := session.Get("auid"); user == nil {
		// No user in session, abort the request
		c.Redirect(http.StatusFound, "/admin/login")
		return
	}

	// User is authenticated, continue to the next handler
	c.Next()
}

func ClientAuthRequired(c *gin.Context) {
	// Get the session from the request context
	session := sessions.Default(c)

	// Try to get the user from the session
	if user := session.Get("cuid"); user == nil {
		// No user in session, abort the request
		c.Redirect(http.StatusFound, "/admin/login")
		return
	}

	// User is authenticated, continue to the next handler
	c.Next()
}

func SetFlags(c *gin.Context) {
	// Get the session from the request context
	session := sessions.Default(c)

	// Try to get the user from the session
	if user := session.Get("cuid"); user != nil {
		c.Keys["isClient"] = true
	}

	if user := session.Get("auid"); user != nil {
		c.Keys["isAdmin"] = true
	}
	// User is authenticated, continue to the next handler
	c.Next()
}
