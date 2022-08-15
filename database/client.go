package database

import (
	"github.com/Ycallaer/golang_crud_app/entities"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var Db * gorm.DB

func Init(connectionString string)(*gorm.DB){
	log.Println("Connecting to database %s",connectionString)
	Db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil && Db != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
	Db.AutoMigrate(&entities.Product{})
	log.Println("Done auto migration")
	return Db
}