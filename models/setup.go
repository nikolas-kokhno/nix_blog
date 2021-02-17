package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func ConnectToDB(dbDriver, dbUser, dbPassword, dbPort, dbHost, dbName string) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)
	database, err := gorm.Open(dbDriver, connectionString)
	if err != nil {
		log.Fatalf("Cannot connect to %s database: %v", dbDriver, err)
	}
	log.Printf("We are connected to %s database", dbDriver)

	database.AutoMigrate(&Users{}, &Posts{}, &Comments{})

	DB = database
}

func InitDatabase() {
	ConnectToDB(
		"postgres",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_DB"))
}
