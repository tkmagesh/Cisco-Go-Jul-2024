package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Model
type Product struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Cost float64 `json:"cost"`
}

var products = []Product{
	{100, "Pen", 10},
	{101, "Pencil", 5},
	{102, "Marker", 50},
}

type AppServer struct {
}

func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v\t\t%v\n", r.Method, r.URL.Path)
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w, "Hello World!")
	case "/products":
		switch r.Method {
		case http.MethodGet:
			if err := json.NewEncoder(w).Encode(products); err != nil {
				http.Error(w, "data serialization error", http.StatusInternalServerError)
			}
		case http.MethodPost:
			var newProduct Product
			if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
				http.Error(w, "payload error", http.StatusBadRequest)
				return
			}
			products = append(products, newProduct)
			w.WriteHeader(http.StatusCreated)
			if err := json.NewEncoder(w).Encode(newProduct); err != nil {
				http.Error(w, "data serialization error", http.StatusInternalServerError)
			}
		default:
			http.Error(w, "method not supported", http.StatusMethodNotAllowed)
		}
	case "/customers":
		fmt.Fprintln(w, "customers list will be served!")
	default:
		http.Error(w, "resource not found!", http.StatusNotFound)
	}

}

func main() {
	appServer := AppServer{}
	http.ListenAndServe(":8080", &appServer)
}
