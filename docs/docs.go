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
        "/v1/tasks": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "取得所有使用 cron 表達式排程的 Lambda 函數清單",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "列出由定時任務觸發的 Lambda 函數",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/entity.Task"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Task": {
            "type": "object",
            "properties": {
                "cronExpression": {
                    "description": "CronExpression 是定時表達式",
                    "type": "string"
                },
                "functionName": {
                    "description": "FunctionName 是 Lambda 函數名稱",
                    "type": "string"
                },
                "lastExecutionStatus": {
                    "description": "LastExecutionStatus 是最後執行狀態",
                    "allOf": [
                        {
                            "$ref": "#/definitions/entity.TaskStatus"
                        }
                    ]
                },
                "lastTriggeredTime": {
                    "description": "LastTriggeredTime 是最後執行時間",
                    "type": "string"
                }
            }
        },
        "entity.TaskStatus": {
            "type": "string",
            "enum": [
                "Success",
                "Failure"
            ],
            "x-enum-varnames": [
                "TaskStatusSuccess",
                "TaskStatusFailure"
            ]
        },
        "http.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
