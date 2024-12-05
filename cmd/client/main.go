package main

import (
	"flag"
	"fmt"
	"os"
)

var baseURL string // Global variable to store the base URL

func main() {
	// Define the `-baseURL` flag
	// Example client -baseURL=http://go-example-server
	flag.StringVar(&baseURL, "baseURL", "http://localhost", "Base URL of the server")
	flag.Parse()

	err := ping()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
