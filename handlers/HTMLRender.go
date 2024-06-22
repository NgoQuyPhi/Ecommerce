package handlers

import (
	"PJ/E-Commerce_2/database"
	"PJ/E-Commerce_2/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Addproduct(c *gin.Context) {
	c.HTML(200, "addproduct.html", nil)
}

func Updateproduct(c *gin.Context) {
	ProductID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.HTML(http.StatusBadRequest, "notice.html", gin.H{
			"notice": err,
		})
		return
	}

	var product models.Product

	err = database.Instance.
		Select("*").
		Table("products").
		Where("product_id = ?", ProductID).
		Scan(&product).Error

	if err != nil {
		c.HTML(http.StatusBadRequest, "notice.html", gin.H{
			"notice": err,
		})
		return
	}

	c.HTML(200, "updateproduct.html", gin.H{
		"product": product,
	})
}

func ShowReviewPage(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.HTML(http.StatusBadRequest, "notice.html", gin.H{
			"notice": err,
		})
		return
	}

	var AllReview []models.ShowReviews

	err = database.Instance.
		Table("reviews").
		Where("product_id = ? ", productID).
		Find(&AllReview).Error
	if err != nil {
		c.HTML(http.StatusBadRequest, "notice.html", gin.H{
			"notice": err,
		})
		return
	}

	c.HTML(200, "review.html", gin.H{
		"reviews":   AllReview,
		"productID": productID,
	})
}

func ShowCreateReviewPage(c *gin.Context) {
	c.HTML(200, "writereview.html", nil)
}

func WriteReview(c *gin.Context) {
	var review models.Reviews
	err := c.ShouldBind(&review)
	if err != nil {
		c.HTML(http.StatusBadRequest, "notice.html", gin.H{
			"notice": err,
		})
		return
	}

	err = database.Instance.Table("reviews").Create(&review).Error

	if err != nil {
		c.HTML(http.StatusBadRequest, "notice.html", gin.H{
			"notice": err,
		})
		return
	}

	c.HTML(200, "notice.html", gin.H{
		"notice": "Successful",
	})
}
