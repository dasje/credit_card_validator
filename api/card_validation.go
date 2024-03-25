package server

import (
	cctypes "credit_card_validation/resources"
	"credit_card_validation/validation_algorithms"
	"encoding/json"
	"net/http"
)

// Validate card number using luhn algorithm, and cvv length.
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

		var returnData cctypes.OutgoingCardNumber
	
		returnData.CardNumber = &cardNumValid
		returnData.CVVNumber = &cvvValid

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(returnData)
	}
}

// Identifiers whether card is accepted based on whether regex pattern can be identified in credit_card_regex resource.
func CardAccepted(w http.ResponseWriter, r *http.Request, regexResource []interface{}) {
	cardNumberString, _, _ := ParseBodyAsString(&w, r, "/card_accepted", "GET")
	if len(*cardNumberString) == 0 {
		http.Error(w, "Card number not found.", http.StatusBadRequest)
	} else {
		cardRegexSlice := regexResource[0]
		retVal := validation_algorithms.CheckCardAccepted(*cardNumberString, cardRegexSlice)
	
		returnData := map[string]interface{} {
			"cardAccepted": retVal,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(returnData)
	}
}
