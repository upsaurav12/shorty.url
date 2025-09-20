package shortner

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type URLShortener struct {
	urls map[string]string
}

type shortURLStruct struct {
	UrlShortened string `json:"urlShortened"`
}

func (us *URLShortener) HandleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	originalURL := r.FormValue("url")
	if originalURL == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	// Generate a unique shortened key for the original URL
	shortKey := generateShortKey()
	us.urls[shortKey] = originalURL

	// Construct the full shortened URL
	shortenedURL := fmt.Sprintf("http://localhost:8080/short/%s", shortKey)

	shortURLStruct := shortURLStruct{
		UrlShortened: shortenedURL,
	}

	fmt.Println("struct: ", shortURLStruct.UrlShortened)

	// Render the HTML response with the shortened URL
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(&shortURLStruct)
}
func (us *URLShortener) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	shortKey := r.URL.Path[len("/short/"):]
	if shortKey == "" {
		http.Error(w, "Shortened key is missing", http.StatusBadRequest)
		return
	}

	// Retrieve the original URL from the `urls` map using the shortened key
	originalURL, found := us.urls[shortKey]
	if !found {
		http.Error(w, "Shortened key not found", http.StatusNotFound)
		return
	}

	// Add scheme if missing
	if !strings.HasPrefix(originalURL, "http://") && !strings.HasPrefix(originalURL, "https://") {
		originalURL = "http://" + originalURL
	}

	// Redirect the user to the original URL
	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

const keyLength = 6

func generateShortKey() string {
	rand.Seed(time.Now().UnixNano())
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}

	return string(shortKey)
}

func NewURLShortner() *URLShortener {
	return &URLShortener{
		urls: make(map[string]string),
	}
}

func URLShortnerHandler(urlShortener *URLShortener) {
	http.HandleFunc("/shorten", urlShortener.HandleShorten)
	http.HandleFunc("/short/", urlShortener.HandleRedirect)
}
