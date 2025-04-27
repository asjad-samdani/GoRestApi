// package database

// import (
// 	"log"

// 	"github.com/asjad-samdani/restApiInGo/model"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB
// var err error

// const DSN = "user:@tcp(127.0.0.1:3306)/rest_api?charset=utf8mb4&parseTime=True&loc=Local"

// // Database Connection
// func ConnectDatabase() {
// 	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Failed to connect to the database:", err)
// 	}
// 	err = DB.AutoMigrate(&model.User{})
// 	if err != nil {
// 		log.Fatalf("Error occurred in database migration: %v", err)
// 	}
// 	log.Println("Database initialized successfully")
// }

package database

import (
	"log"

	"github.com/asjad-samdani/restApiInGo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DSN = "user:mypassword@tcp(127.0.0.1:3306)/rest_api?charset=utf8mb4&parseTime=True&loc=Local"

// ConnectDatabase initializes the DB connection
func ConnectDatabase() {
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Check if the DB connection is properly initialized
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get the underlying SQL database connection: %v", err)
	}

	// Check the connection to the database (pinging)
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("Error occurred in database migration: %v", err)
	}

	log.Println("Database initialized successfully")
}
