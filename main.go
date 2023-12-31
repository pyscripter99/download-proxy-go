package main

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed web/*
var web embed.FS

func main() {
	mux := http.NewServeMux()

	webStripped, err := fs.Sub(web, "web")
	if err != nil {
		log.Panic(err)
	}

	mux.Handle("/", loggerMiddleware(http.FileServer(http.FS(webStripped))))
	mux.HandleFunc("/download", downloadHandler)

	fmt.Println("Starting server on address: ':8080'")
	err = http.ListenAndServe(":8080", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(0)
	}
}
