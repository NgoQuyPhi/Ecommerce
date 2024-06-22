package handlers

import (
	"PJ/E-Commerce_2/database"
	"PJ/E-Commerce_2/models"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var data []models.Product

	err := database.Instance.Table("products").Select("*").Scan(&data).Error
	if err != nil {
		c.HTML(401, "notice.html", gin.H{
			"notice": "Fail to fetching data",
		})
		return
	}

	c.HTML(200, "products.html", gin.H{
		"products": data,
	})
}
