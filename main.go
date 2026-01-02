package main

import (
	"fmt"
	"net/http"
	"log"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You've landed on the home page - there is nothing to see here")
}

func main() {
	http.Handle("/public/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/", home)
	

	log.Fatal(http.ListenAndServe(":8080", nil))
}	
