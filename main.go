package main

import (
	"log"
	"net/http"

	"github.com/korero-chat/backend/routes"
)

func main() {
	router := routes.SetRoutes()
	log.Fatal(http.ListenAndServe(":8888", router))
}
