package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

var products []Product

func initProducts() {

	bs, err := os.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(bs, &products); err != nil {
		log.Fatal(err)
	}
}

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func getProductsIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, product := range products {
		if product.ID == ID {
			err := json.NewEncoder(w).Encode(product)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			} //handles the encoding error and returns the correct status
			return

		} //if the product is equal to the ID encode and send the information,

	}

	w.WriteHeader(http.StatusNotFound)
	return //this means it is not in the slice at all
}

func createProductsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("objects/products.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var products []Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	price, _ := strconv.Atoi(r.FormValue("price"))
	if r.FormValue("price") == "" || r.FormValue("description") == "" || r.FormValue("name") == "" || r.FormValue("id") == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, product := range products {
		if product.ID == id {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	newProduct := Product{
		ID:          id,
		Price:       price,
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
	}

	products = append(products, newProduct)

	productsByte, _ := json.Marshal(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	file, err := os.Create("objects/products.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer file.Close()

	_, writeErr := file.Write(productsByte)
	if writeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func updateProductsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("objects/products.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var products []Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	price, err := strconv.Atoi(r.FormValue("price"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	if r.FormValue("price") == "" || r.FormValue("description") == "" || r.FormValue("name") == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newProduct := Product{
		ID:          id,
		Price:       price,
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
	}

	for i, product := range products {
		if product.ID == id {
			products[i] = newProduct

			productsByte, _ := json.Marshal(products)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			file, err := os.Create("objects/products.json")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			defer file.Close()

			_, writeErr := file.Write(productsByte)
			if writeErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	return
}

func deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("objects/products.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	var products []Product
	err = json.Unmarshal(data, &products)

	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)

			productsByte, _ := json.Marshal(products)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			file, err := os.Create("objects/products.json")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			defer file.Close()

			_, writeErr := file.Write(productsByte)
			if writeErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	return
}

func main() {
	initProducts()
	r := mux.NewRouter()
	r.HandleFunc("/products", getProductsHandler).Methods("GET")
	r.HandleFunc("/products/{id}", getProductsIdHandler).Methods("GET")
	r.HandleFunc("/products", createProductsHandler)
	r.HandleFunc("/products/{id}", updateProductsHandler).Methods("PUT")
	r.HandleFunc("products/{id}", deleteProductHandler).Methods("DELETE")
	// Bind to a port and pass our router in
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}
