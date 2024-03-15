package main

import (
	server "credit_card_validation/api"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/luhn_validation", server.LuhnValidation)
	http.HandleFunc("/industry", server.MajorIndustryValidation)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}