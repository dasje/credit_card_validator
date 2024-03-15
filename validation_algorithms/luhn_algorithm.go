package validation_algorithms

// Takes an array of integers in the same order as the credit
// card number they come from. Include all digits.
// Returns a bool confirming if the Luhn Algorithm validated
// the digits.
func LuhnIsValid(cardNumber *[]int) bool {
	sum := 0
	parity := len(*cardNumber) % 2
	for i := 0; i < (len(*cardNumber) - 1); i++ {
		if (i%2) != parity {
			sum += (*cardNumber)[i]
		} else if ((*cardNumber))[i] > 4 {
			sum += (2*(*cardNumber)[i] - 9)
		} else {
			sum += (2 * (*cardNumber)[i])
		}
	}
	return (*cardNumber)[len(*cardNumber)-1] == (10 - (sum % 10))
}

