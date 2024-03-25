package cctypes

type ResourceTypeIssuer struct {
	Numbers    []int  `json:"numbers"`
	Length     int    `json:"length"`
	Validation string `json:"string"`
}

type IncomingCardNumber struct {
	CardNumber *string `json:"cardNumber"`
	CVVNumber  *string `json:"cvvNumber"`
}

type OutgoingIssuerData struct {
	MajorIndustry string `json:"majorIndustry"`
	CardIssuer    string `json:"cardIssuer"`
}

type OutgoingCardNumber struct {
	CardNumber *bool `json:"cardNumber"`
	CVVNumber  *bool `json:"cvvNumber"`
}

type OutgoingCardAccepted struct {
	CardAccepted bool `json:"cardAccepted"`
}