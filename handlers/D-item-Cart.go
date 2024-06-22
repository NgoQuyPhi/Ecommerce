package handlers

import (
	"PJ/E-Commerce_2/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RemoveFromCart(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	sessionID := getSessionID(c)
	cart := models.GetCart(sessionID)

	for _, item := range cart.Items {
		if item.Product.ProductID == productID {
			quantity := 0 - item.Quantity
			cart.AddItem(item.Product, quantity)
		}
	}

	models.UpdateCart(sessionID, cart)

	c.Redirect(http.StatusFound, "/api/cart")
}

func getSessionID(c *gin.Context) string {
	// Implement this function to retrieve the session ID from the request context
	return ""
}
