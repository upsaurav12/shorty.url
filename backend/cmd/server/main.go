package main

import (
	"log"
	"net/http"

	"github.com/upsaurav12/url_shorty/internal/shortner"
)

func main() {
	urlShortener := shortner.NewURLShortner()

	shortner.URLShortnerHandler(urlShortener)

	log.Printf("server has been started at port :8080")
	http.ListenAndServe(":8080", nil)
}
