package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
)

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s", r.Method, r.URL, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rawURL := r.URL.Query().Get("url")
	if rawURL == "" {
		http.Error(w, "Missing 'url' parameter", http.StatusBadRequest)
		return
	}

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid URL: %s", err), http.StatusBadRequest)
		return
	}

	// Open URL
	resp, err := http.Get(parsedURL.String())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching remote file: %s", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	filename := filepath.Base(parsedURL.Path)
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", fmt.Sprint(resp.ContentLength))

	buffer := make([]byte, 1024)

	for {
		n, err := resp.Body.Read(buffer)
		if err != nil && err != io.EOF {
			http.Error(w, fmt.Sprintf("Error reading remote file: %s", err), http.StatusInternalServerError)
			return
		}

		_, err = w.Write(buffer[:n])
		if err != nil {
			// Client may have disconnected
			return
		}

		if err == io.EOF || n == 0 {
			break
		}
	}
}
