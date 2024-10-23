package storage

import "sync"

// URLStore stores the mapping of shortened URLs to original URLs
var URLStore = make(map[string]string)
var mu sync.Mutex

// SaveURL saves the short URL and its corresponding original URL
func SaveURL(shortURL, originalURL string) {
	mu.Lock()
	defer mu.Unlock()
	URLStore[shortURL] = originalURL
}

// GetURL retrieves the original URL for a given short URL
func GetURL(shortURL string) (string, bool) {
	mu.Lock()
	defer mu.Unlock()
	originalURL, found := URLStore[shortURL]
	return originalURL, found
}
