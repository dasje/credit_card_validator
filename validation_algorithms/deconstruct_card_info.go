package validation_algorithms

import (
	cctypes "credit_card_validation/resources"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

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

