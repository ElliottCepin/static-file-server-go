package main

import (
	"fmt"
	"net/http"
	"log"
)
func main() {
	http.Handle("/public/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You've laneded on the home page - there is nothing to see here")
	})
	

	log.Fatal(http.ListenAndServe(":8080", nil))
}	
