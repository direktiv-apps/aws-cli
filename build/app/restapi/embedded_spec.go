// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Amazon Web Service (AWS) command line interface",
    "title": "aws-cli",
    "version": "1.0",
    "x-direktiv-meta": {
      "categories": [
        "cloud",
        "build"
      ],
      "container": "direktiv.azurecr.io/functions/aws-cli",
      "issues": "https://github.com/direktiv-apps/aws-cli/issues",
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "This function provides AWS's CLI version 2.7.18 and is based on the official [AWS CLI image](https://hub.docker.com/r/amazon/aws-cli) on Docker Hub.  The following additional packages are installed:\n- wget\n- git\n- curl",
      "maintainer": "[direktiv.io](https://www.direktiv.io) ",
      "url": "https://github.com/direktiv-apps/aws-cli"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "default": "development",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "default": "/tmp",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "access-key",
                "secret-key"
              ],
              "properties": {
                "access-key": {
                  "description": "AWS access key.",
                  "type": "string",
                  "example": "ABCABCABCDABCABCABCD"
                },
                "commands": {
                  "description": "Array of commands.",
                  "type": "array",
                  "items": {
                    "type": "object",
                    "properties": {
                      "command": {
                        "description": "Command to run",
                        "type": "string",
                        "example": "aws --version"
                      },
                      "continue": {
                        "description": "Stops excecution if command fails, otherwise proceeds with next command",
                        "type": "boolean"
                      },
                      "print": {
                        "description": "If set to false the command will not print the full command with arguments to logs.",
                        "type": "boolean",
                        "default": true
                      },
                      "silent": {
                        "description": "If set to false the command will not print output to logs.",
                        "type": "boolean",
                        "default": false
                      }
                    }
                  }
                },
                "files": {
                  "description": "File to create before running commands.",
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/direktivFile"
                  }
                },
                "region": {
                  "description": "Region the commands should be executed in.",
                  "type": "string",
                  "default": "us-east-1",
                  "example": "eu-central-1"
                },
                "secret-key": {
                  "description": "AWS secret key.",
                  "type": "string",
                  "example": "Abcd45sa01234+ThIsIsSuPeRsEcReT"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "List of executed commands.",
            "schema": {
              "type": "object",
              "properties": {
                "aws-cli": {
                  "type": "array",
                  "items": {
                    "type": "object",
                    "required": [
                      "success",
                      "result"
                    ],
                    "properties": {
                      "result": {
                        "additionalProperties": false
                      },
                      "success": {
                        "type": "boolean"
                      }
                    }
                  }
                }
              }
            },
            "examples": {
              "aws-cli": [
                {
                  "result": "VTQ3U....c2ZaN0FJaldjVnkra2tKV==",
                  "success": true
                },
                {
                  "result": "exit status 254",
                  "success": false
                }
              ]
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "Direktiv-ErrorCode": {
                "type": "string"
              },
              "Direktiv-ErrorMessage": {
                "type": "string"
              }
            }
          }
        },
        "x-direktiv": {
          "cmds": [
            {
              "action": "foreach",
              "continue": "{{ .Item.Continue }}",
              "env": [
                "AWS_ACCESS_KEY_ID={{ .Body.AccessKey }}",
                "AWS_SECRET_ACCESS_KEY={{ .Body.SecretKey }}",
                "AWS_DEFAULT_REGION={{ default \"us-east-1\" .Body.Region }}",
                "AWS_DEFAULT_OUTPUT=json"
              ],
              "exec": "{{ .Item.Command }}",
              "loop": ".Commands",
              "print": "{{ .Item.Print }}",
              "silent": "{{ .Item.Silent }}"
            }
          ],
          "output": "{\n  \"aws-cli\": {{ index . 0 | toJson }}\n}\n"
        },
        "x-direktiv-errors": {
          "io.direktiv.command.error": "Command execution failed",
          "io.direktiv.output.error": "Template error for output generation of the service",
          "io.direktiv.ri.error": "Can not create information object from request"
        },
        "x-direktiv-examples": [
          {
            "content": "- id: aws-cli\n  type: action\n  action:\n    function: aws-cli\n    secrets: [\"awsAccess\", \"awsSecret\", \"awsRegion\"]\n    input: \n      access-key: jq(.secrets.awsAccess)\n      secret-key: jq(.secrets.awsSecret)\n      region: jq(.secrets.awsRegion)\n      commands:\n      - command: aws ec2 describe-instances",
            "title": "Basic"
          }
        ],
        "x-direktiv-function": "functions:\n- id: aws-cli\n  image: direktiv.azurecr.io/functions/aws-cli:1.0\n  type: knative-workflow",
        "x-direktiv-secrets": [
          {
            "description": "AWS access key (IAM)",
            "name": "awsAccess"
          },
          {
            "description": "AWS secret key (IAM)",
            "name": "awsSecret"
          },
          {
            "description": "AWS region where the commands run",
            "name": "awsRegion"
          }
        ]
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "description": "On cancel Direktiv sends a DELETE request to\nthe action with id in the header\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "x-direktiv": {
          "cancel": "echo 'cancel {{ .DirektivActionID }}'"
        }
      }
    }
  },
  "definitions": {
    "direktivFile": {
      "type": "object",
      "x-go-type": {
        "import": {
          "package": "github.com/direktiv/apps/go/pkg/apps"
        },
        "type": "DirektivFile"
      }
    },
    "error": {
      "type": "object",
      "required": [
        "errorCode",
        "errorMessage"
      ],
      "properties": {
        "errorCode": {
          "type": "string"
        },
        "errorMessage": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Amazon Web Service (AWS) command line interface",
    "title": "aws-cli",
    "version": "1.0",
    "x-direktiv-meta": {
      "categories": [
        "cloud",
        "build"
      ],
      "container": "direktiv.azurecr.io/functions/aws-cli",
      "issues": "https://github.com/direktiv-apps/aws-cli/issues",
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "This function provides AWS's CLI version 2.7.18 and is based on the official [AWS CLI image](https://hub.docker.com/r/amazon/aws-cli) on Docker Hub.  The following additional packages are installed:\n- wget\n- git\n- curl",
      "maintainer": "[direktiv.io](https://www.direktiv.io) ",
      "url": "https://github.com/direktiv-apps/aws-cli"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "default": "development",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "default": "/tmp",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/postParamsBody"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "List of executed commands.",
            "schema": {
              "$ref": "#/definitions/postOKBody"
            },
            "examples": {
              "aws-cli": [
                {
                  "result": "VTQ3U....c2ZaN0FJaldjVnkra2tKV==",
                  "success": true
                },
                {
                  "result": "exit status 254",
                  "success": false
                }
              ]
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "Direktiv-ErrorCode": {
                "type": "string"
              },
              "Direktiv-ErrorMessage": {
                "type": "string"
              }
            }
          }
        },
        "x-direktiv": {
          "cmds": [
            {
              "action": "foreach",
              "continue": "{{ .Item.Continue }}",
              "env": [
                "AWS_ACCESS_KEY_ID={{ .Body.AccessKey }}",
                "AWS_SECRET_ACCESS_KEY={{ .Body.SecretKey }}",
                "AWS_DEFAULT_REGION={{ default \"us-east-1\" .Body.Region }}",
                "AWS_DEFAULT_OUTPUT=json"
              ],
              "exec": "{{ .Item.Command }}",
              "loop": ".Commands",
              "print": "{{ .Item.Print }}",
              "silent": "{{ .Item.Silent }}"
            }
          ],
          "output": "{\n  \"aws-cli\": {{ index . 0 | toJson }}\n}\n"
        },
        "x-direktiv-errors": {
          "io.direktiv.command.error": "Command execution failed",
          "io.direktiv.output.error": "Template error for output generation of the service",
          "io.direktiv.ri.error": "Can not create information object from request"
        },
        "x-direktiv-examples": [
          {
            "content": "- id: aws-cli\n  type: action\n  action:\n    function: aws-cli\n    secrets: [\"awsAccess\", \"awsSecret\", \"awsRegion\"]\n    input: \n      access-key: jq(.secrets.awsAccess)\n      secret-key: jq(.secrets.awsSecret)\n      region: jq(.secrets.awsRegion)\n      commands:\n      - command: aws ec2 describe-instances",
            "title": "Basic"
          }
        ],
        "x-direktiv-function": "functions:\n- id: aws-cli\n  image: direktiv.azurecr.io/functions/aws-cli:1.0\n  type: knative-workflow",
        "x-direktiv-secrets": [
          {
            "description": "AWS access key (IAM)",
            "name": "awsAccess"
          },
          {
            "description": "AWS secret key (IAM)",
            "name": "awsSecret"
          },
          {
            "description": "AWS region where the commands run",
            "name": "awsRegion"
          }
        ]
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "description": "On cancel Direktiv sends a DELETE request to\nthe action with id in the header\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "x-direktiv": {
          "cancel": "echo 'cancel {{ .DirektivActionID }}'"
        }
      }
    }
  },
  "definitions": {
    "direktivFile": {
      "type": "object",
      "x-go-type": {
        "import": {
          "package": "github.com/direktiv/apps/go/pkg/apps"
        },
        "type": "DirektivFile"
      }
    },
    "error": {
      "type": "object",
      "required": [
        "errorCode",
        "errorMessage"
      ],
      "properties": {
        "errorCode": {
          "type": "string"
        },
        "errorMessage": {
          "type": "string"
        }
      }
    },
    "postOKBody": {
      "type": "object",
      "properties": {
        "aws-cli": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/postOKBodyAwsCliItems"
          }
        }
      },
      "x-go-gen-location": "operations"
    },
    "postOKBodyAwsCliItems": {
      "type": "object",
      "required": [
        "success",
        "result"
      ],
      "properties": {
        "result": {
          "additionalProperties": false
        },
        "success": {
          "type": "boolean"
        }
      },
      "x-go-gen-location": "operations"
    },
    "postParamsBody": {
      "type": "object",
      "required": [
        "access-key",
        "secret-key"
      ],
      "properties": {
        "access-key": {
          "description": "AWS access key.",
          "type": "string",
          "example": "ABCABCABCDABCABCABCD"
        },
        "commands": {
          "description": "Array of commands.",
          "type": "array",
          "items": {
            "$ref": "#/definitions/postParamsBodyCommandsItems"
          }
        },
        "files": {
          "description": "File to create before running commands.",
          "type": "array",
          "items": {
            "$ref": "#/definitions/direktivFile"
          }
        },
        "region": {
          "description": "Region the commands should be executed in.",
          "type": "string",
          "default": "us-east-1",
          "example": "eu-central-1"
        },
        "secret-key": {
          "description": "AWS secret key.",
          "type": "string",
          "example": "Abcd45sa01234+ThIsIsSuPeRsEcReT"
        }
      },
      "x-go-gen-location": "operations"
    },
    "postParamsBodyCommandsItems": {
      "type": "object",
      "properties": {
        "command": {
          "description": "Command to run",
          "type": "string",
          "example": "aws --version"
        },
        "continue": {
          "description": "Stops excecution if command fails, otherwise proceeds with next command",
          "type": "boolean"
        },
        "print": {
          "description": "If set to false the command will not print the full command with arguments to logs.",
          "type": "boolean",
          "default": true
        },
        "silent": {
          "description": "If set to false the command will not print output to logs.",
          "type": "boolean",
          "default": false
        }
      },
      "x-go-gen-location": "operations"
    }
  }
}`))
}
