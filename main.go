package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s", r.Method, r.URL, r.RemoteAddr)
		w.Write([]byte("Hello World!"))
	})

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(0)
	}
}
