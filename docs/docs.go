// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/peoples": {
            "put": {
                "description": "Modify People",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Peoples"
                ],
                "summary": "Modify People",
                "parameters": [
                    {
                        "description": "post values",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.People"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/structs.People"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "first_name": {
                                            "type": "string"
                                        },
                                        "last_name": {
                                            "type": "string"
                                        },
                                        "middle_name": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handlers.errorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Show people with sorting and filtering",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Peoples"
                ],
                "summary": "Show all people",
                "parameters": [
                    {
                        "description": "search val",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.Search"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/structs.People"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handlers.errorResponse"
                        }
                    }
                }
            }
        },
        "/peoples/add": {
            "post": {
                "description": "Add one people",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Peoples"
                ],
                "summary": "Add People",
                "parameters": [
                    {
                        "description": "post values",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.PeopleToAdd"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/structs.PeopleToAdd"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handlers.errorResponse"
                        }
                    }
                }
            }
        },
        "/peoples/{people_id}": {
            "post": {
                "description": "Show One People",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Peoples"
                ],
                "summary": "Show People By ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "people_id",
                        "name": "people_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.People"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handlers.errorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "DeletePeoplesById",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Peoples"
                ],
                "summary": "DeletePeoplesById",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "people_id",
                        "name": "people_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "people is delete",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handlers.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.errorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "structs.Filters": {
            "type": "object",
            "properties": {
                "column": {
                    "type": "string",
                    "example": "last_name"
                },
                "value": {
                    "type": "string",
                    "example": "Pushkin"
                }
            }
        },
        "structs.People": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": "Moscow"
                },
                "first_name": {
                    "type": "string",
                    "example": "Evgenij"
                },
                "id": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string",
                    "example": "Kolosov"
                },
                "middle_name": {
                    "type": "string",
                    "example": "Alexandrovich"
                }
            }
        },
        "structs.PeopleToAdd": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": "Moscow"
                },
                "first_name": {
                    "type": "string",
                    "example": "Evgenij"
                },
                "last_name": {
                    "type": "string",
                    "example": "Kolosov"
                },
                "middle_name": {
                    "type": "string",
                    "example": "Alexandrovich"
                }
            }
        },
        "structs.Search": {
            "type": "object",
            "properties": {
                "filters": {
                    "$ref": "#/definitions/structs.Filters"
                },
                "page": {
                    "type": "integer",
                    "example": 1
                },
                "perPage": {
                    "type": "integer",
                    "example": 5
                },
                "sorts": {
                    "$ref": "#/definitions/structs.Sorts"
                }
            }
        },
        "structs.Sorts": {
            "type": "object",
            "properties": {
                "sort": {
                    "type": "string",
                    "example": "p.last_name"
                },
                "way": {
                    "type": "string",
                    "example": "+"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8888",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "CRUD web Server",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}