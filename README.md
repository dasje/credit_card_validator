# credit_card_validator
## Overview
Deployable credit card validation server. Dockerfile and docker-compose.yaml included.

This sever was built using the following public resources:

- Credit card regex patterns: https://gist.github.com/michaelkeevildown/9096cd3aac9029c4e6e05588448a8841
- Major industry identifiers: https://en.wikipedia.org/wiki/ISO/IEC_7812#Major_industry_identifier

- Test card numbers: https://docs.adyen.com/development-resources/testing/test-card-numbers/
## API
- #### /card_accepted
This endpoint accepts a GET request with JSON body. Expected keys are "cardNumber" and "cvvNumber".
This returns a JSON with the key cardAccepted and boolean value.
- #### /validate_card
This endpoint accepts a GET request with JSON body. Expected keys are "cardNumber" and "cvvNumber".
This returns a JSON with the same keys and boolean values.
- #### /card_info
This endpoint accepts a GET request with JSON body. Expected keys are "cardNumber" and "cvvNumber".
This returns a JSON with the following keys: "majorIndustry", "cardIssuer", "accountIdentifier".
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
## Tests
### HTTP tests
These tests are defined in tests.http. I use the REST Client plugin for VSCode.