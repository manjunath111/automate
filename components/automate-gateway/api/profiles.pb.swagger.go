package api

func init() {
	Swagger.Add("profiles", `{
  "swagger": "2.0",
  "info": {
    "title": "external/compliance/profiles/profiles.proto",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/api/v0/compliance/market/read/{name}/version/{version}": {
      "get": {
        "summary": "Show an available profile",
        "description": "Show the details of an un-installed profile using the profile name and version.\nin the UI, these are the profiles under the \"Available\" tab.\nThese profiles are created and maintained by Chef, shipped with Chef Automate.\n\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ncompliance:marketProfiles:get\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "ProfilesService_ReadFromMarket",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Profile"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "name",
            "description": "Name of the profile.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "version",
            "description": "Version of the profile.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "owner",
            "description": "Automate user associated with the profile.",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "ProfilesService"
        ]
      }
    },
    "/api/v0/compliance/profiles/metasearch": {
      "post": {
        "summary": "Check if one or multiple profiles exist in the metadata database.",
        "description": "The endpoint takes an array of compliance profile sha256 IDs and returns the ones that the backend\ndoesn't have metadata (profile title, copyright, controls title, code, tags, etc) for.\nThis is useful when deciding if a compliance report can be sent for ingestion without the associated profile metadata.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ncompliance:profiles:list\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "ProfilesService_MetaSearch",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Missing"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Sha256"
            }
          }
        ],
        "tags": [
          "ProfilesService"
        ]
      }
    },
    "/api/v0/compliance/profiles/read/{owner}/{name}/version/{version}": {
      "get": {
        "summary": "Show an installed profile",
        "description": "Show the details of an installed profile given the profile name, owner (Automate user associated with the profile), and version.\n\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ncompliance:profiles:get\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "ProfilesService_Read",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Profile"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "owner",
            "description": "Automate user associated with the profile.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Name of the profile.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "version",
            "description": "Version of the profile.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "ProfilesService"
        ]
      }
    },
    "/api/v0/compliance/profiles/search": {
      "post": {
        "summary": "List all available profiles",
        "description": "Lists all profiles available for the Automate instance.\nEmpty params return all \"market\" profiles.\nSpecifying the ` + "`" + `owner` + "`" + ` field returns all profiles installed for the specified user.\n\nSupports pagination, sorting, and filtering (wildcard supported).\n\nSupported sort fields: title, name (default: title)\nSupported filter fields: name, version, title\n\nExample:\n` + "`" + `` + "`" + `` + "`" + `\n{\n\"filters\":[\n{\"type\": \"title\", \"values\": [ \"Dev*\"]}\n],\n\"page\": 1,\n\"per_page\": 3,\n\"owner\": \"admin\"\n}\n` + "`" + `` + "`" + `` + "`" + `\n\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ncompliance:profiles:list\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "ProfilesService_List",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Profiles"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Query"
            }
          }
        ],
        "tags": [
          "ProfilesService"
        ]
      }
    },
    "/api/v0/compliance/profiles/{owner}/{name}/version/{version}": {
      "delete": {
        "summary": "Delete an installed profile",
        "description": "Delete an installed profile given the profile name, owner (Automate user associated with the profile), and version.\nNote: this action \"uninstalls\" the profile. This has no impact on the market profiles.\n\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ncompliance:profiles:delete\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "ProfilesService_Delete",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "properties": {}
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "owner",
            "description": "Automate user associated with the profile.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Name of the profile.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "version",
            "description": "Version of the profile.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "ProfilesService"
        ]
      }
    }
  },
  "definitions": {
    "chef.automate.api.compliance.profiles.v1.Attribute": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "options": {
          "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Option"
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.CheckMessage": {
      "type": "object",
      "properties": {
        "file": {
          "type": "string",
          "description": "Profile file where the error or warning exists."
        },
        "line": {
          "type": "integer",
          "format": "int32",
          "description": "Profile line where the error or warning exists."
        },
        "column": {
          "type": "integer",
          "format": "int32",
          "description": "Column where the error or warning exists."
        },
        "control_id": {
          "type": "string",
          "description": "Control ID associated with the error or warning."
        },
        "msg": {
          "type": "string",
          "description": "Message associated with the error or warning."
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.CheckResult": {
      "type": "object",
      "properties": {
        "summary": {
          "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.ResultSummary",
          "description": "Intentionally blank."
        },
        "errors": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.CheckMessage"
          },
          "description": "Errors returned by the ` + "`" + `inspec check` + "`" + ` command."
        },
        "warnings": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.CheckMessage"
          },
          "description": "Warnings returned by the ` + "`" + `inspec check` + "`" + ` command."
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.Chunk": {
      "type": "object",
      "properties": {
        "data": {
          "type": "string",
          "format": "byte"
        },
        "position": {
          "type": "string",
          "format": "int64"
        }
      },
      "description": "Profile contents in byte form."
    },
    "chef.automate.api.compliance.profiles.v1.Control": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "The ID of the control."
        },
        "code": {
          "type": "string",
          "description": "The code (test) for the control."
        },
        "desc": {
          "type": "string",
          "description": "The description of the control."
        },
        "impact": {
          "type": "number",
          "format": "float",
          "description": "The impact of the control."
        },
        "title": {
          "type": "string",
          "description": "The title of the control."
        },
        "source_location": {
          "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.SourceLocation",
          "description": "Intentionally blank."
        },
        "results": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Result"
          },
          "description": "The results of the control tests."
        },
        "refs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Ref"
          },
          "description": "The refs associated with the control."
        },
        "tags": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "description": "The tags associated with the control."
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.Dependency": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Name of the profile."
        },
        "url": {
          "type": "string",
          "description": "URL of the profile."
        },
        "path": {
          "type": "string",
          "description": "Path of the profile."
        },
        "git": {
          "type": "string",
          "description": "Git location of the profile."
        },
        "branch": {
          "type": "string",
          "description": "Branch of the profile."
        },
        "tag": {
          "type": "string",
          "description": "Tag associated with the profile."
        },
        "commit": {
          "type": "string",
          "description": "Commit sha for the profile."
        },
        "version": {
          "type": "string",
          "description": "Version of the profile."
        },
        "supermarket": {
          "type": "string",
          "description": "Supermarket address of the profile."
        },
        "github": {
          "type": "string",
          "description": "Github address of the profile."
        },
        "compliance": {
          "type": "string",
          "description": "Automate address of the profile."
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.Group": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "title": {
          "type": "string"
        },
        "controls": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.ListFilter": {
      "type": "object",
      "properties": {
        "values": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of values to filter on."
        },
        "type": {
          "type": "string",
          "description": "The field to filter on."
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.Metadata": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "title": "Name of the profile (as specified in the inspec.yml)"
        },
        "version": {
          "type": "string",
          "description": "Version of the profile."
        },
        "content_type": {
          "type": "string",
          "title": "Content type of the profile (e.g. application/json, application/x-gtar, application/gzip)"
        }
      },
      "description": "Metadata about the profile."
    },
    "chef.automate.api.compliance.profiles.v1.Missing": {
      "type": "object",
      "properties": {
        "missing_sha256": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "An array of profile sha256 IDs that are missing from the backend metadata store."
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.Option": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string"
        },
        "default": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.Profile": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "title": "The profile name, as specified in the inspec.yml"
        },
        "title": {
          "type": "string",
          "title": "The profile title, as specified in the inspec.yml"
        },
        "maintainer": {
          "type": "string",
          "title": "The profile maintainer, as specified in the inspec.yml"
        },
        "copyright": {
          "type": "string",
          "title": "The profile copyright, as specified in the inspec.yml"
        },
        "copyright_email": {
          "type": "string",
          "title": "The profile copyright email, as specified in the inspec.yml"
        },
        "license": {
          "type": "string",
          "title": "The profile license, as specified in the inspec.yml"
        },
        "summary": {
          "type": "string",
          "title": "The profile summary, as specified in the inspec.yml"
        },
        "version": {
          "type": "string",
          "title": "The profile version, as specified in the inspec.yml"
        },
        "owner": {
          "type": "string",
          "description": "The Automate user associated with the profile."
        },
        "supports": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Support"
          },
          "title": "The list of operating systems compatible with the profile, as specified in the inspec.yml"
        },
        "depends": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Dependency"
          },
          "title": "The list of dependencies the profile has, as specified in the inspec.yml"
        },
        "sha256": {
          "type": "string",
          "description": "The SHA256 of the profile."
        },
        "groups": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Group"
          }
        },
        "controls": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Control"
          },
          "description": "The list of controls in the profile."
        },
        "attributes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Attribute"
          },
          "description": "The list of attributes in the profile."
        },
        "latest_version": {
          "type": "string",
          "description": "The latest version of the profile."
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.ProfileData": {
      "type": "object",
      "properties": {
        "owner": {
          "type": "string",
          "description": "Automate user associated with the profile."
        },
        "name": {
          "type": "string",
          "description": "Name of the profile."
        },
        "version": {
          "type": "string",
          "description": "Version of the profile."
        },
        "data": {
          "type": "string",
          "format": "byte",
          "description": "Profile contents in byte form."
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.Profiles": {
      "type": "object",
      "properties": {
        "profiles": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Profile"
          },
          "description": "List of profiles matching the query."
        },
        "total": {
          "type": "integer",
          "format": "int32",
          "description": "Total count of profiles matching the query."
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.Query": {
      "type": "object",
      "properties": {
        "filters": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.ListFilter"
          },
          "title": "Filters to apply to the query"
        },
        "order": {
          "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Query.OrderType",
          "description": "Order in which to sort. Defaults to ASC."
        },
        "sort": {
          "type": "string",
          "description": "Field on which to sort."
        },
        "page": {
          "type": "integer",
          "format": "int32",
          "description": "Page of results requested."
        },
        "per_page": {
          "type": "integer",
          "format": "int32",
          "description": "Number of results to return per page."
        },
        "owner": {
          "type": "string",
          "description": "Automate user associated with the profile."
        },
        "name": {
          "type": "string",
          "description": "Name of the profile (as defined in ` + "`" + `inspec.yml` + "`" + `)."
        },
        "version": {
          "type": "string",
          "description": "Version of the profile (as defined in ` + "`" + `inspec.yml` + "`" + `)."
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.Query.OrderType": {
      "type": "string",
      "enum": [
        "ASC",
        "DESC"
      ],
      "default": "ASC"
    },
    "chef.automate.api.compliance.profiles.v1.Ref": {
      "type": "object",
      "properties": {
        "url": {
          "type": "string",
          "description": "URL of the ref."
        },
        "ref": {
          "type": "string",
          "description": "Ref for the control."
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.Result": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string",
          "description": "Status of the test results (passed, failed, skipped)."
        },
        "code_desc": {
          "type": "string",
          "description": "The code (test) executed."
        },
        "run_time": {
          "type": "number",
          "format": "float",
          "description": "The amount of time it took to execute the test."
        },
        "start_time": {
          "type": "string",
          "description": "The time the test started."
        },
        "message": {
          "type": "string",
          "description": "The failure message."
        },
        "skip_message": {
          "type": "string",
          "description": "Reason for skipping the test."
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.ResultSummary": {
      "type": "object",
      "properties": {
        "valid": {
          "type": "boolean",
          "description": "Boolean that denotes if the profile is valid or not (as reported by ` + "`" + `inspec check` + "`" + `)."
        },
        "timestamp": {
          "type": "string",
          "description": "Timestamp of when the ` + "`" + `inspec check` + "`" + ` command was executed."
        },
        "location": {
          "type": "string",
          "description": "Path of the checked profile."
        },
        "controls": {
          "type": "integer",
          "format": "int32",
          "description": "Count of controls in the profile."
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.Sha256": {
      "type": "object",
      "properties": {
        "sha256": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "An array of profile sha256 IDs."
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.SourceLocation": {
      "type": "object",
      "properties": {
        "ref": {
          "type": "string"
        },
        "line": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.Support": {
      "type": "object",
      "properties": {
        "os_name": {
          "type": "string",
          "description": "OS name supported by the profile."
        },
        "os_family": {
          "type": "string",
          "description": "OS family supported by the profile."
        },
        "release": {
          "type": "string",
          "description": "OS release supported by the profile."
        },
        "inspec_version": {
          "type": "string",
          "description": "Minimum InSpec version required for the profile."
        },
        "platform": {
          "type": "string",
          "description": "Platform supported by the profile."
        }
      }
    },
    "google.protobuf.Any": {
      "type": "object",
      "properties": {
        "type_url": {
          "type": "string"
        },
        "value": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "grpc.gateway.runtime.Error": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "details": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/google.protobuf.Any"
          }
        }
      }
    },
    "grpc.gateway.runtime.StreamError": {
      "type": "object",
      "properties": {
        "grpc_code": {
          "type": "integer",
          "format": "int32"
        },
        "http_code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "http_status": {
          "type": "string"
        },
        "details": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/google.protobuf.Any"
          }
        }
      }
    }
  }
}
`)
}
