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
        "/blog-post": {
            "get": {
                "description": "Retrieve a list of all blog posts",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Blog"
                ],
                "summary": "Get all blog posts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.BlogPost"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new blog post with title, description, and body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Blog"
                ],
                "summary": "Create a new blog post",
                "parameters": [
                    {
                        "description": "Blog Post Data",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateBlogRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.BlogPost"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/blog-post/{id}": {
            "get": {
                "description": "Get details of a blog post by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Blog"
                ],
                "summary": "Get a single blog post",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Blog Post ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.BlogPost"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a blog post by ID",
                "tags": [
                    "Blog"
                ],
                "summary": "Delete a blog post",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Blog Post ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update a blog post's title, description, or body by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Blog"
                ],
                "summary": "Update a blog post",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Blog Post ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated Blog Post Data",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateBlogRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.BlogPost"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.BlogPost": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.CreateBlogRequest": {
            "type": "object",
            "required": [
                "body",
                "description",
                "title"
            ],
            "properties": {
                "body": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "models.UpdateBlogRequest": {
            "type": "object",
            "properties": {
                "body": {
                    "description": "Optional",
                    "type": "string"
                },
                "description": {
                    "description": "Optional",
                    "type": "string"
                },
                "title": {
                    "description": "Optional",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "assissment-xpx7.onrender.com",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Blog CRUD API",
	Description:      "Simple Blog API using Go-Fiber, PostgreSQL, and Swagger",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
