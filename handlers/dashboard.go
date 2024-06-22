package handlers

import (
	"PJ/E-Commerce_2/database"
	"PJ/E-Commerce_2/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ShowDashboard(c *gin.Context) {
	session := sessions.Default(c)

	name := session.Get("name")
	role := session.Get("role")
	var admin bool
	if role == "admin" {
		admin = true
	} else {
		admin = false
	}
	var data []models.Product

	err := database.Instance.Table("products").Select("*").Scan(&data).Error
	if err != nil {
		c.HTML(401, "notice.html", gin.H{
			"notice": "Fail to fetching data",
		})
		return
	}
	c.HTML(200, "dashboard.html", gin.H{
		"name":     name,
		"admin":    admin,
		"products": data,
	})
}
