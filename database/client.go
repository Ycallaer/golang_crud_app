package database

import (
	"github.com/Ycallaer/golang_crud_app/entities"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect(connectionString string) (*gorm.DB, error){

	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
	return db, err
}

func Migrate() {
	db.AutoMigrate(&entities.Product{})
	log.Println("Database Migration Completed...")
}