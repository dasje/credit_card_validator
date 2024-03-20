package server

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type IncomingCardNumber struct {
	CardNumber *string `json:"cardNumber"`
}

func ParseIncomingCardNumberReturnSlice(
	w *http.ResponseWriter,
	r *http.Request,
	url string,
	method string,
) *[]int {
	var cardInts []int

	if r.URL.Path != url {
		http.Error(*w, "404 Not found.", http.StatusNotFound)
		return &cardInts
	}

	if r.Method != method {
		http.Error(*w, "Method is not supported.", http.StatusNotFound)
		return &cardInts
	}

	body, bodyReadError := io.ReadAll(r.Body)
	if bodyReadError != nil {
		http.Error(*w, "Body not parsable.", http.StatusBadRequest)
		return &cardInts
	}

	// Extract card number from incoming JSON
	var newBody IncomingCardNumber
	var cardNumber string
	jsonReadError := json.Unmarshal(body, &newBody)
	if jsonReadError != nil {
		http.Error(*w, "Body is not valid JSON.", http.StatusBadRequest)
		return &cardInts
		} else if newBody.CardNumber == nil {
		// If card number not present, exit function.
		http.Error(*w, "JSON not parsable.", http.StatusBadRequest)
		return &cardInts
	}
	cardNumber = *newBody.CardNumber
	
	// Convert card number string to slice.
	cardStringInts := strings.SplitAfter(cardNumber, "")
	for _, v := range cardStringInts {
		newV, conversionError := strconv.Atoi(v)
		if conversionError != nil {
			http.Error(*w, "Card number not readable.", http.StatusBadRequest)
			return &cardInts
			} else {
			cardInts = append(cardInts, newV)
		}
	}
	return &cardInts
}

func ParseIncomingCardNumberReturnString(
	w *http.ResponseWriter,
	r *http.Request,
	url string,
	method string,
) string {
	var cardString string
	
	if r.URL.Path != url {
		http.Error(*w, "404 Not found.", http.StatusNotFound)
		return cardString
	}

	if r.Method != method {
		http.Error(*w, "Method is not supported.", http.StatusNotFound)
		return cardString
	}

	body, bodyReadError := io.ReadAll(r.Body)
	if bodyReadError != nil {
		http.Error(*w, "Body not parsable.", http.StatusBadRequest)
		return cardString
	}

	// Extract card number from incoming JSON
	var newBody IncomingCardNumber
	jsonReadError := json.Unmarshal(body, &newBody)
	if jsonReadError != nil {
		http.Error(*w, "Body is not valid JSON.", http.StatusBadRequest)
		return cardString
		} else if newBody.CardNumber == nil {
		// If card number not present, exit function.
		http.Error(*w, "JSON not parsable.", http.StatusBadRequest)
		return cardString
	}
	cardString = *newBody.CardNumber
	return cardString
}