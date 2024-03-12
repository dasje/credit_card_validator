package luhnalgorithm

func isValid(cardNumber []int) bool {
	// cardNumber is array of digits from the card number, minus the last digit
	sum := 0
	parity := len(cardNumber) % 2
	for i := 0; i < len(cardNumber); i++ {
		if i%2 != parity {
			sum += cardNumber[i]
		} else if cardNumber[i] > 4 {
			sum += (2*cardNumber[i] - 9)
		} else {
			sum += (2 * cardNumber[i])
		}
	}
	return cardNumber[len(cardNumber)-1] == (10 - (sum % 10))
}