// data/cart_data.go
package models

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	cartData   = make(map[interface{}]*Cart)
	cartDataMu sync.Mutex
)

// Cart represents the shopping cart data structure
type Cart struct {
	Items []CartItem
	Total float32
}

// CartItem represents a product and its quantity in the cart
type CartItem struct {
	Product  Product
	Quantity int
}

// AddItem adds a product to the cart or updates the quantity if it already exists
func (c *Cart) AddItem(product Product, quantity int) {
	for i, item := range c.Items {
		if item.Product.ProductID == product.ProductID {
			c.Items[i].Quantity += quantity
			c.Total += product.Price * float32(quantity)
			if c.Items[i].Quantity == 0 {
				var newItems []CartItem
				for _, item := range c.Items {
					if item.Product.ProductID != product.ProductID {
						newItems = append(newItems, item)
					}
				}

				c.Items = newItems
				return
			}

			return
		}
	}
	c.Items = append(c.Items, CartItem{Product: product, Quantity: quantity})
	c.Total += product.Price * float32(quantity)
}

// GetCart retrieves the cart data for a given session ID
func GetCart(sessionID interface{}) *Cart {
	cartDataMu.Lock()
	defer cartDataMu.Unlock()
	cart, ok := cartData[sessionID]
	if !ok {
		cart = &Cart{
			Items: []CartItem{},
		}
		cartData[sessionID] = cart
	}
	return cart
}

func UpdateCart(sessionID string, cart *Cart) {
	cartDataMu.Lock()
	defer cartDataMu.Unlock()
	cartData[sessionID] = cart
}

func GetCartFromContext(c *gin.Context) *Cart {
	// Retrieve the session ID from the request context
	sessionID := getSessionID(c)

	// Get the cart data for the current session
	return GetCart(sessionID)
}

// getSessionID retrieves the session ID from the request context
// You'll need to implement this function based on your session management approach
func getSessionID(c *gin.Context) string {
	// Implement this function to retrieve the session ID from the request contexts

	return ""
}
