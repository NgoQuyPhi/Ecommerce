package handlers

import (
	"PJ/E-Commerce_2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCart(c *gin.Context) {
	cart := models.GetCartFromContext(c)

	c.HTML(http.StatusOK, "cart.html", gin.H{
		"cart": cart,
	})
}
