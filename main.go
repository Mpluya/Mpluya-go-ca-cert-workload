package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("USAGE: ca-certificates <url>")
	_, err := http.Head("https://google.com")
	if err != nil {
		fmt.Println("ERROR:", err)
		//os.Exit(2)
	}
}
