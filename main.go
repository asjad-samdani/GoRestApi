package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Routing
func initializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/user", GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}", GetUser).Methods("GET")
	r.HandleFunc("/user", CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")
	
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
	fmt.Println("Server is running on port 8000...")

}

func main() {
	ConnectDatabase()
	initializeRouter()

}
