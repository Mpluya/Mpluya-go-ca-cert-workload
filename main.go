package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("USAGE: ca-certificates <url>")
	_, err := http.Head("https://harbor-repo-ui.vmware.com:443")
	if err != nil {
		fmt.Println("ERROR:", err)
		//os.Exit(2)
	}
}
