package handlers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Checkout_handlers(c *gin.Context) {
	session := sessions.Default(c)

	session.Set("authenticated", false)
	session.Set("role", nil)
	session.Set("username", nil)
	session.Set("name", nil)
	session.Set("id", nil)
	c.Redirect(302, "/login")
}
