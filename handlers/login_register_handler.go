package handlers

import (
	"PJ/E-Commerce_2/database"
	"PJ/E-Commerce_2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowLoginPage(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func ShowRegisterPage(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func RegisterHandlers(c *gin.Context) {

	var data models.User

	err := c.ShouldBind(&data)
	if err != nil {
		c.HTML(http.StatusBadRequest, "notice.html", gin.H{
			"notice": err,
		})
		return
	}

	err = database.Instance.Table("users").Create(&data).Error

	if err != nil {
		c.HTML(http.StatusBadRequest, "notice.html", gin.H{
			"notice": err,
		})
		return
	}

	c.Redirect(http.StatusFound, "/login")

}
