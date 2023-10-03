package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	fmt.Println("Starting http server...")

	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}
