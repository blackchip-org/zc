package main

import (
	"fmt"
	"os"

	"net/http"
)

var addr = ":9090"

func main() {
	fmt.Printf("listening on %v\n", addr)
	err := http.ListenAndServe(addr, http.FileServer(http.Dir("web")))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: unable to start server: %v\n", err)
		os.Exit(1)
	}
}
