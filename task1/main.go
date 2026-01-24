package main

// go get -tool github.com/air-verse/air@latest

import (
	"encoding/json"
	"fmt"
	"log"
	"maps"
	"net/http"
	"strconv"
	"strings"
)

// Produk represents a product in the cashier system
type Product struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
	Stok  int    `json:"stok"`
}

// Cetgory represents a category in the cashier system
type Category struct {
	ID    				int    `json:"id"`
	Name  				string `json:"name"`
	Description 	string `json:"description"`
}




// In-memory storage (sementara, nanti ganti database)
var product = map[int]Product{
	1: {ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 10},
	2: {ID: 2, Nama: "Vit 1000ml", Harga: 3000, Stok: 40},
	3: {ID: 3, Nama: "kecap", Harga: 12000, Stok: 20},
}

// In-memory storage (sementara, nanti ganti database)
var category = map[int]Category{
	1: {ID: 1, Name: "Fashion", Description: "Fashion Category",},
	2: {ID: 2, Name: "Sport", Description: "Sport category",},
	3: {ID: 3, Name: "Kitchen", Description: "Kitchen Category",},
}




func main() {

	// /health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, " ", r.URL.Path, " from ", r.RemoteAddr)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	// GET /api/product/{id}
	// PUT /api/product/{id}
	// DELETE /api/product/{id}
	http.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, " ", r.URL.Path, " from ", r.RemoteAddr)
		if r.Method == "GET" {
			getProductByID(w, r)
		} else if r.Method == "PUT" {
			updateProduct(w, r)
		} else if r.Method == "DELETE" {
			deleteProduct(w, r)
		}
	})

	// GET /api/categories/{id}
	// PUT /api/categories/{id}
	// DELETE /api/categories/{id}
	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, " ", r.URL.Path, " from ", r.RemoteAddr)
		if r.Method == "GET" {
			getCategoryByID(w, r)
		} else if r.Method == "PUT" {
			updateCategory(w, r)
		} else if r.Method == "DELETE" {
			deleteCategory(w, r)
		}
	})

	// GET /api/products
	// POST /api/products
	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, " ", r.URL.Path, " from ", r.RemoteAddr)
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			fmt.Println(maps.Values(product))
			json.NewEncoder(w).Encode(maps.Values(product))
			
		} else if r.Method == "POST" {
			// baca data dari request
			var newProduct Product
			err := json.NewDecoder(r.Body).Decode(&newProduct)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}
			if newProduct.Harga == 0 || newProduct.Nama == "" || newProduct.Stok == 0 {
				http.Error(w, "Value of Harga, Nama, or Stok can't be 0 or empty string", http.StatusBadRequest)
				return
			}
			log.Println(newProduct)

			// masukkin data ke dalam variable product
			productLength := len(product)
			for i := 1; ;i++ {
				id := productLength+i
				_, exist := product[id]
				if exist {
					continue
				}else{
					newProduct.ID = id
					break
				}
			}
			product[newProduct.ID] = newProduct

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated) // 201
			json.NewEncoder(w).Encode(newProduct)
		}
	})

	// GET /api/categories
	// POST /api/categories
	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, " ", r.URL.Path, " from ", r.RemoteAddr)
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			categorySlice := []Category{}
			for v := range maps.Values(category){
				categorySlice = append(categorySlice, v)
			}

			json.NewEncoder(w).Encode(categorySlice)
			
		} else if r.Method == "POST" {
			// baca data dari request
			var newCategory Category
			err := json.NewDecoder(r.Body).Decode(&newCategory)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}
			if newCategory.Description == "" || newCategory.Name == "" {
				http.Error(w, "Value of Name or Description can't be empty string", http.StatusBadRequest)
				return
			}
			log.Println(newCategory)

			// masukkin data ke dalam variable produk
			categoryLength := len(category)
			for i := 1; ;i++ {
				id := categoryLength+i
				_, exist := category[id]
				if exist {
					continue
				}else{
					newCategory.ID = id
					break
				}
			}
			category[newCategory.ID] = newCategory

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated) // 201
			json.NewEncoder(w).Encode(newCategory)
		}
	})

	fmt.Println("Server running di localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("gagal running server")
	}

}



// GET product by id /api/product/{id}
func getProductByID(w http.ResponseWriter, r *http.Request) {
	// Parse ID dari URL path
	// URL: /api/product/123 -> ID = 123
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	// Cari produk dengan ID tersebut
	_, exist := product[id]
	fmt.Println(exist)
	if !exist {
		// Kalau not found
		http.Error(w, "Produk belum ada", http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product[id])
}

// PUT /api/product/{id}
func updateProduct(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")

	// ganti int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	// get data dari request
	var updateProduct Product
	err = json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if updateProduct.Harga == 0 || updateProduct.Nama == "" || updateProduct.Stok == 0 {
		http.Error(w, "Value of Harga, Nama, or Stok can't be 0 or empty string", http.StatusBadRequest)
		return
	}

	// cari id, ganti sesuai data dari request
	_, exist := product[id]
	if !exist {
		http.Error(w, "Produk belum ada", http.StatusNotFound)
		return
	}

	updateProduct.ID = id
	product[id] = updateProduct
	log.Println(updateProduct)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateProduct)
	
}

// DELETE /api/project/{id}
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	// get id
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	
	// ganti id int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}
	
	// cari id yang mau dihapus
	_, exist := product[id]
	if !exist {	
		http.Error(w, "Produk belum ada", http.StatusNotFound)
		return
	}

	delete(product, id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "sukses delete",
	})
}



// GET category by id /api/categories/{id}
func getCategoryByID(w http.ResponseWriter, r *http.Request) {
	// Parse ID dari URL path
	// URL: /api/categories/123 -> ID = 123
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	// Cari category dengan ID tersebut
	_, exist := category[id]
	if !exist {
		// Kalau not found
		http.Error(w, "Category belum ada", http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product[id])
}

// PUT /api/categories/{id}
func updateCategory(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")

	// ganti int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	// get data dari request
	var updateCategory Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if updateCategory.Description == "" || updateCategory.Name == "" {
		http.Error(w, "Value of Name or Description can't be empty string", http.StatusBadRequest)
		return
	}

	// cari id, ganti sesuai data dari request
	_, exist := category[id]
	if !exist {
		http.Error(w, "Category belum ada", http.StatusNotFound)
		return
	}
	
	updateCategory.ID = id
	category[id] = updateCategory
	log.Println(updateCategory)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateCategory)
}

// DELETE /api/categories/{id}
func deleteCategory(w http.ResponseWriter, r *http.Request) {
	// get id
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	
	// ganti id int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}
	
	// cari id yang mau dihapus
	_, exist := category[id]
	if !exist {
		http.Error(w, "Category belum ada", http.StatusNotFound)
		return
	}
	
	// hapus data dg key id
	delete(category, id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "sukses delete",
	})
}
