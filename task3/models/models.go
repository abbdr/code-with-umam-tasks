package models

import "time"

type Product struct {
	ID         int    `json:"id"`
	Name       string `json:"product_name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryId int    `json:"category_id"`
}

type ProductPrint struct {
	ID           int    `json:"id"`
	Name         string `json:"product_name"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
	CategoryName string `json:"category_name"`
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"category_name"`
	Description string `json:"description"`
}

type Transaction struct {
	ID          int                 `json:"id"`
	TotalAmount int                 `json:"total_amount"`
	CreatedAt   time.Time           `json:"created_at"`
	Details     []TransactionDetail `json:"details"`
}

type TransactionDetail struct {
	ID            int    `json:"id"`
	TransactionID int    `json:"transaction_id"`
	ProductID     int    `json:"product_id"`
	ProductName   string `json:"product_name,omitempty"`
	Quantity      int    `json:"quantity"`
	Subtotal      int    `json:"subtotal"`
}

type CheckoutItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CheckoutRequest struct {
	Items []CheckoutItem `json:"items"`
}

type Report struct {
	TransactionReport []Transaction	`json:"report"`
}