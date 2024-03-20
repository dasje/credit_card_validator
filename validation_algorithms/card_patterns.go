package validation_algorithms

import (
	"regexp"
)

// Identify card issuers based on cardRegexes resource.
func CheckCardIssuer(cardNumberString string, cardRegexes interface{}) []string {
	// Force type of incoming interface
	regexSlice := cardRegexes.(map[string]string)

	var issuers []string

	// Remove non-digits from card number
	var digitsRe = regexp.MustCompile(`\D`)
	onlyDigitsCardNumber := digitsRe.ReplaceAllString(cardNumberString, "")

	for k, v := range regexSlice {
		regex := regexp.MustCompile(v)
		if regex.Match([]byte(onlyDigitsCardNumber)) {
			issuers = append(issuers, k)
		}
	}
	return issuers
}

// Identify whether card number matches any regex definitions of issuer card number patterns.
func CheckCardAccepted(cardNumberString string, cardRegexes interface{}) bool {
	// Force type of incoming interface
	regexSlice := cardRegexes.(map[string]string)

	// Remove non-digits from card number
	var digitsRe = regexp.MustCompile(`\D`)
	onlyDigitsCardNumber := digitsRe.ReplaceAllString(cardNumberString, "")

	for _, v := range regexSlice {
		regex := regexp.MustCompile(v)
		if regex.Match([]byte(onlyDigitsCardNumber)) {
			return true
		}
	}
	return false
}