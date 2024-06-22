package handlers

import (
	"PJ/E-Commerce_2/database"
	"PJ/E-Commerce_2/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Add_item_quantity(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	sessionID := getSessionID(c)
	cart := models.GetCart(sessionID)

	var product models.Product
	err = database.Instance.
		Select("product_id", "price").
		Table("products").
		Where("product_id = ?", productID).Scan(&product).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	cart.AddItem(product, 1)
	models.UpdateCart(sessionID, cart)

	c.Redirect(http.StatusFound, "/api/cart")
}
