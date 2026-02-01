package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAll() ([]models.ProductPrint, error) {
	query := "SELECT product_id, product_name, price, stock, category_name FROM products INNER JOIN categories ON products.category_id = categories.category_id"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]models.ProductPrint, 0)
	for rows.Next() {
		var p models.ProductPrint
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryName)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (repo *ProductRepository) Create(product *models.Product) error {
	query := "INSERT INTO products (product_name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING product_id"
	err := repo.db.QueryRow(query, product.Name, product.Price, product.Stock, product.CategoryId).Scan(&product.ID)
	return err
}

// GetByID - ambil produk by ID
func (repo *ProductRepository) GetByID(id int) (*models.ProductPrint, error) {
	query := "SELECT product_id, product_name, price, stock, category_name FROM products INNER JOIN categories ON products.category_id = categories.category_id WHERE product_id = $1"

	var p models.ProductPrint
	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryName)
	if err == sql.ErrNoRows {
		return nil, errors.New("produk tidak ditemukan")
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (repo *ProductRepository) Update(product *models.Product) error {
	query := "UPDATE products SET product_name = $1, price = $2, stock = $3, category_id = $4 WHERE product_id = $5"
	result, err := repo.db.Exec(query, product.Name, product.Price, product.Stock, product.CategoryId, product.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return nil
}

func (repo *ProductRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE product_id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return err
}
