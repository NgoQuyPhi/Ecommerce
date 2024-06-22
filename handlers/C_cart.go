package handlers

import (
	"PJ/E-Commerce_2/database"
	"PJ/E-Commerce_2/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// Retrieve the product from the database
	var product models.Product
	err = database.Instance.
		Table("products").
		Select("product_id", "name", "price", "description").
		Where("product_id = ?", productID).
		Scan(&product).Error

	if err != nil {
		c.HTML(http.StatusInternalServerError, "notice.html", gin.H{
			"notice": err,
		})
		return
	}

	// Add the product to the cart (you'll need to manage the cart data in your application)
	cart := models.GetCartFromContext(c)
	cart.AddItem(product, 1)

	c.JSON(http.StatusOK, gin.H{"message": "Product added to cart"})
}
