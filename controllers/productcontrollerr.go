package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Ycallaer/golang_crud_app/database"
	"github.com/Ycallaer/golang_crud_app/entities"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type handler struct {
    DB *gorm.DB
}

func New(db *gorm.DB) handler {
    return handler{db}
}


func (h handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product entities.Product
	json.NewDecoder(r.Body).Decode(&product)
	log.Println("Creating new database entry")
	h.DB.Create(&product)
	log.Println("Inserting db record")
	json.NewEncoder(w).Encode(product)
}

func (h handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if checkIfProductExists(productId) == false {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product entities.Product
	h.DB.First(&product, productId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting all the products")
	log.Println(h.DB.DB())
	var products []entities.Product
	h.DB.Find(&products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (h handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if checkIfProductExists(productId) == false {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product entities.Product
	h.DB.First(&product, productId)
	json.NewDecoder(r.Body).Decode(&product)
	h.DB.Save(&product)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h handler)DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	productId := mux.Vars(r)["id"]
	if checkIfProductExists(productId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product entities.Product
	h.DB.Delete(&product, productId)
	json.NewEncoder(w).Encode("Product Deleted Successfully!")
}

func checkIfProductExists(productId string) bool {
	var product entities.Product
	database.Db.First(&product, productId)

	if product.ID == 0 {
		return false
	}
	return true
}