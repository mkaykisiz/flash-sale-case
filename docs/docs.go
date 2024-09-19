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
        "/api/v1/flash-sales": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "FlashSale"
                ],
                "summary": "Get flash sales",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.GetFlashSalesResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "FlashSale"
                ],
                "summary": "Add Flash Sale",
                "parameters": [
                    {
                        "description": "Add flash sale request body",
                        "name": "add_flash_sale",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.AddFlashSaleRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.AddFlashSaleResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/flash-sales/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "FlashSale"
                ],
                "summary": "Get a single flash sale",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.GetFlashSaleResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "FlashSale"
                ],
                "summary": "Edit Flash Sale",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Flash Sale ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Edit flash sale request body",
                        "name": "edit_flash_sale",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.EditFlashSaleRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "FlashSale"
                ],
                "summary": "Delete flash sale",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/flash-sales/{id}/buy": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "FlashSale"
                ],
                "summary": "Buy Flash Sale",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Buy flash sale request body",
                        "name": "buy_flash_sale",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.BuyFlashSaleRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/products": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Add product",
                "parameters": [
                    {
                        "description": "Add product request body",
                        "name": "add_product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.AddProductData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/products/{id}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Update product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Edit product request body",
                        "name": "edit_product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.EditProductData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/auth": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Get Auth",
                "parameters": [
                    {
                        "description": "Auth request body",
                        "name": "auth",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.Auth"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.AddFlashSaleRequest": {
            "type": "object",
            "properties": {
                "discount_percent": {
                    "type": "integer"
                },
                "end_time": {
                    "type": "string"
                },
                "product_id": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "string"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "app.AddFlashSaleResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "app.AddProductData": {
            "type": "object",
            "properties": {
                "is_active": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "app.Auth": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "app.BuyFlashSaleRequest": {
            "type": "object",
            "properties": {
                "unit": {
                    "type": "integer"
                }
            }
        },
        "app.EditFlashSaleRequest": {
            "type": "object",
            "properties": {
                "discount_percent": {
                    "type": "integer"
                },
                "end_time": {
                    "type": "string"
                },
                "start_time": {
                    "type": "string"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "app.EditProductData": {
            "type": "object",
            "properties": {
                "is_active": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "app.FlashSale": {
            "type": "object",
            "properties": {
                "discount_percent": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "product_id": {
                    "type": "integer"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "app.GetFlashSaleResponse": {
            "type": "object",
            "properties": {
                "discount_percent": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "product_id": {
                    "type": "integer"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "app.GetFlashSalesResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "flash_sales": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/app.FlashSale"
                    }
                }
            }
        },
        "app.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "token",
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
	Title:            "Flash Sale",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
