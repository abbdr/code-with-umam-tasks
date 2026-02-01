package models

type Product struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
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
