package server

import (
	"credit_card_validation/validation_algorithms"
	"encoding/json"
	"net/http"
)

type OutgoingCardNumber struct {
	CardNumber *bool `json:"cardNumber"`
	CVVNumber *bool `json:"cvvNumber"`
}

func CardValidation(w http.ResponseWriter, r *http.Request) {
	cardNumInts, cardCvvInts := ParseBodyAsSlice(&w, r, "/validate_card", "GET")
	if len(*cardNumInts) == 0 {
		http.Error(w, "Card number not readable.", http.StatusNoContent)
		return
	} else if len(*cardCvvInts) == 0 {
		http.Error(w, "CVV not readable.", http.StatusNoContent)
		return
	} else {
		cardNumValid := validation_algorithms.LuhnIsValid(cardNumInts)
		cvvValid := validation_algorithms.CVVIsValid(*cardCvvInts)

		var returnData OutgoingCardNumber
	
		returnData.CardNumber = &cardNumValid
		returnData.CVVNumber = &cvvValid

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(returnData)
	}
}

func CardAccepted(w http.ResponseWriter, r *http.Request, cardRegexSlice interface{}) {
	cardNumberString, _ := ParseBodyAsString(&w, r, "/card_accepted", "GET")
	if len(*cardNumberString) == 0 {
		return
	} else {
		retVal := validation_algorithms.CheckCardAccepted(*cardNumberString, cardRegexSlice)
	
		returnData := map[string]interface{} {
			"cardAccepted": retVal,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(returnData)
	}
}
