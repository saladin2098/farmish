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
        "/animal": {
            "post": {
                "description": "Create a new animal",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Animal"
                ],
                "summary": "Create Animal",
                "operationId": "create_animal",
                "parameters": [
                    {
                        "description": "Animal data",
                        "name": "animal",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AnimalCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Animal created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/animal/{id}": {
            "get": {
                "description": "Get an animal by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Animal"
                ],
                "summary": "Get Animal",
                "operationId": "get_animal",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Animal ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Animal data",
                        "schema": {
                            "$ref": "#/definitions/models.Animal"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an animal's information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Animal"
                ],
                "summary": "Update Animal",
                "operationId": "update_animal",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Animal ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Animal data",
                        "name": "animal",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AnimalUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Animal updated",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/animals": {
            "get": {
                "description": "Get all animals with optional filters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Animal"
                ],
                "summary": "Get All Animals",
                "operationId": "get_all_animals",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Animal type",
                        "name": "type",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Is Healthy",
                        "name": "is_healthy",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Is Hungry",
                        "name": "is_hungry",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Animals data",
                        "schema": {
                            "$ref": "#/definitions/models.AnimalsGetAll"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/provision": {
            "get": {
                "description": "Get All Provisions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Provision"
                ],
                "summary": "Get All Provisions",
                "operationId": "get_all_provisions",
                "responses": {
                    "200": {
                        "description": "Provisions data",
                        "schema": {
                            "$ref": "#/definitions/models.GetAllProvisions"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create Provision",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Provision"
                ],
                "summary": "Create Provision",
                "operationId": "create_provision",
                "parameters": [
                    {
                        "description": "Created Provision",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BodyProvision"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Provision data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/provision/": {
            "put": {
                "description": "Update a provision",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Provision"
                ],
                "summary": "Update Provision",
                "operationId": "update_provision",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Provision ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Provision data",
                        "name": "provision",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateProvision"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Provision updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/provision/{id}": {
            "delete": {
                "description": "Delete a provision by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Provision"
                ],
                "summary": "Delete Provision",
                "operationId": "delete_provision",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Provision ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Provision deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/rovision/{id}{type}{animal_type}{quantity}": {
            "get": {
                "description": "Get a provision by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Provision"
                ],
                "summary": "Get Provision By ID",
                "operationId": "get_provision_by_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Provision ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Provision data",
                        "schema": {
                            "$ref": "#/definitions/models.GetProvision"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Animal": {
            "type": "object",
            "properties": {
                "animalType": {
                    "type": "string"
                },
                "avgConsumption": {
                    "type": "number"
                },
                "avgWater": {
                    "type": "number"
                },
                "birth": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "integer"
                },
                "feeding": {
                    "$ref": "#/definitions/models.FeedingSchedule"
                },
                "healthCondition": {
                    "$ref": "#/definitions/models.HealthCondition"
                },
                "id": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "weight": {
                    "type": "integer"
                }
            }
        },
        "models.AnimalCreate": {
            "type": "object",
            "properties": {
                "animal_type": {
                    "type": "string"
                },
                "birth": {
                    "type": "string"
                },
                "condition": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_healthy": {
                    "type": "boolean"
                },
                "medication": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "weight": {
                    "type": "integer"
                }
            }
        },
        "models.AnimalGet": {
            "type": "object",
            "properties": {
                "animal_type": {
                    "type": "string"
                },
                "avg_consumption": {
                    "type": "number"
                },
                "avg_water": {
                    "type": "number"
                },
                "birth": {
                    "type": "string"
                },
                "healthCondition": {
                    "$ref": "#/definitions/models.HealthConditionGet"
                },
                "id": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                },
                "weight": {
                    "type": "number"
                }
            }
        },
        "models.AnimalUpdate": {
            "type": "object",
            "properties": {
                "condition": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_healthy": {
                    "type": "boolean"
                },
                "medication": {
                    "type": "string"
                },
                "weight": {
                    "type": "integer"
                }
            }
        },
        "models.AnimalsGetAll": {
            "type": "object",
            "properties": {
                "animals": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.AnimalGet"
                    }
                },
                "count": {
                    "type": "integer"
                }
            }
        },
        "models.BodyProvision": {
            "type": "object",
            "properties": {
                "animalType": {
                    "type": "string"
                },
                "quantity": {
                    "type": "number"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "models.FeedingSchedule": {
            "type": "object",
            "properties": {
                "animalType": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastFedIndex": {
                    "type": "integer"
                },
                "nextFedIndex": {
                    "type": "integer"
                },
                "scheduleID": {
                    "type": "integer"
                }
            }
        },
        "models.GetAllProvisions": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "provisions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.GetProvision"
                    }
                }
            }
        },
        "models.GetProvision": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "number"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "models.HealthCondition": {
            "type": "object",
            "properties": {
                "animalID": {
                    "type": "integer"
                },
                "condition": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isHealthy": {
                    "type": "boolean"
                },
                "medication": {
                    "type": "string"
                }
            }
        },
        "models.HealthConditionGet": {
            "type": "object",
            "properties": {
                "condition": {
                    "type": "string"
                },
                "isHealthy": {
                    "type": "boolean"
                },
                "medication": {
                    "type": "string"
                }
            }
        },
        "models.UpdateProvision": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "number"
                },
                "type": {
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}