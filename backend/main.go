package main

import (
	"log"
	"net/http"

	"github.com/fieldflat/url-shortner/backend/controller/handlers"
)

func main() {
	http.HandleFunc("/short-url/", handlers.GenerateShortURL)
	http.HandleFunc("/", handlers.RedirectOriginURL)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
