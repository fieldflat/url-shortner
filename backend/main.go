package main

import (
	"log"
	"net/http"

	"github.com/fieldflat/url-shortner/backend/controller/handlers"
)

func main() {
	http.HandleFunc("/short-url/", handlers.GenerateShortURL)
	http.HandleFunc("/re/", handlers.RedirectOriginURL)
	http.HandleFunc("/", handlers.ReturnTopPage)

	log.Fatal(http.ListenAndServe(":80", nil))
	log.Println("Server start🎉")
}
