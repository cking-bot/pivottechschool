package main

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Connection struct {
	DB     *gorm.DB
	Router *mux.Router
}
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

var products []Product

func initConnection(filepath string) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Created new Database")
	return db, nil
}

func (c *Connection) getProductsHandler(w http.ResponseWriter, r *http.Request) {
	limiter := r.URL.Query().Get("limit")
	if limiter == "" {
		limiter = "10"
	}
	sortBy := r.URL.Query().Get("sort")
	if sortBy == "" {
		sortBy = "id"
	}

	limitInteger, _ := strconv.Atoi(limiter)
	if limitInteger > 100 {
		limitInteger = 100
	}
	if limitInteger < 0 {
		limitInteger = 10
	}
	switch sortBy {
	case "id":
	case "name":
	case "description":
	case "price":
	default:
		sortBy = "id"

	}

	var prodRet []Product
	c.DB.Table("products").Select("*").Limit(limitInteger).Order(sortBy).Scan(&prodRet)
	log.Println(prodRet)

	rep, err := json.Marshal(&prodRet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(rep)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *Connection) getProductsIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]

	if _, err := strconv.Atoi(ID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var prodRet Product
	c.DB.Table("products").First(&prodRet, "ID = ?", ID).Scan(&prodRet)

	if prodRet.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	rep, err := json.Marshal(&prodRet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(rep)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *Connection) createProductsHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if product.ID == 0 || product.Name == "" || product.Price == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var prodRet Product
	c.DB.Table("products").First(&prodRet, "ID = ?", product.ID).Scan(&prodRet)
	if prodRet.ID == product.ID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c.DB.Create(&product)
	w.WriteHeader(http.StatusCreated)
}

func (c *Connection) updateProductsByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]

	if _, err := strconv.Atoi(ID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if product.ID == 0 || product.Name == "" || product.Price == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var prodRet Product
	c.DB.Table("products").First(&prodRet, "ID = ?", ID).Scan(&prodRet)
	if prodRet.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	idInteger, err := strconv.Atoi(ID)
	if idInteger == product.ID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c.DB.Model(&products).Where("id = ?").Updates(Product{ID: product.ID, Name: product.Name, Description: product.Description, Price: product.Price})
	w.WriteHeader(http.StatusOK)
}

func (c *Connection) deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]

	if _, err := strconv.Atoi(ID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idInteger, err := strconv.Atoi(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if idInteger == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var check Product //looks for product
	c.DB.Table("products").First(&check, "ID = ?", ID).Scan(&check)
	if check.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	c.DB.Where("ID = ?", ID).Delete(&Product{}) //deletes

	var prodRet Product //checks to make sure it was deleted
	c.DB.Table("products").First(&prodRet, "ID = ?", ID).Scan(prodRet)
	if prodRet.ID == 0 {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func main() {
	var err error
	c := Connection{}
	c.Router = mux.NewRouter()

	c.DB, err = initConnection("../../data-driven-product-server/seeder/products.db")
	if err != nil {
		log.Printf("Can't connect to Database\n%v", err)
		os.Exit(1)
	}

	c.Router.HandleFunc("/products", c.getProductsHandler).Methods("GET")
	c.Router.HandleFunc("/products", c.createProductsHandler).Methods("POST")
	c.Router.HandleFunc("/products/{id}", c.getProductsIdHandler).Methods("GET")
	c.Router.HandleFunc("/products/{id}", c.updateProductsByIdHandler).Methods("PUT")
	c.Router.HandleFunc("products/{id}", c.deleteProductHandler).Methods("DELETE")
	// Bind to a port and pass our router in
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", c.Router))

}
