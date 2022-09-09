package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	version := "1.0.20"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is risk tech demo version %v service\n", version)
	})

	fmt.Printf("RSMN demo version %v service running\n", version)

	log.Fatal(http.ListenAndServe(":5555", nil))

}
