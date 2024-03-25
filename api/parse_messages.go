package server

import (
	cctypes "credit_card_validation/resources"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// Parse the received body.
// Unmarshall the JSON received and return the bytes.
func ParseBody(
	w *http.ResponseWriter,
	r *http.Request,
	url string,
	method string,
) cctypes.IncomingCardNumber {
	var newBody cctypes.IncomingCardNumber

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
		http.Error(*w, "Card number not present. JSON not parsable.", http.StatusBadRequest)
		return newBody
	} else if newBody.CVVNumber == nil {
		// If cvv not present, exit function.
		http.Error(*w, "CVV not present. JSON not parsable.", http.StatusBadRequest)
		return newBody
	}
	return newBody
}

// Convert card number and cvv bytes to int slices.
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

// Convert card number and cvv bytes to strings.
func ParseBodyAsString(
	w *http.ResponseWriter,
	r *http.Request,
	url string,
	method string,
) (*string, *string, error) {
	body := ParseBody(w, r, url, method)

	cvvNumber := body.CVVNumber
	cardNumber := body.CardNumber

	if len(*cvvNumber) == 0 {
		http.Error(*w, "CVV number not readable.", http.StatusNoContent)
		return nil, nil, errors.New("no cvv number")
	}
	if len(*cardNumber) == 0 {
		http.Error(*w, "Card number not readable.", http.StatusNoContent)
		return nil, nil, errors.New("no card number")
	}

	return cardNumber, cvvNumber, nil
}