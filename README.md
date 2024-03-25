# credit_card_validator
## Functionality
### Validation algorithms
- #### Luhn
Most credit card numbers can be confirmed using the Luhn algorithm.
-#### CVV validation
Confirms the card CVV has correct number of digits.
- #### Number patterns
Matches number patterns to known card issuer patters using regex.
### Parsing functions
- #### cardNumberToSlice
Returns the card number reduced to digits, with each digit as an element in the slice.
- #### cardNumberToString
Returns the card number reduced to digits, as a single string.
- #### identifyMajorIndustry
Returns the identified major industry based on the first number in the card number.
- #### identifyCardIssuer
Returns the identified card issuer.
- #### identifyAccountIdentifier
Returns the account identifier.
### API
- #### /validate_card
This endpoint accepts a GET request with JSON body. Expected keys are "cardNumber" and "csvNumber".
This returns a JSON with the same keys and boolean values.
- #### /card_info
This endpoint accepts a GET request with JSON body. Expected keys are "cardNumber" and "csvNumber".
This returns a JSON with the following keys: "majorIndustry", "cardIssuer", "accountIdentifier".
### 
## Tests
### HTTP tests
These tests are defined in tests.http. I use the REST Client plugin for VSCode.