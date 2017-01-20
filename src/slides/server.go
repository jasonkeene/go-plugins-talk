package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	wd, _ := os.Getwd()
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir(wd))))
}
