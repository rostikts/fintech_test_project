// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Rostik",
            "email": "rostiktsyapura@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/transactions": {
            "get": {
                "description": "returns array of filtered transactions",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "returns array of transactions",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "filter by terminal_id",
                        "name": "terminal_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "filter by transaction_id",
                        "name": "transaction_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "filter by status",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "filter by payment_type",
                        "name": "payment_type",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "date",
                        "description": "filter from start date",
                        "name": "from",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "date",
                        "description": "filter to ending date",
                        "name": "to",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "partial match by narrative",
                        "name": "payment_narrative",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Transaction"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/transactions/csv": {
            "get": {
                "description": "returns csv file with filtered transactions",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "returns cvs with transactions",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "filter by terminal_id",
                        "name": "terminal_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "filter by transaction_id",
                        "name": "transaction_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "filter by status",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "filter by payment_type",
                        "name": "payment_type",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "date",
                        "description": "filter from start date",
                        "name": "from",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "date",
                        "description": "filter to ending date",
                        "name": "to",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "partial match by narrative",
                        "name": "payment_narrative",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Transaction"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/transactions/parse": {
            "post": {
                "description": "parse document with transactions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "parse document",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/handler.parseDocumentsRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/handler.parseDocumentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "echo.HTTPError": {
            "type": "object",
            "properties": {
                "message": {}
            }
        },
        "handler.parseDocumentResponse": {
            "type": "object",
            "properties": {
                "failed": {
                    "type": "integer"
                },
                "success": {
                    "type": "integer"
                }
            }
        },
        "handler.parseDocumentsRequest": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        },
        "models.Payee": {
            "type": "object",
            "properties": {
                "bankAccount": {
                    "type": "string"
                },
                "bankMfo": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Payment": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "narrative": {
                    "type": "string"
                },
                "number": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "models.Service": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Transaction": {
            "type": "object",
            "properties": {
                "amountOriginal": {
                    "type": "integer"
                },
                "amountTotal": {
                    "type": "integer"
                },
                "commissionClient": {
                    "type": "integer"
                },
                "commissionPS": {
                    "type": "number"
                },
                "commissionProvider": {
                    "type": "number"
                },
                "dateInput": {
                    "type": "string"
                },
                "datePost": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "partnerObjectID": {
                    "type": "integer"
                },
                "payee": {
                    "$ref": "#/definitions/models.Payee"
                },
                "payment": {
                    "$ref": "#/definitions/models.Payment"
                },
                "requestID": {
                    "type": "integer"
                },
                "service": {
                    "$ref": "#/definitions/models.Service"
                },
                "status": {
                    "type": "string"
                },
                "terminalID": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger API",
	Description:      "Documentation for test project for EVO trainee program",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
