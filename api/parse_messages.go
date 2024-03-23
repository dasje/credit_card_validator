package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type IncomingCardNumber struct {
	CardNumber *string `json:"cardNumber"`
	CVVNumber *string `json:"cvvNumber"`
}

func ParseBody(
	w *http.ResponseWriter,
	r *http.Request,
	url string,
	method string,
) IncomingCardNumber {
	var newBody IncomingCardNumber

	if r.URL.Path != url {
		http.Error(*w, "404 Not found.", http.StatusNotFound)
		return newBody
	}

	if r.Method != method {
		http.Error(*w, "Method is not supported.", http.StatusNotFound)
		return newBody
	}

	body, bodyReadError := io.ReadAll(r.Body)
	if bodyReadError != nil {
		http.Error(*w, "Body not parsable.", http.StatusBadRequest)
		return newBody
	}

	jsonReadError := json.Unmarshal(body, &newBody)
	if jsonReadError != nil {
		http.Error(*w, "Body is not valid JSON.", http.StatusBadRequest)
		return newBody
	} else if newBody.CardNumber == nil {
		// If card number not present, exit function.
		http.Error(*w, "JSON not parsable.", http.StatusBadRequest)
		return newBody
	} else if newBody.CVVNumber == nil {
		// If cvv not present, exit function.
		http.Error(*w, "JSON not parsable.", http.StatusBadRequest)
		return newBody
	}
	fmt.Print(newBody)
	return newBody
}

func ParseBodyAsSlice(
	w *http.ResponseWriter,
	r *http.Request,
	url string,
	method string,
) (*[]int, *[]int) {
	body := ParseBody(w, r, url, method)
	
	var cardInts []int
	var cvvInts []int

	var cvvNumber string
	var cardNumber string
	cvvNumber = *body.CVVNumber
	cardNumber = *body.CardNumber
	
	// Convert card number string to slice.
	cardNumberStringInts := strings.SplitAfter(cardNumber, "")
	for _, v := range cardNumberStringInts {
		newV, conversionError := strconv.Atoi(v)
		if conversionError != nil {
			http.Error(*w, "Card number not readable.", http.StatusBadRequest)
			return &cardInts, &cvvInts
		} else {
			cardInts = append(cardInts, newV)
		}
	}
	// Convert cvv number string to slice.
	cvvNumberStringInts := strings.SplitAfter(cvvNumber, "")
	for _, v := range cvvNumberStringInts {
		newV, conversionError := strconv.Atoi(v)
		if conversionError != nil {
			http.Error(*w, "Card number not readable.", http.StatusBadRequest)
			return &cardInts, &cvvInts
		} else {
			cvvInts = append(cvvInts, newV)
		}
	}

	return &cardInts, &cvvInts
}

func ParseBodyAsString(
	w *http.ResponseWriter,
	r *http.Request,
	url string,
	method string,
) (*string, *string) {
	body := ParseBody(w, r, url, method)

	cvvNumber := body.CVVNumber
	cardNumber := body.CardNumber

	return cardNumber, cvvNumber
}