package pb 

const (
swagger = `{
  "swagger": "2.0",
  "info": {
    "title": "pb/service.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/api/v1/contacts": {
      "get": {
        "operationId": "ReapContact",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbContactReqResp"
            }
          }
        },
        "tags": [
          "SimpleGRpcService"
        ]
      },
      "post": {
        "operationId": "CreateContact",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbContactReqResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbContactReqResp"
            }
          }
        ],
        "tags": [
          "SimpleGRpcService"
        ]
      }
    }
  },
  "definitions": {
    "RecipientResourceScope": {
      "type": "string",
      "enum": [
        "Cluster",
        "Namespaced"
      ],
      "default": "Cluster"
    },
    "pbContactReqResp": {
      "type": "object",
      "properties": {
        "recipe": {
          "$ref": "#/definitions/pbRecipient"
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "pbRecipient": {
      "type": "object",
      "properties": {
        "group": {
          "type": "string",
          "format": "string"
        },
        "kind": {
          "type": "string",
          "format": "string"
        },
        "plural": {
          "type": "string",
          "format": "string"
        },
        "resource_scope": {
          "$ref": "#/definitions/RecipientResourceScope"
        },
        "scope": {
          "type": "string",
          "format": "string"
        },
        "singular": {
          "type": "string",
          "format": "string"
        },
        "version": {
          "type": "string",
          "format": "string"
        }
      }
    }
  }
}
`
)
