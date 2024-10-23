package handler

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/pajri/url-shortener/storage"
)

type URLRequest struct {
	URL string `json:"url"`
}

type URLResponse struct {
	ShortenedURL string `json:"shortened_url"`
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// generateShortURL generates a random string to be used as a short URL
func generateShortURL() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// HomeHandler is the root handler that returns a welcome message
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Go URL Shortener!"))
}

// ShortenHandler handles the shortening of the URL
func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	var urlReq URLRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &urlReq)
	if err != nil || urlReq.URL == "" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	shortURL := generateShortURL()
	storage.SaveURL(shortURL, urlReq.URL)

	response := URLResponse{
		ShortenedURL: "http://localhost:8080/r/" + shortURL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RedirectHandler handles the redirection of short URLs to the original URL
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[len("/r/"):]

	originalURL, found := storage.GetURL(shortURL)
	if !found {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}
