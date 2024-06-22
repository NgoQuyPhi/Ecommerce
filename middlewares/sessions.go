package middlewares

import (
	"PJ/E-Commerce_2/database"
	"PJ/E-Commerce_2/models"
	"net/http"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	// Bind the form data to a user struct
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var User models.User
	if err := c.ShouldBind(&input); err != nil {
		c.HTML(201, "notice.html", gin.H{
			"notice": err,
		})
		return
	}

	if input.Username == "" || input.Password == "" {
		c.HTML(401, "notice.html", gin.H{
			"notice": "username or password is invalid",
		})
		return
	}

	err := database.Instance.
		Table("users").
		Where("username = ?", input.Username).
		Scan(&User).Error

	if err != nil {
		// Handle error
		c.HTML(201, "notice.html", gin.H{
			"notice": err,
		})
		return
	}

	isAuthenticated := User.CheckPassword(input.Password)

	if isAuthenticated {
		session := sessions.Default(c)
		session.Set("authenticated", true)
		session.Set("role", User.Role)
		session.Set("username", User.Username)
		session.Set("name", User.Name)
		session.Set("id", User.UserId) // Assuming you have a user ID

		err = session.Save()
		if err != nil {
			// Handle error
			c.HTML(201, "notice.html", gin.H{
				"notice": err,
			})
			return
		}

		c.Redirect(http.StatusFound, "/api/dashboard")
	} else {
		c.HTML(401, "login.html", gin.H{
			"notice": "Invalid credentials",
		})
	}
}
