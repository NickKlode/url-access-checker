package main

import (
	"inhouseadtest/internal/api"
	"inhouseadtest/internal/service"
	"log"
	"net/http"
)

func main() {
	go service.RequestLoop()
	api := api.New()

	log.Fatal(http.ListenAndServe(":8080", api.Router()))

}
