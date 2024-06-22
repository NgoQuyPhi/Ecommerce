package models

import "time"

type Product struct {
	ProductID      int         `json:"product_id"`
	Name           string      `json:"name"`
	Price          float32     `json:"price"`
	Description    string      `json:"description"`
	Stock_quantity int         `json:"stock_quantity"`
	CategoryID     int         `json:"category_id"`
	Create_at      *time.Timer `json:"create_at"`
	Update_at      *time.Timer `json:"update_at"`
}

type Reviews struct {
	ReviewID  int `json:"review_id"`
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	ShowReviews
}

type ShowReviews struct {
	Rating     int    `json:"rating"`
	ReviewText string `json:"review_text"`
}
