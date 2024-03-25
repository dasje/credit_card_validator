package main

// @title Credit Card Validator API documentation
// @version 1.0.0// @host localhost:8000
// @BasePath /

import (
	server "credit_card_validation/api"
	docs "credit_card_validation/docs"
	cctypes "credit_card_validation/resources"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	httpSwagger "github.com/swaggo/http-swagger"
)

// Load static JSON resources into memory.
// Returns:
// Map of major industries, where the key is the industry number.
// Map of issuer data, where keys are numbers (card number identifiers), validation method, and max card number length.
// Map of regex patterns for accepted card number. Keys are issuer names.
func loadResources() (map[string]string, map[string]cctypes.ResourceTypeIssuer, map[string]string) {
	industryContent, industryErr := os.ReadFile("./resources/industry.json")
	var industryPayload map[string]string
	if industryErr != nil {
		fmt.Println(industryErr)
	} else {
		jsonError := json.Unmarshal(industryContent, &industryPayload)
		if jsonError != nil {
			fmt.Println(jsonError)
		}
	}

	issuerContent, issuerErr := os.ReadFile("./resources/issuers.json")
	var issuerPayload map[string]cctypes.ResourceTypeIssuer
	if issuerErr != nil {
		fmt.Print(issuerErr)
	} else {
		jsonError := json.Unmarshal(issuerContent, &issuerPayload)
		if jsonError != nil {
			fmt.Print(jsonError)
		}
	}

	cardRegexContent, cardRegexErr := os.ReadFile("./resources/credit_card_regex.json")
	var cardRegexPayload map[string]string
	if cardRegexErr != nil {
		fmt.Print(cardRegexErr)
	} else {
		jsonError := json.Unmarshal(cardRegexContent, &cardRegexPayload)
		if jsonError != nil {
			fmt.Print(jsonError)
		}
	}

	return industryPayload, issuerPayload, cardRegexPayload
}

type serverFunc func(http.ResponseWriter, *http.Request, []interface{})

// Wrapper function for http handler func. Permits passing resources to handler functions.
func handleWithResources(serverFunc serverFunc, resource ...interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		serverFunc(w, r, resource)
	}
}

func main() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Credit Card Validator API"
	docs.SwaggerInfo.Description = "Server for validation credit card numbers."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	industryResources, issuerResources, cardRegex := loadResources()
	
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/docs/", httpSwagger.WrapHandler.ServeHTTP)
	http.HandleFunc("/card_accepted", handleWithResources(server.CardAccepted, cardRegex))
	http.HandleFunc("/validate_card", server.CardValidation)
	http.HandleFunc("/card_info", handleWithResources(server.DeconstructCardInfo, issuerResources, industryResources))

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}