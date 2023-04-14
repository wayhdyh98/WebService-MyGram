// Code generated by swaggo/swag. DO NOT EDIT.

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
            "email": "soberkoder@swagger.io"
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
        "/comments": {
            "get": {
                "description": "Get all comments from given user Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "Get all comments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new comment corresponding to the input data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "Create comment from given data",
                "parameters": [
                    {
                        "description": "create comment",
                        "name": "models.Comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            }
        },
        "/comments/{commentId}": {
            "get": {
                "description": "Get comment corresponding to the input Comment id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "Get comment for a given Comment id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the comment",
                        "name": "commentId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete the comment corresponding to the input Comment id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "Delete comment identified by the given Comment id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the comment to be deleted",
                        "name": "commentId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            },
            "patch": {
                "description": "Update the comment corresponding to the input Comment id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "Update comment identified by the given Comment id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the comment to be updated",
                        "name": "commentId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            }
        },
        "/medias": {
            "get": {
                "description": "Get all medias from given user Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "media"
                ],
                "summary": "Get all medias",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Socialmedia"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new media corresponding to the input data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "media"
                ],
                "summary": "Create media from given data",
                "parameters": [
                    {
                        "description": "create media",
                        "name": "models.Socialmedia",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Socialmedia"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Socialmedia"
                        }
                    }
                }
            }
        },
        "/medias/{mediaId}": {
            "get": {
                "description": "Get media corresponding to the input Media id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "media"
                ],
                "summary": "Get media for a given Media id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the media",
                        "name": "mediaId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Socialmedia"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete the media corresponding to the input Media id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "media"
                ],
                "summary": "Delete media identified by the given Media id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the media to be deleted",
                        "name": "mediaId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            },
            "patch": {
                "description": "Update the media corresponding to the input Media id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "media"
                ],
                "summary": "Update media identified by the given Media id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the media to be updated",
                        "name": "mediaId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Socialmedia"
                        }
                    }
                }
            }
        },
        "/photos": {
            "get": {
                "description": "Get all photos from given user Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photo"
                ],
                "summary": "Get all photos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new photo corresponding to the input data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photo"
                ],
                "summary": "Create photo from given data",
                "parameters": [
                    {
                        "description": "create photo",
                        "name": "models.Photo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                }
            }
        },
        "/photos/{photoId}": {
            "get": {
                "description": "Get photo corresponding to the input Photo id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photo"
                ],
                "summary": "Get photo for a given Photo id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the photo",
                        "name": "photoId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete the photo corresponding to the input Photo id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photo"
                ],
                "summary": "Delete photo identified by the given Photo id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the photo to be deleted",
                        "name": "photoId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            },
            "patch": {
                "description": "Update the photo corresponding to the input Photo id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photo"
                ],
                "summary": "Update photo identified by the given Photo id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the photo to be updated",
                        "name": "photoId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "Get user profile corresponding to the input Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get users for a given Id",
                "parameters": [
                    {
                        "description": "login user",
                        "name": "models.User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "Create new user corresponding to the input data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create users from given data",
                "parameters": [
                    {
                        "description": "create user",
                        "name": "models.User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Comment": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "photo": {
                    "$ref": "#/definitions/models.Photo"
                },
                "photo_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "models.Photo": {
            "type": "object",
            "properties": {
                "caption": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "photo_url": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "models.Socialmedia": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "social_media_url": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "comments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Comment"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "medsos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Socialmedia"
                    }
                },
                "photos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Photo"
                    }
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "4.0",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "MyGram API",
	Description:      "This is a simple service for managing MyGram",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
