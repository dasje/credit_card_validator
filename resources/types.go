package cctypes

// Card issuer metadata.
type ResourceTypeIssuer struct {
	Numbers    []int  `json:"numbers"`
	Length     int    `json:"length"`
	Validation string `json:"string"`
}

// Expected JSON in body of GET requests.
type IncomingCardNumber struct {
	CardNumber *string `json:"cardNumber"`
	CVVNumber  *string `json:"cvvNumber"`
}

// JSON structure of succesful responses to /card_info endpoint.
type OutgoingIssuerData struct {
	MajorIndustry string `json:"majorIndustry"`
	CardIssuer    string `json:"cardIssuer"`
}

// JSON structure of succesful responses to /validate_card endpoint.
type OutgoingCardNumber struct {
	CardNumber *bool `json:"cardNumber"`
	CVVNumber  *bool `json:"cvvNumber"`
}

// JSON structure of succesful responses to /card_accepted endpoint.
type OutgoingCardAccepted struct {
	CardAccepted bool `json:"cardAccepted"`
}