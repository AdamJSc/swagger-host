package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed static/*
var fs embed.FS

func main() {
	http.Handle("/", http.FileServer(http.FS(fs)))
	addr := ":8080"
	log.Printf("listening on %s...", addr)
	log.Fatal(http.ListenAndServe(addr, http.DefaultServeMux))
}
