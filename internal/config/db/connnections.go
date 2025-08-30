package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/KalpeshJangir23/todolist_go/internal/models"

)

// Global DB instance
var DB *gorm.DB

// Connect opens a connection to PostgreSQL
func Connect() *gorm.DB {
	// You can load these from env later
	host := "localhost"
	user := "todouser"
	password := "todopass"
	dbname := "todolist_go"
	port := "5432"

	// DSN = connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	// Open DB
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database: ", err)
	}

	fmt.Println("✅ Connected to PostgreSQL successfully")

	err = DB.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatal("❌ Failed to migrate database: ", err)
	}
	fmt.Println("✅ Database migrated successfully")

	return DB;
}
