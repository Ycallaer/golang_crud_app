package main

import (
	"fmt"
	"github.com/Ycallaer/golang_crud_app/controllers"
	"github.com/Ycallaer/golang_crud_app/database"
	"github.com/Ycallaer/golang_crud_app/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"

)

func main() {

	// Load Configurations from config.json using Viper
	config.LoadAppConfig()

	// Initialize Database
	Db := database.Init(config.AppConfig.ConnectionString)

	
	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	h := controllers.New(Db)

	// Register Routes
	//RegisterProductRoutes(router)
	router.HandleFunc("/api/products", h.GetProducts).Methods(http.MethodGet)
	router.HandleFunc("/api/products/{id}", h.GetProductById).Methods(http.MethodGet)
	router.HandleFunc("/api/products", h.CreateProduct).Methods(http.MethodPost)
	router.HandleFunc("/api/products/{id}", h.UpdateProduct).Methods(http.MethodPut)
	router.HandleFunc("/api/products/{id}", h.DeleteProduct).Methods(http.MethodDelete)
	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", config.AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.AppConfig.Port), router))
}
