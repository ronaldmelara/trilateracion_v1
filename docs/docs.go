// Code generated by swaggo/swag. DO NOT EDIT
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
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/topsecret/": {
            "post": {
                "description": "Obtains the position of the enemy ship and deciphers the message received on the satellites",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TopSecret"
                ],
                "summary": "Get position and message",
                "parameters": [
                    {
                        "description": "Se debe enviar por el body la siguiente estructura JSON",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.TopSecret"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ResponseTopSecret"
                        }
                    }
                }
            }
        },
        "/topsecret_split/{satellite_name}": {
            "get": {
                "description": "Obtains the position of the enemy ship and deciphers the message received on the satellites",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TopSecretSplit"
                ],
                "summary": "Get position and message for a specific satellite",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Nombre del Satélite",
                        "name": "satellite_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ResponseTopSecret"
                        }
                    }
                }
            },
            "post": {
                "description": "Obtains the position of the enemy ship and deciphers the message received on the satellites",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TopSecretSplit"
                ],
                "summary": "Get position and message for a specific satellite",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Nombre del Satélite",
                        "name": "satellite_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Distancia de la nave enemiga",
                        "name": "Data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.TopSecret"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ResponseTopSecret"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.Entry": {
            "type": "object",
            "properties": {
                "distance": {
                    "type": "number"
                },
                "message": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.Position": {
            "type": "object",
            "properties": {
                "x": {
                    "type": "number"
                },
                "y": {
                    "type": "number"
                }
            }
        },
        "dto.ResponseTopSecret": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "position": {
                    "$ref": "#/definitions/dto.Position"
                }
            }
        },
        "dto.TopSecret": {
            "type": "object",
            "properties": {
                "satellites": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Entry"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "API Operación Fuego de Quasar",
	Description:      "Esta API corresponde al desarrollo del desafío tecnico a desarrollar en GO",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}