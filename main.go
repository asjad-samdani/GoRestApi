package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/asjad-samdani/restApiInGo/database"
	"github.com/asjad-samdani/restApiInGo/handlers"
	"github.com/gorilla/mux"
)

// Routing

func main() {
	database.ConnectDatabase()
	r := mux.NewRouter()
	r.HandleFunc("/user", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/user", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", handlers.DeleteUser).Methods("DELETE")

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
	fmt.Println("Server is running on port 8000...")

}
