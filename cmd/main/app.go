package main

import (
	"log"
	"net/http"
	"user-rest-api/internal/user"

	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Println("router initializing")
	router := httprouter.New()

	log.Println("handler initializing")
	handler := user.NewHandler()
	handler.Register(router)

	log.Println("server running...")
	http.ListenAndServe(":8080", router)
}
