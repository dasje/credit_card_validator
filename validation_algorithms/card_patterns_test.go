package validation_algorithms

import (
	"testing"
)
func TestAcceptance(t *testing.T) {
	var input = "1,7,8,9,3,7,2,9,9,7,4"
	var regexes = map[string]string{"Amex Card": "^3[47][0-9]{13}$"}
	output := CheckCardAccepted(input, regexes)
	expectedOutput := true
	
	if output != expectedOutput {
		t.Error(output, expectedOutput)
	}
}