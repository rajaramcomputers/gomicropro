package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Customer Service")
	})

	log.Fatal(http.ListenAndServeTLS(":3000", "./cert.pem", "./key.pem", nil))
}
