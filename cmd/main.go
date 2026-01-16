package main

import (
	"log"
	"net/http"
)

func main() {
	dir := "/Users/ruslanmuradov/github.com:gasuhwbab/static_file_server/templates/"
	handler := http.FileServer(http.Dir(dir))
	http.Handle("/", handler)
	log.Println("Listening on localhost:8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
