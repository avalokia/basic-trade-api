package database

import (
	"basic-trade-api/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// Get DB credentials from .env file and connect the DB
func StartDB() {
	// Load .env file
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalln("Error loading .env file:", err)
	// 	// return err
	// }
	// Get DB credentials from .env file
	host := os.Getenv("HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")
	// Define connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, dbport)
	// Open connection with postgreSQL
	fmt.Println("dsn:", dsn)
	fmt.Println("Opening connection with PSQL...")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed connecting DB:", err)
		// log.Fatalln("Failed connecting DB:", err)
	}

	// Auto migration
	fmt.Println("Auto migrating DB...")
	err = db.AutoMigrate(models.Admin{}, models.Product{}, models.Variant{})
	if err != nil {
		log.Fatalln("Failed to auto migrate:", err)
	}
}

// Get DB that has been started by StartDB()
func GetDB() *gorm.DB {
	// If db hasn't been started, start db
	if db == nil {
		fmt.Println("DB is nil. Starting DB..")
		StartDB()
		// RECONNECT DB
		// Get DB credentials from .env file
		host := os.Getenv("HOST")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")
		dbport := os.Getenv("DB_PORT")
		// Define connection string
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, dbport)
		// Open connection with postgreSQL
		fmt.Println("dsn:", dsn)
		fmt.Println("Opening connection with PSQL...")
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println("Failed connecting DB:", err)
			// log.Fatalln("Failed connecting DB:", err)
		}
		return db
	}
	return db
}
