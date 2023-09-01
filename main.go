package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/omprakas123/controller"
)

func main() {
	r := mux.NewRouter()

	// handling with routers
	r.HandleFunc("/user/signup", controller.Signup).Methods("POST")
	r.HandleFunc("/user/login", controller.Login).Methods("POST")
	r.HandleFunc("/create/book", controller.BookCreation).Methods("POST")
	r.HandleFunc("/book/availablebooks", controller.AvailableBooks).Methods("GET")
	r.HandleFunc("/user/buybook", controller.BookPurchase).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", r))
	fmt.Println("My server is running on port 5000")
}
