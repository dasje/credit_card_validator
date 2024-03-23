package validation_algorithms

import (
	"testing"
)
func TestAlgorithm(t *testing.T) {
	var input = []int{1,7,8,9,3,7,2,9,9,7,4}
	output := LuhnIsValid(&input)
	expectedOutput := true
	
	if output != expectedOutput {
		t.Error(output, expectedOutput)
	}
}