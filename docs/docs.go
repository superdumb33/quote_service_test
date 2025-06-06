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
        "/quotes": {
            "get": {
                "description": "Returns all of the quotes or filters by author if query is provided",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "quotes"
                ],
                "summary": "Get quotes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Quote author",
                        "name": "author",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.GetAllQuotesResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new record with provided quote",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "quotes"
                ],
                "summary": "Create a new quote",
                "parameters": [
                    {
                        "description": "Quote payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateQuoteRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateQuoteResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/quotes/random": {
            "get": {
                "description": "Returs random quote",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "quotes"
                ],
                "summary": "Get a random quote",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetRandomQuoteResponse"
                        }
                    },
                    "404": {
                        "description": "if there's no stored qoutes",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/quotes/{id}": {
            "delete": {
                "description": "Deletes a quote by it's UUID",
                "tags": [
                    "quotes"
                ],
                "summary": "Delete quote by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Quote ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "invalid ID",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateQuoteRequest": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "quote": {
                    "type": "string"
                }
            }
        },
        "dto.CreateQuoteResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "dto.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "dto.GetAllQuotesResponse": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "quote": {
                    "type": "string"
                }
            }
        },
        "dto.GetRandomQuoteResponse": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "quote": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
