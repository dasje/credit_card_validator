package server

import (
	"credit_card_validation/validation_algorithms"
	"encoding/json"
	"fmt"
	"net/http"
)

func CardValidation(w http.ResponseWriter, r *http.Request, issuerMap interface{}) {
	cardInts := ParseIncomingCardNumberReturnSlice(&w, r, "/validate_card", "GET")
	if len(*cardInts) == 0 {
		return
	} else {
		retVal := validation_algorithms.LuhnIsValid(cardInts)
	
		returnData := map[string]interface{} {
			"cardValid": retVal,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(returnData)
	}
}

func CardAccepted(w http.ResponseWriter, r *http.Request, cardRegexSlice interface{}) {
	cardNumberString := ParseIncomingCardNumberReturnString(&w, r, "/check_card_accepted", "GET")
	fmt.Print(cardNumberString)
	if len(cardNumberString) == 0 {
		return
	} else {
		retVal := validation_algorithms.CheckCardAccepted(cardNumberString, cardRegexSlice)
	
		returnData := map[string]interface{} {
			"cardAccepted": retVal,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(returnData)
	}
}
