package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AdminAuthRequired is a middleware that checks if the user has a valid session.
// It should be used on routes that require authentication.
// If no valid session exists, it aborts the request with 401 Unauthorized.
func AdminAuthRequired(c *gin.Context) {
	// Get the session from the request context
	session := sessions.Default(c)

	// Try to get the user from the session
	if user := session.Get("asid"); user == nil {
		// No user in session, abort the request
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// User is authenticated, continue to the next handler
	c.Next()
}

// login is a handler that parses a form and checks for specific data.
func login(c *gin.Context) {
	session := sessions.Default(c)

	login := c.PostForm("login")
	password := c.PostForm("password")

	// Validate form input
	if strings.Trim(login, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check for username and password match, usually from a database
	if login != "hello" || password != "itsme" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// Save the username in the session
	session.Set("auid", auid) // In real world usage you'd set this to the users ID
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}

// logout is the handler called for the user to log out.
func logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete(userkey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

/*
v := session.Get("count")
if v == nil {
count = 0
} else {
count = v.(int)
count++
}
session.Set("count", count)
session.Save()
*/
