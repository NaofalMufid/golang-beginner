package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var (
	products     []Product
	productsLock sync.Mutex
	PORT         = ":5050"
)

func main() {
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getProducts(w, r)
		case http.MethodPost:
			createProduct(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getProductID(w, r)
		case http.MethodPut:
			updateProduct(w, r)
		case http.MethodDelete:
			deleteProduct(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server running on port", PORT)
	http.ListenAndServe(PORT, nil)
}

// controller
func createProduct(w http.ResponseWriter, r *http.Request) {
	var new_product Product
	if err := json.NewDecoder(r.Body).Decode(&new_product); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	productsLock.Lock()
	defer productsLock.Unlock()

	new_product.ID = len(products) + 1
	products = append(products, new_product)

	w.WriteHeader(http.StatusCreated)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	productsLock.Lock()
	defer productsLock.Unlock()

	productJSON, err := json.Marshal(products)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(productJSON)
}

func getProductID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/products/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	productsLock.Lock()
	defer productsLock.Unlock()

	for _, product := range products {
		if product.ID == id {
			productJSON, err := json.Marshal(product)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(productJSON)
			return
		}
	}

	http.Error(w, "Not Found", http.StatusNotFound)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/products/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var update_product Product
	if err := json.NewDecoder(r.Body).Decode(&update_product); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	productsLock.Lock()
	defer productsLock.Unlock()

	for i := range products {
		if products[i].ID == id {
			update_product.ID = id
			products[i] = update_product
			return
		}
	}

	http.Error(w, "Not Found", http.StatusNotFound)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/products/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	productsLock.Lock()
	defer productsLock.Unlock()

	for i, product := range products {
		if product.ID == id {
			// hapus product
			products = append(products[:i], products[i+1:]...)
			return
		}
	}

	http.Error(w, "Not Found", http.StatusNotFound)
}
