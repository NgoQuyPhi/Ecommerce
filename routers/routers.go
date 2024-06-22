package routers

import (
	"PJ/E-Commerce_2/handlers"
	"PJ/E-Commerce_2/middlewares"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Use(sessions.Sessions("mysession", cookie.NewStore([]byte("secret"))))
	router.GET("/login", handlers.ShowLoginPage)
	router.GET("/register", handlers.ShowRegisterPage)
	router.POST("/register", handlers.RegisterHandlers)
	router.POST("/login", middlewares.LoginHandler)
	router.GET("/checkout", handlers.Checkout_handlers)

	api := router.Group("/api")

	api.Use(authRequired)
	{
		api.GET("/dashboard", handlers.ShowDashboard)
		api.GET("/cart", handlers.GetCart)
		api.POST("/cart/:id", handlers.AddToCart)
		api.POST("/cart/remove/:id", handlers.RemoveFromCart)
		api.POST("/cart/add/:id", handlers.Add_item_quantity)
		api.POST("/cart/subtract/:id", handlers.Subtract_item_quantity)
		api.GET("/review/:id", handlers.ShowReviewPage)
		api.GET("/createreview/:id", handlers.ShowCreateReviewPage)
		api.POST("/review/:id", handlers.WriteReview)
	}

	admin := router.Group("/admin")

	admin.Use(adminRequired)
	{
		admin.GET("/addproduct", handlers.Addproduct)
		admin.POST("/addproduct", handlers.AddProduct)
		admin.GET("/edit/:id", handlers.EditHandler)
		admin.GET("/update/:id", handlers.Updateproduct)
		admin.PUT("/update/:id", handlers.UpdateProduct)
		admin.DELETE("/remove/:id", handlers.RemoveProduct)
	}

	return router
}

func authRequired(c *gin.Context) {
	session := sessions.Default(c)
	auth := session.Get("authenticated")

	if auth != true {
		c.Redirect(http.StatusFound, "/login")
	}
}

func adminRequired(c *gin.Context) {
	session := sessions.Default(c)
	role := session.Get("role")
	if role != "admin" {
		c.Redirect(302, "/login")
	}
}
