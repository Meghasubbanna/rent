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
	// r.HandleFunc("/alluser", ArticlesHandler)

	log.Fatal(http.ListenAndServe(":3000", r))
	fmt.Println("My server is running on port 60000")
}
