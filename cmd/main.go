package main

import (
	"HTTP-API/internal"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/rs/cors"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/{id}", internal.GetUserHandler).Methods("GET")
	router.HandleFunc("/api/users", internal.GetAllUsersHandler).Methods("GET")
	router.HandleFunc("/api/createUser", internal.CreateUserHandler).Methods("POST")

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
