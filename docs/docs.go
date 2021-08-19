// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "accelerite.com",
        "contact": {
            "name": "API Support",
            "url": "http://accelerite.com/support"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Show the status of server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/workspace/create": {
            "post": {
                "description": "creates workspace diretory",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "workspaces"
                ],
                "summary": "Creates a WorkSpace",
                "parameters": [
                    {
                        "description": "Add workspace",
                        "name": "workpace",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.WorkSpace"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/workspace/{UUID}": {
            "get": {
                "description": "get Workpace by UUID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "workspaces"
                ],
                "summary": "Show an workspace",
                "parameters": [
                    {
                        "type": "string",
                        "description": "worskspace ID",
                        "name": "UUID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Found workspace",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/workspace/{uuid}": {
            "delete": {
                "description": "Delete by Workspace UUID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "workspaces"
                ],
                "summary": "Delete a Workspace",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Workspace ID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Succssfully Deleted Workspace",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/workspace/{uuid}/add/resource/{resourceName}": {
            "post": {
                "description": "Enter WorkspaceUUID and resource Name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "workspaces"
                ],
                "summary": "Add resource to workspace",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Workspace ID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Resource Name",
                        "name": "resourceName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully added resource to workspace",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/workspace/{uuid}/add/template/": {
            "post": {
                "description": "Enter WorkspaceUUID and template Name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "workspaces"
                ],
                "summary": "Add template to workspace",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Workspace UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Template UUID",
                        "name": "templateUUID",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully added resource to workspace",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/worskpace/{UUID}": {
            "patch": {
                "description": "Update by json Workspace",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "workspaces"
                ],
                "summary": "Update a Workspace",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Workspace UUID",
                        "name": "UUID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update workspace",
                        "name": "workpace",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.WorkSpace"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated Workspace Successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "models.Organization": {
            "type": "object",
            "properties": {
                "attributes": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "account name"
                },
                "updatedAt": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string",
                    "format": "uuid",
                    "example": "workspace-550e8400"
                }
            }
        },
        "models.WorkSpace": {
            "type": "object",
            "required": [
                "OrganizationID",
                "name"
            ],
            "properties": {
                "OrganizationID": {
                    "description": "Required: true\nMin Length: 1",
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                },
                "attributes": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer",
                    "format": "int64",
                    "example": 2
                },
                "inputVariables": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "name": {
                    "description": "Required: true\nMin Length: 1",
                    "type": "string",
                    "example": "account name"
                },
                "organization": {
                    "$ref": "#/definitions/models.Organization"
                },
                "outputVariable": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "state": {
                    "type": "string",
                    "example": "creating"
                },
                "updatedAt": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string",
                    "format": "uuid",
                    "example": "workspace-550e8400"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8090",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Swagger Multicloud service API",
	Description: "This is multicloud server.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
