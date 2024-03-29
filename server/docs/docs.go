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
        "/login": {
            "post": {
                "description": "Logs in a user with provided credentials.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "Login Credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LogInDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful login",
                        "schema": {
                            "$ref": "#/definitions/dto.LoggedInDTO"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "description": "Signs up a user with provided credentials.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "User signup",
                "parameters": [
                    {
                        "description": "Signup Credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SignUpDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful signup",
                        "schema": {
                            "$ref": "#/definitions/dto.LoggedInDTO"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.UserDTO"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create a new User",
                "parameters": [
                    {
                        "description": "User DTO",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UserDTO"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update a new User",
                "parameters": [
                    {
                        "description": "Update DTO",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateUserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UserDTO"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get a single User based on the provided ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UserDTO"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateUserDTO": {
            "type": "object",
            "required": [
                "fullname",
                "username"
            ],
            "properties": {
                "birthDate": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "introduction": {
                    "type": "string"
                },
                "lookingFor": {
                    "type": "string"
                },
                "photoUrl": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.LogInDTO": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.LoggedInDTO": {
            "type": "object",
            "required": [
                "fullname",
                "id",
                "username"
            ],
            "properties": {
                "birthDate": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "introduction": {
                    "type": "string"
                },
                "jwtToken": {
                    "type": "string"
                },
                "lookingFor": {
                    "type": "string"
                },
                "photos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.PhotoDTO"
                    }
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.PhotoDTO": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "isMain": {
                    "type": "boolean"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "dto.SignUpDTO": {
            "type": "object",
            "required": [
                "fullname",
                "password",
                "username"
            ],
            "properties": {
                "birthDate": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "introduction": {
                    "type": "string"
                },
                "lookingFor": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "photoUrl": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateUserDTO": {
            "type": "object",
            "properties": {
                "birthDate": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "introduction": {
                    "type": "string"
                },
                "lookingFor": {
                    "type": "string"
                }
            }
        },
        "dto.UserDTO": {
            "type": "object",
            "required": [
                "fullname",
                "id",
                "username"
            ],
            "properties": {
                "birthDate": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "introduction": {
                    "type": "string"
                },
                "lookingFor": {
                    "type": "string"
                },
                "photos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.PhotoDTO"
                    }
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "CitasApp API",
	Description:      "This is a HTTP API for CitasApp.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
