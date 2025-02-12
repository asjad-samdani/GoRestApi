package database

import (
	"log"

	"github.com/asjad-samdani/restApiInGo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DSN = "root:root@tcp(127.0.0.1:3306)/restapi?charset=utf8mb4&parseTime=True&loc=Local"

// Database Connection
func ConnectDatabase() {
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("Error occurred in database migration: %v", err)
	}
	log.Println("Database initialized successfully")
}
