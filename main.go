package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	version := "1.0.11"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello risk-tech\n")
	})

	fmt.Printf("RSMN demo version %v service running\n", version)

	log.Fatal(http.ListenAndServe(":5555", nil))

}
