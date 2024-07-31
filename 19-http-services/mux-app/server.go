package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Cost float64 `json:"cost"`
}

var products = []Product{
	{101, "Pen", 10},
	{102, "Pencil", 5},
	{103, "Marker", 50},
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "error encoding data", http.StatusInternalServerError)
	}
}

func NewProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "error parsing payload", http.StatusBadRequest)
		return
	}
	products = append(products, newProduct)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newProduct); err != nil {
		http.Error(w, "error encoding data", http.StatusInternalServerError)
	}
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestedProductId := vars["id"]
	fmt.Printf("Product with id %q is requested", requestedProductId)
	fmt.Fprintf(w, "Product with id %q will be served!\n", requestedProductId)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", GetProductsHandler).Methods("GET")
	r.HandleFunc("/products", NewProductHandler).Methods("POST")
	r.HandleFunc("/products/{id:[0-9]+}", GetProductById).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
