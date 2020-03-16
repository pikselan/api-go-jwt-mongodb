package main

import (
	"github.com/gorilla/mux"
	"github.com/pikselan/api-go-jwt-mongodb/controller"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", controller.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", controller.LoginHandler).Methods("POST")
	r.HandleFunc("/profile", controller.ProfileHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", r))
}