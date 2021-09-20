package main

import (
	"HTTP-API-TestTask/internal"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	server := internal.NewUserServer()

	mux.HandleFunc("/api/user", server.UserHandler)
	mux.HandleFunc("api/user?id=", server.UserHandler)

	log.Fatal(http.ListenAndServe("localhost:"+os.Getenv("SERVERPORT"), mux))
}
