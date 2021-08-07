// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package api

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://xingkong.io/support",
            "email": "zhenhua32@xingkong.io"
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
        "/book/{id}": {
            "get": {
                "description": "返回小说的详情",
                "consumes": [
                    "text/html"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "book"
                ],
                "summary": "返回小说的详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "小说ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Book"
                        }
                    }
                }
            }
        },
        "/book/{id}/directory": {
            "get": {
                "description": "返回小说的目录",
                "consumes": [
                    "text/html"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "book"
                ],
                "summary": "返回小说的目录",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "小说ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/book.GetBookDirectoryResp"
                        }
                    }
                }
            }
        },
        "/chapter/{id}": {
            "get": {
                "description": "返回章节的详情",
                "consumes": [
                    "text/html"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chapter"
                ],
                "summary": "返回章节的详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "章节ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Chapter"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "测试连接是否正常",
                "consumes": [
                    "text/html"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "监控"
                ],
                "summary": "测试连接是否正常",
                "operationId": "ping",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ping.PingResp"
                        }
                    }
                }
            }
        },
        "/search": {
            "post": {
                "description": "返回小说的搜索结果",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "搜索"
                ],
                "summary": "返回小说的搜索结果",
                "operationId": "search",
                "parameters": [
                    {
                        "description": "搜索参数",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/search.SearchReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/search.SearchResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "book.Chapter": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "章节名称",
                    "type": "string"
                },
                "url": {
                    "description": "链接",
                    "type": "string"
                }
            }
        },
        "book.GetBookDirectoryResp": {
            "type": "object",
            "properties": {
                "data_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Chapter"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "deleteat.DeletedAt": {
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
        "model.Book": {
            "type": "object",
            "properties": {
                "author": {
                    "description": "作者",
                    "type": "string"
                },
                "book_type": {
                    "description": "类型",
                    "type": "string"
                },
                "brief": {
                    "description": "简介",
                    "type": "string"
                },
                "chapter_list": {
                    "description": "定义 Has Many 关系",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Chapter"
                    }
                },
                "create_at": {
                    "type": "string"
                },
                "delete_at": {
                    "$ref": "#/definitions/deleteat.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "img_url": {
                    "description": "图片链接",
                    "type": "string"
                },
                "last_chapter": {
                    "description": "最近更新章节",
                    "$ref": "#/definitions/book.Chapter"
                },
                "last_update_time": {
                    "description": "最近更新时间",
                    "type": "string"
                },
                "name": {
                    "description": "书名",
                    "type": "string"
                },
                "source": {
                    "description": "来源",
                    "type": "string"
                },
                "update_at": {
                    "type": "string"
                },
                "url": {
                    "description": "链接",
                    "type": "string"
                }
            }
        },
        "model.Chapter": {
            "type": "object",
            "properties": {
                "book_id": {
                    "type": "integer"
                },
                "content": {
                    "type": "string"
                },
                "create_at": {
                    "type": "string"
                },
                "delete_at": {
                    "$ref": "#/definitions/deleteat.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "description": "章节名称",
                    "type": "string"
                },
                "update_at": {
                    "type": "string"
                },
                "url": {
                    "description": "链接",
                    "type": "string"
                }
            }
        },
        "ping.PingResp": {
            "type": "object",
            "properties": {
                "hello": {
                    "type": "string"
                }
            }
        },
        "search.SearchReq": {
            "type": "object",
            "required": [
                "keyword"
            ],
            "properties": {
                "keyword": {
                    "description": "搜索关键字",
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 1
                },
                "limit": {
                    "description": "限制结果数量",
                    "type": "integer",
                    "default": 10,
                    "maximum": 100,
                    "minimum": 1
                }
            }
        },
        "search.SearchResp": {
            "type": "object",
            "properties": {
                "data_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Book"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
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
	Host:        "127.0.0.1:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "行空 API",
	Description: "行空 API, 用于小说搜索",
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
