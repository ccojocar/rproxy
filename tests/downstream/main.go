package main

import (
	"fmt"
	"net/http"
	"os"
)

var name string

func main() {
	if len(os.Args) != 3 {
		fmt.Println("command arguments missing: 'downstream <name> <address>'")
		os.Exit(1)
	}
	name = os.Args[1]
	address := os.Args[2]

	http.HandleFunc("/", handleTestDownstreamService)
	if err := http.ListenAndServe(address, nil); err != nil {
		fmt.Printf("failed to listen on: %s", address)
	}
}

func handleTestDownstreamService(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "%s\n", name)
}
