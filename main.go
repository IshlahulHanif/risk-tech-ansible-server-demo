package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello risk-tech")
	})

	log.Fatal(http.ListenAndServe(":5555", nil))

}
