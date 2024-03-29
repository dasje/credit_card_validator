// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/card_accepted": {
            "post": {
                "description": "Identifiers whether card is accepted based on whether regex pattern can be identified in credit_card_regex resource.",
                "consumes": [
                    "application/json"
                ],
                "summary": "Checks if card accepted",
                "parameters": [
                    {
                        "description": "Card numbers",
                        "name": "cardNumbers",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/cctypes.IncomingCardNumber"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/cctypes.OutgoingCardAccepted"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/card_info": {
            "post": {
                "description": "Identify the major industry and card issuer from the provided card number.",
                "consumes": [
                    "application/json"
                ],
                "summary": "Identify major industry",
                "parameters": [
                    {
                        "description": "Card numbers",
                        "name": "cardNumbers",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/cctypes.IncomingCardNumber"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/cctypes.OutgoingIssuerData"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/validate_card": {
            "post": {
                "description": "Validate card number using luhn algorithm, and cvv length.",
                "consumes": [
                    "application/json"
                ],
                "summary": "Validate card number",
                "parameters": [
                    {
                        "description": "Card numbers",
                        "name": "cardNumbers",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/cctypes.IncomingCardNumber"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/cctypes.OutgoingCardNumber"
                        }
                    },
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "cctypes.IncomingCardNumber": {
            "type": "object",
            "properties": {
                "cardNumber": {
                    "type": "string"
                },
                "cvvNumber": {
                    "type": "string"
                }
            }
        },
        "cctypes.OutgoingCardAccepted": {
            "type": "object",
            "properties": {
                "cardAccepted": {
                    "type": "boolean"
                }
            }
        },
        "cctypes.OutgoingCardNumber": {
            "type": "object",
            "properties": {
                "cardNumber": {
                    "type": "boolean"
                },
                "cvvNumber": {
                    "type": "boolean"
                }
            }
        },
        "cctypes.OutgoingIssuerData": {
            "type": "object",
            "properties": {
                "cardIssuer": {
                    "type": "string"
                },
                "majorIndustry": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0// @host localhost:8000",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Credit Card Validator API documentation",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
