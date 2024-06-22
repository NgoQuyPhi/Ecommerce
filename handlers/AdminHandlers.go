package handlers

import (
	"PJ/E-Commerce_2/database"
	"PJ/E-Commerce_2/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddProduct(c *gin.Context) {
	var products models.Product

	err := c.ShouldBind(&products)

	if err != nil {
		c.HTML(http.StatusBadRequest, "notice.html", gin.H{
			"notice": err,
		})
		return
	}

	err = database.Instance.Table("products").Create(&products).Error

	if err != nil {
		c.HTML(http.StatusBadRequest, "notice.html", gin.H{
			"notice": err,
		})
		return
	}

	c.HTML(200, "notice.html", gin.H{
		"notice": "Successfully",
	})

}

func EditHandler(c *gin.Context) {
	ProductID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "notice.html", gin.H{
			"notice": err,
		})
		return
	}

	var product models.Product

	err = database.Instance.
		Table("products").
		Select("*").
		Where("product_id = ?", ProductID).
		Scan(&product).Error
	if err != nil {
		c.HTML(http.StatusBadRequest, "notice.html", gin.H{
			"notice": err,
		})
		return
	}
	c.HTML(200, "edit.html", gin.H{
		"product": product,
	})
}

func UpdateProduct(c *gin.Context) {
	var data models.Product

	err := c.ShouldBind(&data)

	if err != nil {
		c.HTML(http.StatusBadRequest, "notice.html", gin.H{
			"notice": err,
		})
		return
	}

	productID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.HTML(http.StatusBadRequest, "notice.html", gin.H{
			"notice": err,
		})
		return
	}

	err = database.Instance.
		Table("products").
		Where("product_id = ?", productID).
		Updates(&data).Error

	if err != nil {
		c.HTML(http.StatusBadRequest, "notice.html", gin.H{
			"notice": err,
		})
		return
	}

	c.HTML(200, "notice.html", gin.H{
		"notice": "Successfully",
	})
}

func RemoveProduct(c *gin.Context) {
	ProductID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.HTML(http.StatusBadRequest, "notice.html", gin.H{
			"notice": err,
		})
		return
	}

	err = database.Instance.Table("products").Where("product_id = ?", ProductID).Delete(nil).Error

	if err != nil {
		c.HTML(http.StatusBadRequest, "notice.html", gin.H{
			"notice": err,
		})
		return
	}

	c.HTML(200, "notice.html", gin.H{
		"notice": "Successfully",
	})
}
