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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/order/": {
            "post": {
                "description": "Places an order on the market or a limit order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "order",
                "parameters": [
                    {
                        "description": "Market",
                        "name": "market",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Type of the order (MARKET or LIMIT)",
                        "name": "type",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Direction of the order (BUY или SELL)",
                        "name": "bid",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    {
                        "description": "Price (only for limit order)",
                        "name": "price",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Volume",
                        "name": "size",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/ping/": {
            "get": {
                "summary": "ping",
                "responses": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Qoinify API",
	Description:      "API Server for Qoinify application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}