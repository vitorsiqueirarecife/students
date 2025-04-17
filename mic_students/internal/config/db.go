package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/vitorsiqueirarecife/students/internal/modules/student/domain"
)

func InitDB() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env:", err)
	}
	// Load environment variables
	dbHost := os.Getenv("MYSQL_HOST")
	dbName := os.Getenv("MYSQL_DATABASE")
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&domain.Student{})
	return db, nil
}
