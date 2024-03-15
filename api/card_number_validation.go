package server

import (
	"credit_card_validation/validation_algorithms"
	"encoding/json"
	"fmt"
	"net/http"
)

func LuhnValidation(w http.ResponseWriter, r *http.Request) {
	cardInts := ParseIncomingCardNumber(&w, r, "/luhn_validation", "GET")
	fmt.Println(cardInts)
	if len(*cardInts) == 0 {
		return
	} else {
		retVal := validation_algorithms.LuhnIsValid(cardInts)
		fmt.Println(retVal)
	
		returnData := map[string]interface{} {
			"cardValid": retVal,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(returnData)
	}
}
