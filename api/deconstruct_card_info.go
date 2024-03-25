package server

import (
	cctypes "credit_card_validation/resources"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

// Identify the major industry and card issuer from the provided card number.
func DeconstructCardInfo(w http.ResponseWriter, r *http.Request, industryResources []interface{}) {
	cardNumber, _, parseErr := ParseBodyAsString(&w, r, "/card_info", "GET")
	if parseErr != nil {
		fmt.Print("Parsing error.")
		return
	} else {
		majorIndustries := industryResources[1].(map[string]string)
		majorIndustry, miErr := IdentifyMajorIndustry(*cardNumber, majorIndustries)
		if miErr != nil {
			http.Error(w, "Industry cannot be identified", http.StatusUnprocessableEntity)
		}

		issuers := industryResources[0].(map[string]cctypes.ResourceTypeIssuer)
		issuer, issuerErr := IdentifyCardIssuer(*cardNumber, issuers)
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

// Identify the card issuer by comparing initial values in card number with issuer card number patterns.
func IdentifyCardIssuer(cardNumber string, issuers map[string]cctypes.ResourceTypeIssuer) (string, error) {
	for issuerName, issuerData := range issuers {
		for _, v := range issuerData.Numbers {
			strValue := strconv.Itoa(v)
			strValue = "^" + strValue
			var startRegex = regexp.MustCompile(strValue)
			if startRegex.Match([]byte(cardNumber)) {
				return issuerName, nil
			}
		}
	}
	return "", errors.New("no issuer identified")
}

// Identify the major industry associated with the card issuer based on index[0] of card number.
func IdentifyMajorIndustry(cardNumber string, industries map[string]string) (string, error) {
	firstIndex := strings.Split(cardNumber, "")[0]
	mi := industries[firstIndex]
	if len(mi) == 0 {
		return "", errors.New("no identified industry")
	}
	return mi, nil
}

