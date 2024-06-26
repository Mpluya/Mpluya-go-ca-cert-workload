package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("helloworld: received a request")
	target := os.Getenv("TARGET")
	if target == "" {
		target = "Mae"
	}
	fmt.Fprintf(w, "Hello %s!\n", target)
}

func main() {
	log.Print("helloworld: starting server...")
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("USAGE: ca-certificates <url>")
	_, err := http.Head("https://support.broadcom.com")
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(2)
	}

	log.Printf("helloworld: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
