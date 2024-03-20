package main

import (
	server "credit_card_validation/api"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func loadResources() (map[string]string, map[string]map[string]interface{}) {
	industryContent, industryErr := os.ReadFile("./resources/industry.json")
	var industryPayload map[string]string
	issuerContent, issuerErr := os.ReadFile("./resources/issuers.json")
	var issuerPayload map[string]map[string]interface{}
	if industryErr != nil {
		fmt.Println(industryErr)
	} else {
		jsonError := json.Unmarshal(industryContent, &industryPayload)
		if jsonError != nil {
			fmt.Println(jsonError)
		}
	}
	if issuerErr != nil {
		fmt.Print(issuerErr)
	} else {
		jsonError := json.Unmarshal(issuerContent, &issuerPayload)
		if jsonError != nil {
			fmt.Print(jsonError)
		}
	}
	return industryPayload, issuerPayload
}

type serverFunc func(http.ResponseWriter, *http.Request, interface{})

func handleWithResources(resource interface{}, serverFunc serverFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		serverFunc(w, r, resource)
	}
}

func main() {
	industryResources, issuerResources := loadResources()
	fmt.Print("Resources loaded from file.")
	
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/luhn_validation", handleWithResources(issuerResources, server.LuhnValidation))
	http.HandleFunc("/industry_validation", handleWithResources(industryResources, server.MajorIndustryValidation))

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}