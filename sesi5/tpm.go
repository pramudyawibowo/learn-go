package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var products []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"status": "success",
		"data":   products,
	}
	json.NewEncoder(w).Encode(response)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	products = append(products, product)
	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"status": "success",
		"data":   product,
	}
	json.NewEncoder(w).Encode(response)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var updatedProduct Product
	_ = json.NewDecoder(r.Body).Decode(&updatedProduct)

	var product Product
	var found bool
	for i, p := range products {
		if p.ID == id {
			products[i].Name = updatedProduct.Name
			products[i].Price = updatedProduct.Price
			product = products[i]
			found = true
			break
		}
	}
	if !found {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"status": "success",
		"data":   product,
	}
	json.NewEncoder(w).Encode(response)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() // Menutup body request

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var found bool
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"status": "success",
		"data":   nil,
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	products = append(products, Product{1, "Produk 1", 100})
	products = append(products, Product{2, "Produk 2", 200})

	http.HandleFunc("GET /products", getProducts)
	http.HandleFunc("POST /products", createProduct)
	http.HandleFunc("PUT /products", updateProduct)
	http.HandleFunc("DELETE /products", deleteProduct)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
