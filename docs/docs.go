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
        "/opening": {
            "post": {
                "description": "Create a new job opening",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Openings"
                ],
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.CreateOpeningRequest"
                        }
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorReponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorReponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a new job opening",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Openings"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Opening indentification",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorReponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorReponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.CreateOpeningRequest": {
            "type": "object",
            "properties": {
                "company": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "remote": {
                    "type": "boolean"
                },
                "role": {
                    "type": "string"
                },
                "salary": {
                    "type": "integer"
                }
            }
        },
        "handler.ErrorReponse": {
            "type": "object",
            "properties": {
                "errorCode": {
                    "type": "string"
                },
                "message": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
