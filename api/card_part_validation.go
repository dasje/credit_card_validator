package server

import (
	"encoding/json"
	"net/http"
	"strconv"
)


func MajorIndustryValidation(w http.ResponseWriter, r *http.Request, industryResources interface{}) {
	cardInts := ParseIncomingCardNumberReturnSlice(&w, r, "/industry_validation", "GET")
	if len(*cardInts) == 0 {
		return
	} else {
		industryResources := industryResources.(map[string]string)
		ci := *cardInts
		firstNumber := strconv.Itoa(ci[0])
		retVal := industryResources[firstNumber]
		returnData := map[string]interface{} {
			"cardIndustry": retVal,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(returnData)
	}
}