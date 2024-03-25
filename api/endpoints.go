package server

import (
	cctypes "credit_card_validation/resources"
	"credit_card_validation/validation_algorithms"
	"encoding/json"
	"fmt"
	"net/http"
)

// CardAccepted ... Validate card number
// @Summary Validate card number
// @Description Validate card number using luhn algorithm, and cvv length.
// @Accept json
// @Param cardNumbers body cctypes.IncomingCardNumber true "Card numbers"
// @Success 200 {object} cctypes.OutgoingCardNumber
// @Failure 204 {object} object
// @Router /validate_card [post]
func CardValidation(w http.ResponseWriter, r *http.Request) {
	cardNumInts, cardCvvInts := ParseBodyAsSlice(&w, r, "/validate_card", "POST")
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

// CardAccepted ... Checks if card accepted
// @Summary Checks if card accepted
// @Description Identifiers whether card is accepted based on whether regex pattern can be identified in credit_card_regex resource.
// @Accept json
// @Param cardNumbers body cctypes.IncomingCardNumber true "Card numbers"
// @Success 200 {object} cctypes.OutgoingCardAccepted
// @Failure 400 {object} object
// @Router /card_accepted [post]
func CardAccepted(w http.ResponseWriter, r *http.Request, regexResource []interface{}) {
	cardNumberString, _, _ := ParseBodyAsString(&w, r, "/card_accepted", "POST")
	if len(*cardNumberString) == 0 {
		http.Error(w, "Card number not found.", http.StatusBadRequest)
	} else {
		cardRegexSlice := regexResource[0]
		retVal := validation_algorithms.CheckCardAccepted(*cardNumberString, cardRegexSlice)
	
		var returnData cctypes.OutgoingCardAccepted
		returnData.CardAccepted = retVal

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(returnData)
	}
}

// CardAccepted ... Identify major industry
// @Summary Identify major industry
// @Description Identify the major industry and card issuer from the provided card number.
// @Accept json
// @Param cardNumbers body cctypes.IncomingCardNumber true "Card numbers"
// @Success 200 {object} cctypes.OutgoingIssuerData
// @Failure 422 {object} object
// @Router /card_info [post]
func DeconstructCardInfo(w http.ResponseWriter, r *http.Request, industryResources []interface{}) {
	cardNumber, _, parseErr := ParseBodyAsString(&w, r, "/card_info", "POST")
	if parseErr != nil {
		fmt.Print("Parsing error.")
		return
	} else {
		majorIndustries := industryResources[1].(map[string]string)
		majorIndustry, miErr := validation_algorithms.IdentifyMajorIndustry(*cardNumber, majorIndustries)
		if miErr != nil {
			http.Error(w, "Industry cannot be identified", http.StatusUnprocessableEntity)
		}

		issuers := industryResources[0].(map[string]cctypes.ResourceTypeIssuer)
		issuer, issuerErr := validation_algorithms.IdentifyCardIssuer(*cardNumber, issuers)
		if issuerErr != nil {
			http.Error(w, "Issuer cannot be identified", http.StatusUnprocessableEntity)
		}

		var returnData cctypes.OutgoingIssuerData 
		returnData.MajorIndustry = majorIndustry
		returnData.CardIssuer = issuer
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(returnData)
	}
}