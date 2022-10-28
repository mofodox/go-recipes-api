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
        "contact": {
            "name": "Khairul Akmal",
            "url": "https://mofodox.com",
            "email": "kai@mofodox.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/recipes": {
            "get": {
                "description": "Returns a list of recipes",
                "produces": [
                    "application/json"
                ],
                "summary": "Returns a list of recipes",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Recipe"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new recipe",
                "produces": [
                    "application/json"
                ],
                "summary": "To create a new recipe",
                "parameters": [
                    {
                        "description": "recipe data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Recipe"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.Recipe"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/recipes/:id": {
            "put": {
                "description": "Updates an existing recipe",
                "produces": [
                    "application/json"
                ],
                "summary": "Updates an existing recipe",
                "parameters": [
                    {
                        "description": "recipe data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Recipe"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Recipe"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes an existing recipe",
                "produces": [
                    "application/json"
                ],
                "summary": "Deletes an existing recipe",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/recipes/search": {
            "get": {
                "description": "Searches for recipes by tags",
                "produces": [
                    "application/json"
                ],
                "summary": "Searches for recipes by tags",
                "parameters": [
                    {
                        "type": "string",
                        "format": "tag",
                        "description": "recipe search by tags",
                        "name": "tag",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Recipe": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "ingredients": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "instructions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "publishedAt": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{"http"},
	Title:            "Recipes API",
	Description:      "This is a project on building applications using Gin",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
