package rest
var SwaggerJson = `{
  "swagger": "2.0",
  "info": {
    "title": "Pydio Cells Rest API",
    "version": "1.0",
    "contact": {
      "name": "Pydio",
      "url": "https://pydio.com"
    }
  },
  "schemes": [
    "http",
    "https",
    "wss"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
"responses": {
    "401":{
      "description":"User is not authenticated",
      "schema":{
        "$ref": "#/definitions/restError"
      }
    },
    "403":{
      "description":"User has no permission to access this particular resource",
      "schema":{
        "$ref": "#/definitions/restError"
      }
    },
    "404":{
      "description":"Resource does not exist in the system",
      "schema":{
        "$ref": "#/definitions/restError"
      }
    },
    "500":{
      "description":"An internal error occurred in the backend",
      "schema":{
        "$ref": "#/definitions/restError"
      }
    }
  },
  "paths": {
    "/acl": {
      "post": {
        "summary": "Search Acls",
        "operationId": "SearchAcls",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restACLCollection"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restSearchACLRequest"
            }
          }
        ],
        "tags": [
          "ACLService"
        ]
      },
      "put": {
        "summary": "Store an ACL",
        "operationId": "PutAcl",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/idmACL"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/idmACL"
            }
          }
        ],
        "tags": [
          "ACLService"
        ]
      }
    },
    "/acl/bulk/delete": {
      "post": {
        "summary": "Delete one or more ACLs",
        "operationId": "DeleteAcl",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restDeleteResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/idmACL"
            }
          }
        ],
        "tags": [
          "ACLService"
        ]
      }
    },
    "/activity/stream": {
      "post": {
        "summary": "Load the the feeds of the currently logged user",
        "operationId": "Stream",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/activityObject"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/activityStreamActivitiesRequest"
            }
          }
        ],
        "tags": [
          "ActivityService"
        ]
      }
    },
    "/activity/subscribe": {
      "post": {
        "summary": "Manage subscriptions to other users/nodes feeds",
        "operationId": "Subscribe",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/activitySubscription"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/activitySubscription"
            }
          }
        ],
        "tags": [
          "ActivityService"
        ]
      }
    },
    "/activity/subscriptions": {
      "post": {
        "summary": "Load subscriptions to other users/nodes feeds",
        "operationId": "SearchSubscriptions",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restSubscriptionsCollection"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/activitySearchSubscriptionsRequest"
            }
          }
        ],
        "tags": [
          "ActivityService"
        ]
      }
    },
    "/auth/reset-password": {
      "post": {
        "summary": "Finish up the reset password process by providing the unique token",
        "operationId": "ResetPassword",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restResetPasswordResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restResetPasswordRequest"
            }
          }
        ],
        "tags": [
          "TokenService"
        ]
      }
    },
    "/auth/reset-password-token/{UserLogin}": {
      "put": {
        "summary": "Generate a unique token for the reset password process",
        "operationId": "ResetPasswordToken",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restResetPasswordTokenResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "UserLogin",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "TokenService"
        ]
      }
    },
    "/auth/token/revoke": {
      "post": {
        "summary": "Revoke a JWT token",
        "operationId": "Revoke",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restRevokeResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restRevokeRequest"
            }
          }
        ],
        "tags": [
          "TokenService"
        ]
      }
    },
    "/changes/{SeqID}": {
      "post": {
        "summary": "Get Changes",
        "operationId": "GetChanges",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restChangeCollection"
            }
          }
        },
        "parameters": [
          {
            "name": "SeqID",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "int64"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restChangeRequest"
            }
          }
        ],
        "tags": [
          "ChangeService"
        ]
      }
    },
    "/config/ctl": {
      "get": {
        "summary": "List all services and their status",
        "operationId": "ListServices",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restServiceCollection"
            }
          }
        },
        "parameters": [
          {
            "name": "StatusFilter",
            "in": "query",
            "required": false,
            "type": "string",
            "enum": [
              "ANY",
              "STOPPED",
              "STARTING",
              "STOPPING",
              "STARTED"
            ],
            "default": "ANY"
          }
        ],
        "tags": [
          "ConfigService"
        ]
      },
      "post": {
        "summary": "[Not Implemented]  Start/Stop a service",
        "operationId": "ControlService",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/ctlService"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restControlServiceRequest"
            }
          }
        ],
        "tags": [
          "ConfigService"
        ]
      }
    },
    "/config/datasource": {
      "get": {
        "summary": "List all defined datasources",
        "operationId": "ListDataSources",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restDataSourceCollection"
            }
          }
        },
        "tags": [
          "ConfigService"
        ]
      }
    },
    "/config/datasource/{Name}": {
      "get": {
        "summary": "Load datasource information",
        "operationId": "GetDataSource",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/objectDataSource"
            }
          }
        },
        "parameters": [
          {
            "name": "Name",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "Disabled",
            "in": "query",
            "required": false,
            "type": "boolean",
            "format": "boolean"
          },
          {
            "name": "StorageType",
            "in": "query",
            "required": false,
            "type": "string",
            "enum": [
              "LOCAL",
              "S3",
              "SMB",
              "CELLS",
              "AZURE",
              "GCS",
              "B2",
              "MANTA",
              "SIA"
            ],
            "default": "LOCAL"
          },
          {
            "name": "ObjectsServiceName",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "ObjectsHost",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "ObjectsPort",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "ObjectsSecure",
            "in": "query",
            "required": false,
            "type": "boolean",
            "format": "boolean"
          },
          {
            "name": "ObjectsBucket",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "ObjectsBaseFolder",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "ApiKey",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "ApiSecret",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "PeerAddress",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "Watch",
            "in": "query",
            "required": false,
            "type": "boolean",
            "format": "boolean"
          },
          {
            "name": "EncryptionMode",
            "in": "query",
            "required": false,
            "type": "string",
            "enum": [
              "CLEAR",
              "MASTER",
              "USER",
              "USER_PWD"
            ],
            "default": "CLEAR"
          },
          {
            "name": "EncryptionKey",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "VersioningPolicyName",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "CreationDate",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "LastSynchronizationDate",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          }
        ],
        "tags": [
          "ConfigService"
        ]
      },
      "delete": {
        "summary": "Delete a datasource",
        "operationId": "DeleteDataSource",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restDeleteDataSourceResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "Name",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "ConfigService"
        ]
      },
      "post": {
        "summary": "Create or update a datasource",
        "operationId": "PutDataSource",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/objectDataSource"
            }
          }
        },
        "parameters": [
          {
            "name": "Name",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/objectDataSource"
            }
          }
        ],
        "tags": [
          "ConfigService"
        ]
      }
    },
    "/config/discovery": {
      "get": {
        "summary": "Publish available endpoints",
        "operationId": "EndpointsDiscovery",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restDiscoveryResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "EndpointType",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "ConfigService"
        ]
      }
    },
    "/config/discovery/forms/{ServiceName}": {
      "get": {
        "summary": "Publish Forms definition for building screens in frontend",
        "operationId": "ConfigFormsDiscovery",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restDiscoveryResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "ServiceName",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "ConfigService"
        ]
      }
    },
    "/config/discovery/openapi": {
      "get": {
        "summary": "Publish available REST APIs",
        "operationId": "OpenApiDiscovery",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restOpenApiResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "EndpointType",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "ConfigService"
        ]
      }
    },
    "/config/encryption/create": {
      "post": {
        "summary": "Create a new master key",
        "operationId": "CreateEncryptionKey",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/encryptionAdminCreateKeyResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/encryptionAdminCreateKeyRequest"
            }
          }
        ],
        "tags": [
          "ConfigService"
        ]
      }
    },
    "/config/encryption/delete": {
      "post": {
        "summary": "Delete an existing master key",
        "operationId": "DeleteEncryptionKey",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/encryptionAdminDeleteKeyResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/encryptionAdminDeleteKeyRequest"
            }
          }
        ],
        "tags": [
          "ConfigService"
        ]
      }
    },
    "/config/encryption/export": {
      "post": {
        "summary": "Export a master key for backup purpose, protected with a password",
        "operationId": "ExportEncryptionKey",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/encryptionAdminExportKeyResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/encryptionAdminExportKeyRequest"
            }
          }
        ],
        "tags": [
          "ConfigService"
        ]
      }
    },
    "/config/encryption/import": {
      "put": {
        "summary": "Import a previously exported master key, requires the password created at export time",
        "operationId": "ImportEncryptionKey",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/encryptionAdminImportKeyResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/encryptionAdminImportKeyRequest"
            }
          }
        ],
        "tags": [
          "ConfigService"
        ]
      }
    },
    "/config/encryption/list": {
      "post": {
        "summary": "List registered master keys",
        "operationId": "ListEncryptionKeys",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/encryptionAdminListKeysResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/encryptionAdminListKeysRequest"
            }
          }
        ],
        "tags": [
          "ConfigService"
        ]
      }
    },
    "/config/peers": {
      "get": {
        "summary": "List all detected peers (servers on which the app is running)",
        "operationId": "ListPeersAddresses",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restListPeersAddressesResponse"
            }
          }
        },
        "tags": [
          "ConfigService"
        ]
      }
    },
    "/config/peers/{PeerAddress}": {
      "post": {
        "summary": "List folders on a peer, starting from root",
        "operationId": "ListPeerFolders",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restNodesCollection"
            }
          }
        },
        "parameters": [
          {
            "name": "PeerAddress",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restListPeerFoldersRequest"
            }
          }
        ],
        "tags": [
          "ConfigService"
        ]
      }
    },
    "/config/processes": {
      "post": {
        "summary": "List running Processes, with option PeerId or ServiceName filter",
        "operationId": "ListProcesses",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restListProcessesResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restListProcessesRequest"
            }
          }
        ],
        "tags": [
          "ConfigService"
        ]
      }
    },
    "/config/versioning": {
      "get": {
        "summary": "List all defined versioning policies",
        "operationId": "ListVersioningPolicies",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restVersioningPolicyCollection"
            }
          }
        },
        "tags": [
          "ConfigService"
        ]
      }
    },
    "/config/versioning/{Uuid}": {
      "get": {
        "summary": "Load a given versioning policy",
        "operationId": "GetVersioningPolicy",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/treeVersioningPolicy"
            }
          }
        },
        "parameters": [
          {
            "name": "Uuid",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "Name",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "Description",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "VersionsDataSourceName",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "VersionsDataSourceBucket",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "MaxTotalSize",
            "in": "query",
            "required": false,
            "type": "string",
            "format": "int64"
          },
          {
            "name": "MaxSizePerFile",
            "in": "query",
            "required": false,
            "type": "string",
            "format": "int64"
          },
          {
            "name": "IgnoreFilesGreaterThan",
            "in": "query",
            "required": false,
            "type": "string",
            "format": "int64"
          }
        ],
        "tags": [
          "ConfigService"
        ]
      }
    },
    "/config/virtualnodes": {
      "get": {
        "summary": "List all defined virtual nodes",
        "operationId": "ListVirtualNodes",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restNodesCollection"
            }
          }
        },
        "tags": [
          "ConfigService"
        ]
      }
    },
    "/config/{FullPath}": {
      "get": {
        "summary": "Generic config Get using a full path in the config tree",
        "operationId": "GetConfig",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restConfiguration"
            }
          }
        },
        "parameters": [
          {
            "name": "FullPath",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "Data",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "ConfigService"
        ]
      },
      "put": {
        "summary": "Generic config Put, using a full path in the config tree",
        "operationId": "PutConfig",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restConfiguration"
            }
          }
        },
        "parameters": [
          {
            "name": "FullPath",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restConfiguration"
            }
          }
        ],
        "tags": [
          "ConfigService"
        ]
      }
    },
    "/frontend/binaries/{BinaryType}/{Uuid}": {
      "get": {
        "summary": "Serve frontend binaries directly (avatars / logos / bg images)",
        "operationId": "FrontServeBinary",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restFrontBinaryResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "BinaryType",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "Uuid",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "FrontendService"
        ]
      },
      "post": {
        "summary": "Upload frontend binaries (avatars / logos / bg images)",
        "operationId": "FrontPutBinary",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restFrontBinaryResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "BinaryType",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "Uuid",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restFrontBinaryRequest"
            }
          }
        ],
        "tags": [
          "FrontendService"
        ]
      }
    },
    "/frontend/bootconf": {
      "get": {
        "summary": "Add some data to the initial set of parameters loaded by the frontend",
        "operationId": "FrontBootConf",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restFrontBootConfResponse"
            }
          }
        },
        "tags": [
          "FrontendService"
        ]
      }
    },
    "/frontend/enroll": {
      "post": {
        "summary": "Generic endpoint that can be implemented by 2FA systems for enrollment",
        "operationId": "FrontEnrollAuth",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restFrontEnrollAuthResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restFrontEnrollAuthRequest"
            }
          }
        ],
        "tags": [
          "FrontendService"
        ]
      }
    },
    "/frontend/messages/{Lang}": {
      "get": {
        "summary": "Serve list of I18n messages",
        "operationId": "FrontMessages",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restFrontMessagesResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "Lang",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "FrontendService"
        ]
      }
    },
    "/frontend/plugins/{Lang}": {
      "get": {
        "summary": "Serve list of I18n messages",
        "operationId": "FrontPlugins",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restFrontPluginsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "Lang",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "FrontendService"
        ]
      }
    },
    "/frontend/session": {
      "post": {
        "summary": "Handle JWT",
        "operationId": "FrontSession",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restFrontSessionResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restFrontSessionRequest"
            }
          }
        ],
        "tags": [
          "FrontendService"
        ]
      }
    },
    "/frontend/settings-menu": {
      "get": {
        "summary": "Sends a tree of nodes to be used a menu in the Settings panel",
        "operationId": "SettingsMenu",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restSettingsMenuResponse"
            }
          }
        },
        "tags": [
          "FrontendService"
        ]
      }
    },
    "/frontend/state": {
      "get": {
        "summary": "Send XML state registry",
        "operationId": "FrontState",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restFrontStateResponse"
            }
          }
        },
        "tags": [
          "FrontendService"
        ]
      }
    },
    "/graph/relation/{UserId}": {
      "get": {
        "summary": "Compute relation of context user with another user",
        "operationId": "Relation",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restRelationResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "UserId",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "GraphService"
        ]
      }
    },
    "/graph/state/{Segment}": {
      "get": {
        "summary": "Compute accessible workspaces for a given user",
        "operationId": "UserState",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restUserStateResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "Segment",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "GraphService"
        ]
      }
    },
    "/install": {
      "get": {
        "summary": "Loads default values for install form",
        "operationId": "GetInstall",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/installGetDefaultsResponse"
            }
          }
        },
        "tags": [
          "InstallService"
        ]
      },
      "post": {
        "summary": "Post values to be saved for install",
        "operationId": "PostInstall",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/installInstallResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/installInstallRequest"
            }
          }
        ],
        "tags": [
          "InstallService"
        ]
      }
    },
    "/install/agreement": {
      "get": {
        "summary": "Load a textual agreement for using the software",
        "operationId": "GetAgreement",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/installGetAgreementResponse"
            }
          }
        },
        "tags": [
          "InstallService"
        ]
      }
    },
    "/install/check": {
      "post": {
        "summary": "Perform a check during install (like a valid DB connection)",
        "operationId": "PerformInstallCheck",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/installPerformCheckResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/installPerformCheckRequest"
            }
          }
        ],
        "tags": [
          "InstallService"
        ]
      }
    },
    "/jobs/tasks/delete": {
      "post": {
        "summary": "Send a control command to clean tasks on a given job",
        "operationId": "UserDeleteTasks",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/jobsDeleteTasksResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/jobsDeleteTasksRequest"
            }
          }
        ],
        "tags": [
          "JobsService"
        ]
      }
    },
    "/jobs/tasks/logs": {
      "post": {
        "summary": "Technical Logs, in Json or CSV format",
        "operationId": "ListTasksLogs",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restLogMessageCollection"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/logListLogRequest"
            }
          }
        ],
        "tags": [
          "JobsService"
        ]
      }
    },
    "/jobs/user": {
      "post": {
        "summary": "List jobs associated with current user",
        "operationId": "UserListJobs",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restUserJobsCollection"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/jobsListJobsRequest"
            }
          }
        ],
        "tags": [
          "JobsService"
        ]
      },
      "put": {
        "summary": "Send Control Commands to one or many jobs / tasks",
        "operationId": "UserControlJob",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/jobsCtrlCommandResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/jobsCtrlCommand"
            }
          }
        ],
        "tags": [
          "JobsService"
        ]
      }
    },
    "/jobs/user/{JobName}": {
      "put": {
        "summary": "Create a predefined job to be run directly",
        "operationId": "UserCreateJob",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restUserJobResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "JobName",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restUserJobRequest"
            }
          }
        ],
        "tags": [
          "JobsService"
        ]
      }
    },
    "/log/sys": {
      "post": {
        "summary": "Technical Logs, in Json or CSV format",
        "operationId": "Syslog",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restLogMessageCollection"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/logListLogRequest"
            }
          }
        ],
        "tags": [
          "LogService"
        ]
      }
    },
    "/mailer/send": {
      "post": {
        "summary": "Send an email to a user or any email address",
        "operationId": "Send",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/mailerSendMailResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/mailerMail"
            }
          }
        ],
        "tags": [
          "MailerService"
        ]
      }
    },
    "/meta/bulk/get": {
      "post": {
        "summary": "List meta for a list of nodes, or a full directory using /path/* syntax",
        "operationId": "GetBulkMeta",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restBulkMetaResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restGetBulkMetaRequest"
            }
          }
        ],
        "tags": [
          "MetaService"
        ]
      }
    },
    "/meta/delete/{NodePath}": {
      "post": {
        "summary": "Delete metadata of a given node",
        "operationId": "DeleteMeta",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/treeNode"
            }
          }
        },
        "parameters": [
          {
            "name": "NodePath",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restMetaNamespaceRequest"
            }
          }
        ],
        "tags": [
          "MetaService"
        ]
      }
    },
    "/meta/get/{NodePath}": {
      "post": {
        "summary": "Load metadata for a given node",
        "operationId": "GetMeta",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/treeNode"
            }
          }
        },
        "parameters": [
          {
            "name": "NodePath",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restMetaNamespaceRequest"
            }
          }
        ],
        "tags": [
          "MetaService"
        ]
      }
    },
    "/meta/set/{NodePath}": {
      "post": {
        "summary": "Update metadata for a given node",
        "operationId": "SetMeta",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/treeNode"
            }
          }
        },
        "parameters": [
          {
            "name": "NodePath",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restMetaCollection"
            }
          }
        ],
        "tags": [
          "MetaService"
        ]
      }
    },
    "/policy": {
      "post": {
        "summary": "List all defined security policies",
        "operationId": "ListPolicies",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/idmListPolicyGroupsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/idmListPolicyGroupsRequest"
            }
          }
        ],
        "tags": [
          "PolicyService"
        ]
      }
    },
    "/role": {
      "post": {
        "summary": "Search Roles",
        "operationId": "SearchRoles",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restRolesCollection"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restSearchRoleRequest"
            }
          }
        ],
        "tags": [
          "RoleService"
        ]
      }
    },
    "/role/{Uuid}": {
      "get": {
        "summary": "Get a Role by ID",
        "operationId": "GetRole",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/idmRole"
            }
          }
        },
        "parameters": [
          {
            "name": "Uuid",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "Label",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "IsTeam",
            "in": "query",
            "required": false,
            "type": "boolean",
            "format": "boolean"
          },
          {
            "name": "GroupRole",
            "in": "query",
            "required": false,
            "type": "boolean",
            "format": "boolean"
          },
          {
            "name": "UserRole",
            "in": "query",
            "required": false,
            "type": "boolean",
            "format": "boolean"
          },
          {
            "name": "LastUpdated",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "AutoApplies",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            }
          },
          {
            "name": "PoliciesContextEditable",
            "in": "query",
            "required": false,
            "type": "boolean",
            "format": "boolean"
          },
          {
            "name": "ForceOverride",
            "in": "query",
            "required": false,
            "type": "boolean",
            "format": "boolean"
          }
        ],
        "tags": [
          "RoleService"
        ]
      },
      "delete": {
        "summary": "Delete a Role by ID",
        "operationId": "DeleteRole",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/idmRole"
            }
          }
        },
        "parameters": [
          {
            "name": "Uuid",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "RoleService"
        ]
      },
      "put": {
        "summary": "Create or update a Role",
        "operationId": "SetRole",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/idmRole"
            }
          }
        },
        "parameters": [
          {
            "name": "Uuid",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/idmRole"
            }
          }
        ],
        "tags": [
          "RoleService"
        ]
      }
    },
    "/search/nodes": {
      "post": {
        "summary": "Search indexed nodes (files/folders) on various aspects",
        "operationId": "Nodes",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restSearchResults"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/treeSearchRequest"
            }
          }
        ],
        "tags": [
          "SearchService"
        ]
      }
    },
    "/share/cell": {
      "put": {
        "summary": "Put or Create a share room",
        "operationId": "PutCell",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restCell"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restPutCellRequest"
            }
          }
        ],
        "tags": [
          "ShareService"
        ]
      }
    },
    "/share/cell/{Uuid}": {
      "get": {
        "summary": "Load a share room",
        "operationId": "GetCell",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restCell"
            }
          }
        },
        "parameters": [
          {
            "name": "Uuid",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "ShareService"
        ]
      },
      "delete": {
        "summary": "Delete a share room",
        "operationId": "DeleteCell",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restDeleteCellResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "Uuid",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "ShareService"
        ]
      }
    },
    "/share/link": {
      "put": {
        "summary": "Put or Create a share room",
        "operationId": "PutShareLink",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restShareLink"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restPutShareLinkRequest"
            }
          }
        ],
        "tags": [
          "ShareService"
        ]
      }
    },
    "/share/link/{Uuid}": {
      "get": {
        "summary": "Load a share link with all infos",
        "operationId": "GetShareLink",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restShareLink"
            }
          }
        },
        "parameters": [
          {
            "name": "Uuid",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "ShareService"
        ]
      },
      "delete": {
        "summary": "Delete Share Link",
        "operationId": "DeleteShareLink",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restDeleteShareLinkResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "Uuid",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "ShareService"
        ]
      }
    },
    "/share/policies": {
      "put": {
        "summary": "Updates policies associated to the underlying workspace for a Cell or a ShareLink",
        "operationId": "UpdateSharePolicies",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restUpdateSharePoliciesResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restUpdateSharePoliciesRequest"
            }
          }
        ],
        "tags": [
          "ShareService"
        ]
      }
    },
    "/share/resources": {
      "post": {
        "summary": "List Shared Resources for current user or all users",
        "operationId": "ListSharedResources",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restListSharedResourcesResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restListSharedResourcesRequest"
            }
          }
        ],
        "tags": [
          "ShareService"
        ]
      }
    },
    "/templates": {
      "get": {
        "summary": "List available templates",
        "operationId": "ListTemplates",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restListTemplatesResponse"
            }
          }
        },
        "tags": [
          "TemplatesService"
        ]
      }
    },
    "/tree/admin/list": {
      "post": {
        "summary": "List files and folders starting at the root (first level lists the datasources)",
        "operationId": "ListAdminTree",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restNodesCollection"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/treeListNodesRequest"
            }
          }
        ],
        "tags": [
          "AdminTreeService"
        ]
      }
    },
    "/tree/admin/stat": {
      "post": {
        "summary": "Read a node information inside the admin tree",
        "operationId": "StatAdminTree",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/treeReadNodeResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/treeReadNodeRequest"
            }
          }
        ],
        "tags": [
          "AdminTreeService"
        ]
      }
    },
    "/tree/create": {
      "post": {
        "summary": "Create dirs or empty files inside the tree",
        "operationId": "CreateNodes",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restNodesCollection"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restCreateNodesRequest"
            }
          }
        ],
        "tags": [
          "TreeService"
        ]
      }
    },
    "/tree/delete": {
      "post": {
        "summary": "Handle nodes deletion",
        "operationId": "DeleteNodes",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restDeleteNodesResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restDeleteNodesRequest"
            }
          }
        ],
        "tags": [
          "TreeService"
        ]
      }
    },
    "/tree/restore": {
      "post": {
        "summary": "Handle nodes restoration from recycle bin",
        "operationId": "RestoreNodes",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restRestoreNodesResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restRestoreNodesRequest"
            }
          }
        ],
        "tags": [
          "TreeService"
        ]
      }
    },
    "/tree/selection": {
      "post": {
        "summary": "Create a temporary selection for further action (namely download)",
        "operationId": "CreateSelection",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restCreateSelectionResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restCreateSelectionRequest"
            }
          }
        ],
        "tags": [
          "TreeService"
        ]
      }
    },
    "/tree/stat/{Node}": {
      "get": {
        "summary": "Return node meta without the node content itself",
        "operationId": "HeadNode",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restHeadNodeResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "Node",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "TreeService"
        ]
      }
    },
    "/tree/stats": {
      "post": {
        "summary": "List meta for a list of nodes, or a full directory using /path/* syntax",
        "operationId": "BulkStatNodes",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restBulkMetaResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restGetBulkMetaRequest"
            }
          }
        ],
        "tags": [
          "TreeService"
        ]
      }
    },
    "/update": {
      "post": {
        "summary": "Check the remote server to see if there are available binaries",
        "operationId": "UpdateRequired",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/updateUpdateResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/updateUpdateRequest"
            }
          }
        ],
        "tags": [
          "UpdateService"
        ]
      }
    },
    "/update/{TargetVersion}": {
      "patch": {
        "summary": "Apply an update to a given version",
        "operationId": "ApplyUpdate",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/updateApplyUpdateResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "TargetVersion",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/updateApplyUpdateRequest"
            }
          }
        ],
        "tags": [
          "UpdateService"
        ]
      }
    },
    "/user": {
      "post": {
        "summary": "List/Search users",
        "operationId": "SearchUsers",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restUsersCollection"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restSearchUserRequest"
            }
          }
        ],
        "tags": [
          "UserService"
        ]
      }
    },
    "/user-meta/bookmarks": {
      "post": {
        "summary": "Special API for Bookmarks, will load userMeta and the associated nodes, and return\nas a node list",
        "operationId": "UserBookmarks",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restBulkMetaResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restUserBookmarksRequest"
            }
          }
        ],
        "tags": [
          "UserMetaService"
        ]
      }
    },
    "/user-meta/namespace": {
      "get": {
        "summary": "List defined meta namespaces",
        "operationId": "ListUserMetaNamespace",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restUserMetaNamespaceCollection"
            }
          }
        },
        "tags": [
          "UserMetaService"
        ]
      },
      "put": {
        "summary": "Admin: update namespaces",
        "operationId": "UpdateUserMetaNamespace",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/idmUpdateUserMetaNamespaceResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/idmUpdateUserMetaNamespaceRequest"
            }
          }
        ],
        "tags": [
          "UserMetaService"
        ]
      }
    },
    "/user-meta/search": {
      "post": {
        "summary": "Search a list of meta by node Id or by User id and by namespace",
        "operationId": "SearchUserMeta",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restUserMetaCollection"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/idmSearchUserMetaRequest"
            }
          }
        ],
        "tags": [
          "UserMetaService"
        ]
      }
    },
    "/user-meta/tags/{Namespace}": {
      "get": {
        "summary": "List Tags for a given namespace",
        "operationId": "ListUserMetaTags",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restListUserMetaTagsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "Namespace",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "UserMetaService"
        ]
      },
      "post": {
        "summary": "Add a new value to Tags for a given namespace",
        "operationId": "PutUserMetaTag",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restPutUserMetaTagResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "Namespace",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restPutUserMetaTagRequest"
            }
          }
        ],
        "tags": [
          "UserMetaService"
        ]
      }
    },
    "/user-meta/tags/{Namespace}/{Tags}": {
      "delete": {
        "summary": "Delete one or all tags for a given namespace (use * for all tags)",
        "operationId": "DeleteUserMetaTags",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restDeleteUserMetaTagsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "Namespace",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "Tags",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "UserMetaService"
        ]
      }
    },
    "/user-meta/update": {
      "put": {
        "summary": "Update/delete user meta",
        "operationId": "UpdateUserMeta",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/idmUpdateUserMetaResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/idmUpdateUserMetaRequest"
            }
          }
        ],
        "tags": [
          "UserMetaService"
        ]
      }
    },
    "/user/roles/{Login}": {
      "put": {
        "summary": "Just save a user roles, without other datas",
        "operationId": "PutRoles",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/idmUser"
            }
          }
        },
        "parameters": [
          {
            "name": "Login",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/idmUser"
            }
          }
        ],
        "tags": [
          "UserService"
        ]
      }
    },
    "/user/{Login}": {
      "get": {
        "summary": "Get a user by login",
        "operationId": "GetUser",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/idmUser"
            }
          }
        },
        "parameters": [
          {
            "name": "Login",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "Uuid",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "GroupPath",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "Password",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "OldPassword",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "IsGroup",
            "description": "Group specific data.",
            "in": "query",
            "required": false,
            "type": "boolean",
            "format": "boolean"
          },
          {
            "name": "GroupLabel",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "PoliciesContextEditable",
            "in": "query",
            "required": false,
            "type": "boolean",
            "format": "boolean"
          }
        ],
        "tags": [
          "UserService"
        ]
      },
      "delete": {
        "summary": "Delete a user",
        "operationId": "DeleteUser",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restDeleteResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "Login",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "UserService"
        ]
      },
      "put": {
        "summary": "Create or update a user",
        "operationId": "PutUser",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/idmUser"
            }
          }
        },
        "parameters": [
          {
            "name": "Login",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/idmUser"
            }
          }
        ],
        "tags": [
          "UserService"
        ]
      }
    },
    "/workspace": {
      "post": {
        "summary": "Search workspaces on certain keys",
        "operationId": "SearchWorkspaces",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restWorkspaceCollection"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/restSearchWorkspaceRequest"
            }
          }
        ],
        "tags": [
          "WorkspaceService"
        ]
      }
    },
    "/workspace/{Slug}": {
      "delete": {
        "summary": "Delete an existing workspace",
        "operationId": "DeleteWorkspace",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/restDeleteResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "Slug",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "WorkspaceService"
        ]
      },
      "put": {
        "summary": "Create or update a workspace",
        "operationId": "PutWorkspace",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/idmWorkspace"
            }
          }
        },
        "parameters": [
          {
            "name": "Slug",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/idmWorkspace"
            }
          }
        ],
        "tags": [
          "WorkspaceService"
        ]
      }
    }
  },
  "definitions": {
    "ListLogRequestLogFormat": {
      "type": "string",
      "enum": [
        "JSON",
        "CSV",
        "XLSX"
      ],
      "default": "JSON",
      "title": "Output Format"
    },
    "ListSharedResourcesRequestListShareType": {
      "type": "string",
      "enum": [
        "ANY",
        "LINKS",
        "CELLS"
      ],
      "default": "ANY"
    },
    "ListSharedResourcesResponseSharedResource": {
      "type": "object",
      "properties": {
        "Node": {
          "$ref": "#/definitions/treeNode"
        },
        "Link": {
          "$ref": "#/definitions/restShareLink"
        },
        "Cells": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/restCell"
          }
        }
      },
      "title": "Container for ShareLink or Cell"
    },
    "NodeChangeEventEventType": {
      "type": "string",
      "enum": [
        "CREATE",
        "READ",
        "UPDATE_PATH",
        "UPDATE_CONTENT",
        "UPDATE_META",
        "UPDATE_USER_META",
        "DELETE"
      ],
      "default": "CREATE"
    },
    "PackagePackageStatus": {
      "type": "string",
      "enum": [
        "Draft",
        "Pending",
        "Released"
      ],
      "default": "Draft"
    },
    "ResourcePolicyQueryQueryType": {
      "type": "string",
      "enum": [
        "CONTEXT",
        "ANY",
        "NONE",
        "USER"
      ],
      "default": "CONTEXT"
    },
    "UpdateUserMetaNamespaceRequestUserMetaNsOp": {
      "type": "string",
      "enum": [
        "PUT",
        "DELETE"
      ],
      "default": "PUT"
    },
    "UpdateUserMetaRequestUserMetaOp": {
      "type": "string",
      "enum": [
        "PUT",
        "DELETE"
      ],
      "default": "PUT"
    },
    "activityObject": {
      "type": "object",
      "properties": {
        "jsonLdContext": {
          "type": "string"
        },
        "type": {
          "$ref": "#/definitions/activityObjectType"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "summary": {
          "type": "string"
        },
        "context": {
          "$ref": "#/definitions/activityObject"
        },
        "attachment": {
          "$ref": "#/definitions/activityObject"
        },
        "attributedTo": {
          "$ref": "#/definitions/activityObject"
        },
        "audience": {
          "$ref": "#/definitions/activityObject"
        },
        "content": {
          "$ref": "#/definitions/activityObject"
        },
        "startTime": {
          "type": "string",
          "format": "date-time"
        },
        "endTime": {
          "type": "string",
          "format": "date-time"
        },
        "published": {
          "type": "string",
          "format": "date-time"
        },
        "updated": {
          "type": "string",
          "format": "date-time"
        },
        "duration": {
          "type": "string",
          "format": "date-time"
        },
        "url": {
          "$ref": "#/definitions/activityObject"
        },
        "mediaType": {
          "type": "string"
        },
        "icon": {
          "$ref": "#/definitions/activityObject"
        },
        "image": {
          "$ref": "#/definitions/activityObject"
        },
        "preview": {
          "$ref": "#/definitions/activityObject"
        },
        "location": {
          "$ref": "#/definitions/activityObject"
        },
        "inReplyTo": {
          "$ref": "#/definitions/activityObject"
        },
        "replies": {
          "$ref": "#/definitions/activityObject"
        },
        "tag": {
          "$ref": "#/definitions/activityObject"
        },
        "generator": {
          "$ref": "#/definitions/activityObject"
        },
        "to": {
          "$ref": "#/definitions/activityObject"
        },
        "bto": {
          "$ref": "#/definitions/activityObject"
        },
        "cc": {
          "$ref": "#/definitions/activityObject"
        },
        "bcc": {
          "$ref": "#/definitions/activityObject"
        },
        "actor": {
          "$ref": "#/definitions/activityObject",
          "title": "Activity Properties"
        },
        "object": {
          "$ref": "#/definitions/activityObject"
        },
        "target": {
          "$ref": "#/definitions/activityObject"
        },
        "result": {
          "$ref": "#/definitions/activityObject"
        },
        "origin": {
          "$ref": "#/definitions/activityObject"
        },
        "instrument": {
          "$ref": "#/definitions/activityObject"
        },
        "href": {
          "type": "string",
          "title": "Link Properties"
        },
        "rel": {
          "type": "string"
        },
        "hreflang": {
          "type": "string"
        },
        "height": {
          "type": "integer",
          "format": "int32"
        },
        "width": {
          "type": "integer",
          "format": "int32"
        },
        "oneOf": {
          "$ref": "#/definitions/activityObject",
          "title": "Question Properties"
        },
        "anyOf": {
          "$ref": "#/definitions/activityObject"
        },
        "closed": {
          "type": "string",
          "format": "date-time"
        },
        "subject": {
          "$ref": "#/definitions/activityObject",
          "title": "Relationship Properties"
        },
        "relationship": {
          "$ref": "#/definitions/activityObject"
        },
        "formerType": {
          "$ref": "#/definitions/activityObjectType",
          "title": "Tombstone Properties"
        },
        "deleted": {
          "type": "string",
          "format": "date-time"
        },
        "accuracy": {
          "type": "number",
          "format": "float",
          "title": "Place Properties"
        },
        "altitude": {
          "type": "number",
          "format": "float"
        },
        "latitude": {
          "type": "number",
          "format": "float"
        },
        "longitude": {
          "type": "number",
          "format": "float"
        },
        "radius": {
          "type": "number",
          "format": "float"
        },
        "units": {
          "type": "string"
        },
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/activityObject"
          },
          "title": "Collection Properties"
        },
        "totalItems": {
          "type": "integer",
          "format": "int32"
        },
        "current": {
          "$ref": "#/definitions/activityObject"
        },
        "first": {
          "$ref": "#/definitions/activityObject"
        },
        "last": {
          "$ref": "#/definitions/activityObject"
        },
        "partOf": {
          "$ref": "#/definitions/activityObject"
        },
        "next": {
          "$ref": "#/definitions/activityObject"
        },
        "prev": {
          "$ref": "#/definitions/activityObject"
        }
      }
    },
    "activityObjectType": {
      "type": "string",
      "enum": [
        "BaseObject",
        "Activity",
        "Link",
        "Mention",
        "Collection",
        "OrderedCollection",
        "CollectionPage",
        "OrderedCollectionPage",
        "Application",
        "Group",
        "Organization",
        "Person",
        "Service",
        "Article",
        "Audio",
        "Document",
        "Event",
        "Image",
        "Note",
        "Page",
        "Place",
        "Profile",
        "Relationship",
        "Tombstone",
        "Video",
        "Accept",
        "Add",
        "Announce",
        "Arrive",
        "Block",
        "Create",
        "Delete",
        "Dislike",
        "Flag",
        "Follow",
        "Ignore",
        "Invite",
        "Join",
        "Leave",
        "Like",
        "Listen",
        "Move",
        "Offer",
        "Question",
        "Reject",
        "Read",
        "Remove",
        "TentativeReject",
        "TentativeAccept",
        "Travel",
        "Undo",
        "Update",
        "View",
        "Workspace",
        "Digest",
        "Folder",
        "Cell",
        "Share"
      ],
      "default": "BaseObject",
      "title": "- Collection: CollectionTypes\n - Application: Actor Types\n - Article: Objects Types\n - Accept: Activity Types\n - Workspace: Pydio Types"
    },
    "activityOwnerType": {
      "type": "string",
      "enum": [
        "NODE",
        "USER"
      ],
      "default": "NODE"
    },
    "activitySearchSubscriptionsRequest": {
      "type": "object",
      "properties": {
        "UserIds": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "ObjectTypes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/activityOwnerType"
          }
        },
        "ObjectIds": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "activityStreamActivitiesRequest": {
      "type": "object",
      "properties": {
        "Context": {
          "$ref": "#/definitions/activityStreamContext"
        },
        "ContextData": {
          "type": "string"
        },
        "StreamFilter": {
          "type": "string"
        },
        "BoxName": {
          "type": "string"
        },
        "UnreadCountOnly": {
          "type": "boolean",
          "format": "boolean"
        },
        "Offset": {
          "type": "string",
          "format": "int64"
        },
        "Limit": {
          "type": "string",
          "format": "int64"
        },
        "AsDigest": {
          "type": "boolean",
          "format": "boolean"
        },
        "PointOfView": {
          "$ref": "#/definitions/activitySummaryPointOfView"
        },
        "Language": {
          "type": "string"
        }
      }
    },
    "activityStreamContext": {
      "type": "string",
      "enum": [
        "MYFEED",
        "USER_ID",
        "NODE_ID"
      ],
      "default": "MYFEED"
    },
    "activitySubscription": {
      "type": "object",
      "properties": {
        "UserId": {
          "type": "string"
        },
        "ObjectType": {
          "$ref": "#/definitions/activityOwnerType"
        },
        "ObjectId": {
          "type": "string"
        },
        "Events": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "activitySummaryPointOfView": {
      "type": "string",
      "enum": [
        "GENERIC",
        "ACTOR",
        "SUBJECT"
      ],
      "default": "GENERIC"
    },
    "ctlPeer": {
      "type": "object",
      "properties": {
        "Id": {
          "type": "string"
        },
        "Address": {
          "type": "string"
        },
        "Port": {
          "type": "integer",
          "format": "int32"
        },
        "Metadata": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        }
      }
    },
    "ctlService": {
      "type": "object",
      "properties": {
        "Name": {
          "type": "string"
        },
        "Version": {
          "type": "string"
        },
        "Description": {
          "type": "string"
        },
        "Tag": {
          "type": "string"
        },
        "Controllable": {
          "type": "boolean",
          "format": "boolean"
        },
        "Status": {
          "$ref": "#/definitions/ctlServiceStatus"
        },
        "RunningPeers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ctlPeer"
          }
        }
      }
    },
    "ctlServiceCommand": {
      "type": "string",
      "enum": [
        "START",
        "STOP"
      ],
      "default": "START"
    },
    "ctlServiceStatus": {
      "type": "string",
      "enum": [
        "ANY",
        "STOPPED",
        "STARTING",
        "STOPPING",
        "STARTED"
      ],
      "default": "ANY"
    },
    "encryptionAdminCreateKeyRequest": {
      "type": "object",
      "properties": {
        "KeyID": {
          "type": "string"
        },
        "Label": {
          "type": "string"
        }
      }
    },
    "encryptionAdminCreateKeyResponse": {
      "type": "object",
      "properties": {
        "Success": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "encryptionAdminDeleteKeyRequest": {
      "type": "object",
      "properties": {
        "KeyID": {
          "type": "string"
        }
      }
    },
    "encryptionAdminDeleteKeyResponse": {
      "type": "object",
      "properties": {
        "Success": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "encryptionAdminExportKeyRequest": {
      "type": "object",
      "properties": {
        "KeyID": {
          "type": "string"
        },
        "StrPassword": {
          "type": "string"
        }
      }
    },
    "encryptionAdminExportKeyResponse": {
      "type": "object",
      "properties": {
        "Key": {
          "$ref": "#/definitions/encryptionKey"
        }
      }
    },
    "encryptionAdminImportKeyRequest": {
      "type": "object",
      "properties": {
        "Key": {
          "$ref": "#/definitions/encryptionKey"
        },
        "StrPassword": {
          "type": "string"
        },
        "Override": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "encryptionAdminImportKeyResponse": {
      "type": "object",
      "properties": {
        "Success": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "encryptionAdminListKeysRequest": {
      "type": "object"
    },
    "encryptionAdminListKeysResponse": {
      "type": "object",
      "properties": {
        "Keys": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/encryptionKey"
          }
        }
      }
    },
    "encryptionExport": {
      "type": "object",
      "properties": {
        "By": {
          "type": "string"
        },
        "Date": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "encryptionImport": {
      "type": "object",
      "properties": {
        "By": {
          "type": "string"
        },
        "Date": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "encryptionKey": {
      "type": "object",
      "properties": {
        "Owner": {
          "type": "string"
        },
        "ID": {
          "type": "string"
        },
        "Label": {
          "type": "string"
        },
        "Content": {
          "type": "string"
        },
        "CreationDate": {
          "type": "integer",
          "format": "int32"
        },
        "Info": {
          "$ref": "#/definitions/encryptionKeyInfo"
        }
      }
    },
    "encryptionKeyInfo": {
      "type": "object",
      "properties": {
        "Exports": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/encryptionExport"
          }
        },
        "Imports": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/encryptionImport"
          }
        }
      }
    },
    "idmACL": {
      "type": "object",
      "properties": {
        "ID": {
          "type": "string"
        },
        "Action": {
          "$ref": "#/definitions/idmACLAction"
        },
        "RoleID": {
          "type": "string"
        },
        "WorkspaceID": {
          "type": "string"
        },
        "NodeID": {
          "type": "string"
        }
      }
    },
    "idmACLAction": {
      "type": "object",
      "properties": {
        "Name": {
          "type": "string"
        },
        "Value": {
          "type": "string"
        }
      }
    },
    "idmACLSingleQuery": {
      "type": "object",
      "properties": {
        "Actions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idmACLAction"
          }
        },
        "RoleIDs": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "WorkspaceIDs": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "NodeIDs": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "not": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "idmListPolicyGroupsRequest": {
      "type": "object"
    },
    "idmListPolicyGroupsResponse": {
      "type": "object",
      "properties": {
        "PolicyGroups": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idmPolicyGroup"
          }
        },
        "Total": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "idmNodeType": {
      "type": "string",
      "enum": [
        "UNKNOWN",
        "USER",
        "GROUP"
      ],
      "default": "UNKNOWN"
    },
    "idmPolicy": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "subjects": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "resources": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "actions": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "effect": {
          "$ref": "#/definitions/idmPolicyEffect"
        },
        "conditions": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/idmPolicyCondition"
          }
        }
      }
    },
    "idmPolicyCondition": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string"
        },
        "jsonOptions": {
          "type": "string"
        }
      }
    },
    "idmPolicyEffect": {
      "type": "string",
      "enum": [
        "unknown",
        "deny",
        "allow"
      ],
      "default": "unknown"
    },
    "idmPolicyGroup": {
      "type": "object",
      "properties": {
        "Uuid": {
          "type": "string"
        },
        "Name": {
          "type": "string"
        },
        "Description": {
          "type": "string"
        },
        "OwnerUuid": {
          "type": "string"
        },
        "ResourceGroup": {
          "$ref": "#/definitions/idmPolicyResourceGroup"
        },
        "LastUpdated": {
          "type": "integer",
          "format": "int32"
        },
        "Policies": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idmPolicy"
          }
        }
      }
    },
    "idmPolicyResourceGroup": {
      "type": "string",
      "enum": [
        "rest",
        "acl",
        "oidc"
      ],
      "default": "rest"
    },
    "idmRole": {
      "type": "object",
      "properties": {
        "Uuid": {
          "type": "string"
        },
        "Label": {
          "type": "string"
        },
        "IsTeam": {
          "type": "boolean",
          "format": "boolean"
        },
        "GroupRole": {
          "type": "boolean",
          "format": "boolean"
        },
        "UserRole": {
          "type": "boolean",
          "format": "boolean"
        },
        "LastUpdated": {
          "type": "integer",
          "format": "int32"
        },
        "AutoApplies": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "Policies": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/serviceResourcePolicy"
          }
        },
        "PoliciesContextEditable": {
          "type": "boolean",
          "format": "boolean"
        },
        "ForceOverride": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "idmRoleSingleQuery": {
      "type": "object",
      "properties": {
        "Uuid": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "Label": {
          "type": "string"
        },
        "IsTeam": {
          "type": "boolean",
          "format": "boolean"
        },
        "IsGroupRole": {
          "type": "boolean",
          "format": "boolean"
        },
        "IsUserRole": {
          "type": "boolean",
          "format": "boolean"
        },
        "HasAutoApply": {
          "type": "boolean",
          "format": "boolean"
        },
        "not": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "idmSearchUserMetaRequest": {
      "type": "object",
      "properties": {
        "MetaUuids": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "NodeUuids": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "Namespace": {
          "type": "string"
        },
        "ResourceSubjectOwner": {
          "type": "string"
        },
        "ResourceQuery": {
          "$ref": "#/definitions/serviceResourcePolicyQuery"
        }
      },
      "title": "Request for searching UserMeta by NodeUuid or by Namespace"
    },
    "idmUpdateUserMetaNamespaceRequest": {
      "type": "object",
      "properties": {
        "Operation": {
          "$ref": "#/definitions/UpdateUserMetaNamespaceRequestUserMetaNsOp"
        },
        "Namespaces": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idmUserMetaNamespace"
          }
        }
      },
      "title": "Modify UserMetaNamespaces"
    },
    "idmUpdateUserMetaNamespaceResponse": {
      "type": "object",
      "properties": {
        "Namespaces": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idmUserMetaNamespace"
          }
        }
      },
      "title": "Response of the"
    },
    "idmUpdateUserMetaRequest": {
      "type": "object",
      "properties": {
        "Operation": {
          "$ref": "#/definitions/UpdateUserMetaRequestUserMetaOp"
        },
        "MetaDatas": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idmUserMeta"
          }
        }
      },
      "title": "Request for modifying UserMeta"
    },
    "idmUpdateUserMetaResponse": {
      "type": "object",
      "properties": {
        "MetaDatas": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idmUserMeta"
          }
        }
      },
      "title": "Response of UpdateUserMeta service"
    },
    "idmUser": {
      "type": "object",
      "properties": {
        "Uuid": {
          "type": "string"
        },
        "GroupPath": {
          "type": "string"
        },
        "Attributes": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        },
        "Roles": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idmRole"
          }
        },
        "Login": {
          "type": "string",
          "title": "User specific data"
        },
        "Password": {
          "type": "string"
        },
        "OldPassword": {
          "type": "string"
        },
        "IsGroup": {
          "type": "boolean",
          "format": "boolean",
          "title": "Group specific data"
        },
        "GroupLabel": {
          "type": "string"
        },
        "Policies": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/serviceResourcePolicy"
          }
        },
        "PoliciesContextEditable": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "idmUserMeta": {
      "type": "object",
      "properties": {
        "Uuid": {
          "type": "string"
        },
        "NodeUuid": {
          "type": "string"
        },
        "Namespace": {
          "type": "string"
        },
        "JsonValue": {
          "type": "string"
        },
        "Policies": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/serviceResourcePolicy"
          }
        },
        "PoliciesContextEditable": {
          "type": "boolean",
          "format": "boolean"
        }
      },
      "title": "Piece of metadata attached to a node"
    },
    "idmUserMetaNamespace": {
      "type": "object",
      "properties": {
        "Namespace": {
          "type": "string"
        },
        "Label": {
          "type": "string"
        },
        "Order": {
          "type": "integer",
          "format": "int32"
        },
        "Indexable": {
          "type": "boolean",
          "format": "boolean"
        },
        "JsonDefinition": {
          "type": "string"
        },
        "Policies": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/serviceResourcePolicy"
          }
        }
      },
      "title": "Globally declared Namespace with associated policies"
    },
    "idmUserSingleQuery": {
      "type": "object",
      "properties": {
        "Uuid": {
          "type": "string"
        },
        "Login": {
          "type": "string"
        },
        "Password": {
          "type": "string"
        },
        "GroupPath": {
          "type": "string",
          "title": "Search on group path, and if so, search recursively"
        },
        "Recursive": {
          "type": "boolean",
          "format": "boolean"
        },
        "FullPath": {
          "type": "string",
          "title": "Search a specific group by path"
        },
        "AttributeName": {
          "type": "string",
          "title": "Search on attribute"
        },
        "AttributeValue": {
          "type": "string"
        },
        "AttributeAnyValue": {
          "type": "boolean",
          "format": "boolean"
        },
        "HasRole": {
          "type": "string",
          "title": "Search on roles"
        },
        "NodeType": {
          "$ref": "#/definitions/idmNodeType"
        },
        "not": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "idmWorkspace": {
      "type": "object",
      "properties": {
        "UUID": {
          "type": "string"
        },
        "Label": {
          "type": "string"
        },
        "Description": {
          "type": "string"
        },
        "Slug": {
          "type": "string"
        },
        "Scope": {
          "$ref": "#/definitions/idmWorkspaceScope"
        },
        "LastUpdated": {
          "type": "integer",
          "format": "int32"
        },
        "Policies": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/serviceResourcePolicy"
          }
        },
        "Attributes": {
          "type": "string"
        },
        "RootUUIDs": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "RootNodes": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/treeNode"
          }
        },
        "PoliciesContextEditable": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "idmWorkspaceScope": {
      "type": "string",
      "enum": [
        "ANY",
        "ADMIN",
        "ROOM",
        "LINK"
      ],
      "default": "ANY"
    },
    "idmWorkspaceSingleQuery": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string"
        },
        "label": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "slug": {
          "type": "string"
        },
        "scope": {
          "$ref": "#/definitions/idmWorkspaceScope"
        },
        "not": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "installCheckResult": {
      "type": "object",
      "properties": {
        "Name": {
          "type": "string"
        },
        "Success": {
          "type": "boolean",
          "format": "boolean"
        },
        "JsonResult": {
          "type": "string"
        }
      }
    },
    "installGetAgreementResponse": {
      "type": "object",
      "properties": {
        "Text": {
          "type": "string"
        }
      }
    },
    "installGetDefaultsResponse": {
      "type": "object",
      "properties": {
        "config": {
          "$ref": "#/definitions/installInstallConfig"
        }
      }
    },
    "installInstallConfig": {
      "type": "object",
      "properties": {
        "internalUrl": {
          "type": "string"
        },
        "dbConnectionType": {
          "type": "string"
        },
        "dbTCPHostname": {
          "type": "string"
        },
        "dbTCPPort": {
          "type": "string"
        },
        "dbTCPName": {
          "type": "string"
        },
        "dbTCPUser": {
          "type": "string"
        },
        "dbTCPPassword": {
          "type": "string"
        },
        "dbSocketFile": {
          "type": "string"
        },
        "dbSocketName": {
          "type": "string"
        },
        "dbSocketUser": {
          "type": "string"
        },
        "dbSocketPassword": {
          "type": "string"
        },
        "dbManualDSN": {
          "type": "string"
        },
        "dsName": {
          "type": "string"
        },
        "dsPort": {
          "type": "string"
        },
        "dsFolder": {
          "type": "string"
        },
        "externalMicro": {
          "type": "string"
        },
        "externalGateway": {
          "type": "string"
        },
        "externalWebsocket": {
          "type": "string"
        },
        "externalFrontPlugins": {
          "type": "string"
        },
        "externalDAV": {
          "type": "string"
        },
        "externalWOPI": {
          "type": "string"
        },
        "externalDex": {
          "type": "string"
        },
        "externalDexID": {
          "type": "string"
        },
        "externalDexSecret": {
          "type": "string"
        },
        "frontendHosts": {
          "type": "string"
        },
        "frontendLogin": {
          "type": "string"
        },
        "frontendPassword": {
          "type": "string"
        },
        "frontendRepeatPassword": {
          "type": "string"
        },
        "frontendApplicationTitle": {
          "type": "string"
        },
        "frontendDefaultLanguage": {
          "type": "string"
        },
        "licenseRequired": {
          "type": "boolean",
          "format": "boolean"
        },
        "licenseString": {
          "type": "string"
        },
        "CheckResults": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/installCheckResult"
          }
        }
      }
    },
    "installInstallRequest": {
      "type": "object",
      "properties": {
        "config": {
          "$ref": "#/definitions/installInstallConfig"
        }
      }
    },
    "installInstallResponse": {
      "type": "object",
      "properties": {
        "success": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "installPerformCheckRequest": {
      "type": "object",
      "properties": {
        "Name": {
          "type": "string"
        },
        "Config": {
          "$ref": "#/definitions/installInstallConfig"
        }
      }
    },
    "installPerformCheckResponse": {
      "type": "object",
      "properties": {
        "Result": {
          "$ref": "#/definitions/installCheckResult"
        }
      }
    },
    "jobsAction": {
      "type": "object",
      "properties": {
        "ID": {
          "type": "string",
          "title": "String Identifier for specific action"
        },
        "NodesSelector": {
          "$ref": "#/definitions/jobsNodesSelector",
          "title": "Nodes Selector"
        },
        "UsersSelector": {
          "$ref": "#/definitions/jobsUsersSelector",
          "title": "Users Selector"
        },
        "NodesFilter": {
          "$ref": "#/definitions/jobsNodesSelector",
          "title": "Node Filter"
        },
        "UsersFilter": {
          "$ref": "#/definitions/jobsUsersSelector",
          "title": "User Filter"
        },
        "SourceFilter": {
          "$ref": "#/definitions/jobsSourceFilter",
          "title": "Source filter"
        },
        "Parameters": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "Defined parameters for this action"
        },
        "ChainedActions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/jobsAction"
          },
          "title": "Other actions to perform after this one is finished,\nusing the Output of this action as Input for the next.\nIf there are many, it is considered they can be triggered\nin parallel"
        }
      }
    },
    "jobsActionLog": {
      "type": "object",
      "properties": {
        "Action": {
          "$ref": "#/definitions/jobsAction"
        },
        "InputMessage": {
          "$ref": "#/definitions/jobsActionMessage"
        },
        "OutputMessage": {
          "$ref": "#/definitions/jobsActionMessage"
        }
      }
    },
    "jobsActionMessage": {
      "type": "object",
      "properties": {
        "Event": {
          "$ref": "#/definitions/protobufAny"
        },
        "Nodes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/treeNode"
          }
        },
        "Users": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idmUser"
          }
        },
        "Activities": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/activityObject"
          }
        },
        "OutputChain": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/jobsActionOutput"
          }
        }
      },
      "title": "Message passed along from one action to another, main properties\nare modified by the various actions.\nOutputChain is being stacked up when passing through actions"
    },
    "jobsActionOutput": {
      "type": "object",
      "properties": {
        "Success": {
          "type": "boolean",
          "format": "boolean"
        },
        "RawBody": {
          "type": "string",
          "format": "byte"
        },
        "StringBody": {
          "type": "string"
        },
        "JsonBody": {
          "type": "string",
          "format": "byte"
        },
        "ErrorString": {
          "type": "string"
        },
        "Ignored": {
          "type": "boolean",
          "format": "boolean"
        },
        "Time": {
          "type": "integer",
          "format": "int32"
        }
      },
      "title": "Standard output of an action. Success value is required\nother are optional"
    },
    "jobsCommand": {
      "type": "string",
      "enum": [
        "None",
        "Pause",
        "Resume",
        "Stop",
        "Delete",
        "RunOnce",
        "Inactive",
        "Active"
      ],
      "default": "None"
    },
    "jobsCtrlCommand": {
      "type": "object",
      "properties": {
        "Cmd": {
          "$ref": "#/definitions/jobsCommand"
        },
        "JobId": {
          "type": "string"
        },
        "TaskId": {
          "type": "string"
        },
        "OwnerId": {
          "type": "string"
        }
      }
    },
    "jobsCtrlCommandResponse": {
      "type": "object",
      "properties": {
        "Msg": {
          "type": "string"
        }
      }
    },
    "jobsDeleteTasksRequest": {
      "type": "object",
      "properties": {
        "JobId": {
          "type": "string",
          "title": "Id of the job"
        },
        "TaskID": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "Ids of tasks to delete"
        },
        "Status": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/jobsTaskStatus"
          },
          "title": "If no TaskID and/or no JobID are passed, delete tasks by status"
        },
        "PruneLimit": {
          "type": "integer",
          "format": "int32",
          "title": "If deleting by status, optionally keep only a number of tasks"
        }
      }
    },
    "jobsDeleteTasksResponse": {
      "type": "object",
      "properties": {
        "Deleted": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "jobsJob": {
      "type": "object",
      "properties": {
        "ID": {
          "type": "string",
          "title": "Unique ID for this Job"
        },
        "Label": {
          "type": "string",
          "title": "Human-readable Label"
        },
        "Owner": {
          "type": "string",
          "title": "Who created this Job"
        },
        "Inactive": {
          "type": "boolean",
          "format": "boolean",
          "title": "Admin can temporarily disable this job"
        },
        "Languages": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "Optional list of languages detected in the context at launch time"
        },
        "EventNames": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "How the job will be triggered.\nOne of these must be set (not exclusive)\nListen to a given set of events"
        },
        "Schedule": {
          "$ref": "#/definitions/jobsSchedule",
          "title": "Schedule a periodic repetition"
        },
        "AutoStart": {
          "type": "boolean",
          "format": "boolean",
          "title": "Start task as soon as job is inserted"
        },
        "AutoClean": {
          "type": "boolean",
          "format": "boolean",
          "title": "Remove job automatically once it is finished (success only)"
        },
        "Actions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/jobsAction"
          },
          "title": "Chain of actions to perform"
        },
        "MaxConcurrency": {
          "type": "integer",
          "format": "int32",
          "title": "Task properties"
        },
        "TasksSilentUpdate": {
          "type": "boolean",
          "format": "boolean",
          "title": "Do not send notification on task update"
        },
        "Tasks": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/jobsTask"
          },
          "title": "Filled with currently running tasks"
        },
        "NodeEventFilter": {
          "$ref": "#/definitions/jobsNodesSelector"
        },
        "UserEventFilter": {
          "$ref": "#/definitions/jobsUsersSelector"
        }
      }
    },
    "jobsListJobsRequest": {
      "type": "object",
      "properties": {
        "Owner": {
          "type": "string"
        },
        "EventsOnly": {
          "type": "boolean",
          "format": "boolean"
        },
        "TimersOnly": {
          "type": "boolean",
          "format": "boolean"
        },
        "LoadTasks": {
          "$ref": "#/definitions/jobsTaskStatus"
        },
        "JobIDs": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "TasksOffset": {
          "type": "integer",
          "format": "int32"
        },
        "TasksLimit": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "jobsNodesSelector": {
      "type": "object",
      "properties": {
        "All": {
          "type": "boolean",
          "format": "boolean",
          "title": "Select all files - ignore any other condition"
        },
        "Pathes": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "Preset list of node pathes"
        },
        "Nodes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/treeNode"
          },
          "title": "Preset set of nodes"
        },
        "Query": {
          "$ref": "#/definitions/serviceQuery",
          "title": "Query to apply to select users (or filter a given node passed by event)"
        },
        "Collect": {
          "type": "boolean",
          "format": "boolean",
          "title": "Wether to trigger one action per node or one action\nwith all nodes as selection"
        }
      },
      "title": "/////////////////\nJOB  SERVICE  //\n/////////////////"
    },
    "jobsSchedule": {
      "type": "object",
      "properties": {
        "Iso8601Schedule": {
          "type": "string",
          "description": "ISO 8601 Description of the scheduling for instance \"R2/2015-06-04T19:25:16.828696-07:00/PT4S\"\nwhere first part is the number of repetitions (if 0, infinite repetition), \nsecond part the starting date and last part, the duration between 2 occurrences."
        },
        "Iso8601MinDelta": {
          "type": "string",
          "title": "Minimum time between two runs"
        }
      }
    },
    "jobsSourceFilter": {
      "type": "object",
      "properties": {
        "Query": {
          "$ref": "#/definitions/serviceQuery",
          "title": "Can be built with SourceSingleQuery or ActionOutputQuery"
        }
      }
    },
    "jobsTask": {
      "type": "object",
      "properties": {
        "ID": {
          "type": "string"
        },
        "JobID": {
          "type": "string"
        },
        "Status": {
          "$ref": "#/definitions/jobsTaskStatus"
        },
        "StatusMessage": {
          "type": "string"
        },
        "TriggerOwner": {
          "type": "string"
        },
        "StartTime": {
          "type": "integer",
          "format": "int32"
        },
        "EndTime": {
          "type": "integer",
          "format": "int32"
        },
        "CanStop": {
          "type": "boolean",
          "format": "boolean",
          "title": "Can be interrupted"
        },
        "CanPause": {
          "type": "boolean",
          "format": "boolean",
          "title": "Can be paused/resumed"
        },
        "HasProgress": {
          "type": "boolean",
          "format": "boolean",
          "title": "Tasks publish a progress"
        },
        "Progress": {
          "type": "number",
          "format": "float",
          "title": "Float value of the progress between 0 and 1"
        },
        "ActionsLogs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/jobsActionLog"
          },
          "title": "Logs of all the actions performed"
        }
      }
    },
    "jobsTaskStatus": {
      "type": "string",
      "enum": [
        "Unknown",
        "Idle",
        "Running",
        "Finished",
        "Interrupted",
        "Paused",
        "Any",
        "Error",
        "Queued"
      ],
      "default": "Unknown",
      "title": "/////////////////\nTASK SERVICE  //\n/////////////////"
    },
    "jobsUsersSelector": {
      "type": "object",
      "properties": {
        "All": {
          "type": "boolean",
          "format": "boolean",
          "title": "Select all users"
        },
        "Users": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idmUser"
          },
          "title": "Preset set of Users"
        },
        "Query": {
          "$ref": "#/definitions/serviceQuery",
          "title": "Filter users using this query"
        },
        "Collect": {
          "type": "boolean",
          "format": "boolean",
          "title": "Wether to trigger one action per user or one action\nwith all user as a selection"
        }
      }
    },
    "logListLogRequest": {
      "type": "object",
      "properties": {
        "Query": {
          "type": "string",
          "title": "Bleve-type Query stsring"
        },
        "Page": {
          "type": "integer",
          "format": "int32",
          "title": "Start at page"
        },
        "Size": {
          "type": "integer",
          "format": "int32",
          "title": "Number of results"
        },
        "Format": {
          "$ref": "#/definitions/ListLogRequestLogFormat"
        }
      },
      "description": "ListLogRequest launches a parameterised query in the log repository and streams the results."
    },
    "logLogMessage": {
      "type": "object",
      "properties": {
        "Ts": {
          "type": "integer",
          "format": "int32",
          "title": "Generic zap fields"
        },
        "Level": {
          "type": "string"
        },
        "Logger": {
          "type": "string"
        },
        "Msg": {
          "type": "string"
        },
        "MsgId": {
          "type": "string",
          "title": "Pydio specific"
        },
        "UserName": {
          "type": "string",
          "title": "User Info"
        },
        "UserUuid": {
          "type": "string"
        },
        "GroupPath": {
          "type": "string"
        },
        "Profile": {
          "type": "string"
        },
        "RoleUuids": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "RemoteAddress": {
          "type": "string",
          "title": "Client info"
        },
        "UserAgent": {
          "type": "string"
        },
        "HttpProtocol": {
          "type": "string"
        },
        "NodeUuid": {
          "type": "string",
          "title": "Tree Info"
        },
        "NodePath": {
          "type": "string"
        },
        "WsUuid": {
          "type": "string"
        },
        "WsScope": {
          "type": "string"
        },
        "SpanUuid": {
          "type": "string",
          "title": "Span Info"
        },
        "SpanParentUuid": {
          "type": "string"
        },
        "SpanRootUuid": {
          "type": "string"
        },
        "OperationUuid": {
          "type": "string",
          "title": "High Level Operation Info"
        },
        "OperationLabel": {
          "type": "string"
        }
      },
      "description": "LogMessage is the format used to transmit log messages to clients via the REST API."
    },
    "mailerMail": {
      "type": "object",
      "properties": {
        "From": {
          "$ref": "#/definitions/mailerUser"
        },
        "To": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mailerUser"
          }
        },
        "Cc": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mailerUser"
          }
        },
        "DateSent": {
          "type": "string",
          "format": "int64"
        },
        "Subject": {
          "type": "string"
        },
        "ContentPlain": {
          "type": "string"
        },
        "ContentHtml": {
          "type": "string"
        },
        "ContentMarkdown": {
          "type": "string"
        },
        "Attachments": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "ThreadUuid": {
          "type": "string",
          "title": "Could be used for Re: ... conversations"
        },
        "ThreadIndex": {
          "type": "string"
        },
        "TemplateId": {
          "type": "string"
        },
        "TemplateData": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        },
        "Retries": {
          "type": "integer",
          "format": "int32"
        },
        "sendErrors": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "Sender": {
          "$ref": "#/definitions/mailerUser"
        }
      }
    },
    "mailerSendMailResponse": {
      "type": "object",
      "properties": {
        "Success": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "mailerUser": {
      "type": "object",
      "properties": {
        "Uuid": {
          "type": "string"
        },
        "Address": {
          "type": "string"
        },
        "Name": {
          "type": "string"
        },
        "Language": {
          "type": "string"
        }
      }
    },
    "objectDataSource": {
      "type": "object",
      "properties": {
        "Name": {
          "type": "string"
        },
        "Disabled": {
          "type": "boolean",
          "format": "boolean"
        },
        "StorageType": {
          "$ref": "#/definitions/objectStorageType"
        },
        "StorageConfiguration": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        },
        "ObjectsServiceName": {
          "type": "string"
        },
        "ObjectsHost": {
          "type": "string"
        },
        "ObjectsPort": {
          "type": "integer",
          "format": "int32"
        },
        "ObjectsSecure": {
          "type": "boolean",
          "format": "boolean"
        },
        "ObjectsBucket": {
          "type": "string"
        },
        "ObjectsBaseFolder": {
          "type": "string"
        },
        "ApiKey": {
          "type": "string"
        },
        "ApiSecret": {
          "type": "string"
        },
        "PeerAddress": {
          "type": "string"
        },
        "Watch": {
          "type": "boolean",
          "format": "boolean"
        },
        "EncryptionMode": {
          "$ref": "#/definitions/objectEncryptionMode"
        },
        "EncryptionKey": {
          "type": "string"
        },
        "VersioningPolicyName": {
          "type": "string"
        },
        "CreationDate": {
          "type": "integer",
          "format": "int32"
        },
        "LastSynchronizationDate": {
          "type": "integer",
          "format": "int32"
        }
      },
      "title": "DataSource Object description"
    },
    "objectEncryptionMode": {
      "type": "string",
      "enum": [
        "CLEAR",
        "MASTER",
        "USER",
        "USER_PWD"
      ],
      "default": "CLEAR",
      "title": "Type of Encryption"
    },
    "objectStorageType": {
      "type": "string",
      "enum": [
        "LOCAL",
        "S3",
        "SMB",
        "CELLS",
        "AZURE",
        "GCS",
        "B2",
        "MANTA",
        "SIA"
      ],
      "default": "LOCAL",
      "title": "Type of Gateway"
    },
    "protobufAny": {
      "type": "object",
      "properties": {
        "type_url": {
          "type": "string",
          "description": "A URL/resource name whose content describes the type of the\nserialized protocol buffer message.\n\nFor URLs which use the scheme http, https, or no scheme, the\nfollowing restrictions and interpretations apply:\n\n* If no scheme is provided, https is assumed.\n* The last segment of the URL's path must represent the fully\n  qualified name of the type (as in path/google.protobuf.Duration).\n  The name should be in a canonical form (e.g., leading \".\" is\n  not accepted).\n* An HTTP GET on the URL must yield a [google.protobuf.Type][]\n  value in binary format, or produce an error.\n* Applications are allowed to cache lookup results based on the\n  URL, or have them precompiled into a binary to avoid any\n  lookup. Therefore, binary compatibility needs to be preserved\n  on changes to types. (Use versioned type names to manage\n  breaking changes.)\n\nSchemes other than http, https (or the empty scheme) might be\nused with implementation specific semantics."
        },
        "value": {
          "type": "string",
          "format": "byte",
          "description": "Must be a valid serialized protocol buffer of the above specified type."
        }
      },
      "description": "Any contains an arbitrary serialized protocol buffer message along with a\nURL that describes the type of the serialized message.\n\nProtobuf library provides support to pack/unpack Any values in the form\nof utility functions or additional generated methods of the Any type.\n\nExample 1: Pack and unpack a message in C++.\n\n    Foo foo = ...;\n    Any any;\n    any.PackFrom(foo);\n    ...\n    if (any.UnpackTo(\u0026foo)) {\n      ...\n    }\n\nExample 2: Pack and unpack a message in Java.\n\n    Foo foo = ...;\n    Any any = Any.pack(foo);\n    ...\n    if (any.is(Foo.class)) {\n      foo = any.unpack(Foo.class);\n    }\n\n Example 3: Pack and unpack a message in Python.\n\n    foo = Foo(...)\n    any = Any()\n    any.Pack(foo)\n    ...\n    if any.Is(Foo.DESCRIPTOR):\n      any.Unpack(foo)\n      ...\n\n Example 4: Pack and unpack a message in Go\n\n     foo := \u0026pb.Foo{...}\n     any, err := ptypes.MarshalAny(foo)\n     ...\n     foo := \u0026pb.Foo{}\n     if err := ptypes.UnmarshalAny(any, foo); err != nil {\n       ...\n     }\n\nThe pack methods provided by protobuf library will by default use\n'type.googleapis.com/full.type.name' as the type URL and the unpack\nmethods only use the fully qualified type name after the last '/'\nin the type URL, for example \"foo.bar.com/x/y.z\" will yield type\nname \"y.z\".\n\n\nJSON\n====\nThe JSON representation of an Any value uses the regular\nrepresentation of the deserialized, embedded message, with an\nadditional field @type which contains the type URL. Example:\n\n    package google.profile;\n    message Person {\n      string first_name = 1;\n      string last_name = 2;\n    }\n\n    {\n      \"@type\": \"type.googleapis.com/google.profile.Person\",\n      \"firstName\": \u003cstring\u003e,\n      \"lastName\": \u003cstring\u003e\n    }\n\nIf the embedded message type is well-known and has a custom JSON\nrepresentation, that representation will be embedded adding a field\nvalue which holds the custom JSON in addition to the @type\nfield. Example (for message [google.protobuf.Duration][]):\n\n    {\n      \"@type\": \"type.googleapis.com/google.protobuf.Duration\",\n      \"value\": \"1.212s\"\n    }"
    },
    "restACLCollection": {
      "type": "object",
      "properties": {
        "ACLs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idmACL"
          }
        },
        "Total": {
          "type": "integer",
          "format": "int32"
        }
      },
      "title": "Response for search request"
    },
    "restBackgroundJobResult": {
      "type": "object",
      "properties": {
        "Uuid": {
          "type": "string"
        },
        "Label": {
          "type": "string"
        },
        "NodeUuid": {
          "type": "string"
        }
      }
    },
    "restBulkMetaResponse": {
      "type": "object",
      "properties": {
        "Nodes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/treeNode"
          }
        },
        "Pagination": {
          "$ref": "#/definitions/restPagination"
        }
      }
    },
    "restCell": {
      "type": "object",
      "properties": {
        "Uuid": {
          "type": "string"
        },
        "Label": {
          "type": "string"
        },
        "Description": {
          "type": "string"
        },
        "RootNodes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/treeNode"
          }
        },
        "ACLs": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/restCellAcl"
          }
        },
        "Policies": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/serviceResourcePolicy"
          }
        },
        "PoliciesContextEditable": {
          "type": "boolean",
          "format": "boolean"
        }
      },
      "title": "Model for representing a shared room"
    },
    "restCellAcl": {
      "type": "object",
      "properties": {
        "RoleId": {
          "type": "string"
        },
        "Actions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idmACLAction"
          }
        },
        "IsUserRole": {
          "type": "boolean",
          "format": "boolean"
        },
        "User": {
          "$ref": "#/definitions/idmUser"
        },
        "Group": {
          "$ref": "#/definitions/idmUser"
        },
        "Role": {
          "$ref": "#/definitions/idmRole"
        }
      },
      "title": "Group collected acls by subjects"
    },
    "restChangeCollection": {
      "type": "object",
      "properties": {
        "Changes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/treeSyncChange"
          }
        },
        "LastSeqId": {
          "type": "string",
          "format": "int64"
        }
      }
    },
    "restChangeRequest": {
      "type": "object",
      "properties": {
        "SeqID": {
          "type": "string",
          "format": "int64"
        },
        "filter": {
          "type": "string"
        },
        "flatten": {
          "type": "boolean",
          "format": "boolean"
        },
        "stream": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "restConfiguration": {
      "type": "object",
      "properties": {
        "FullPath": {
          "type": "string"
        },
        "Data": {
          "type": "string"
        }
      },
      "title": "Configuration message. Data is an Json representation of any value"
    },
    "restControlServiceRequest": {
      "type": "object",
      "properties": {
        "ServiceName": {
          "type": "string"
        },
        "NodeName": {
          "type": "string"
        },
        "Command": {
          "$ref": "#/definitions/ctlServiceCommand"
        }
      }
    },
    "restCreateNodesRequest": {
      "type": "object",
      "properties": {
        "Nodes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/treeNode"
          }
        },
        "Recursive": {
          "type": "boolean",
          "format": "boolean"
        },
        "TemplateUUID": {
          "type": "string"
        }
      }
    },
    "restCreateSelectionRequest": {
      "type": "object",
      "properties": {
        "Nodes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/treeNode"
          }
        },
        "TargetAction": {
          "type": "string"
        },
        "Persist": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "restCreateSelectionResponse": {
      "type": "object",
      "properties": {
        "SelectionUUID": {
          "type": "string"
        },
        "Nodes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/treeNode"
          }
        }
      }
    },
    "restDataSourceCollection": {
      "type": "object",
      "properties": {
        "DataSources": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/objectDataSource"
          }
        },
        "Total": {
          "type": "integer",
          "format": "int32"
        }
      },
      "title": "Collection of datasources"
    },
    "restDeleteCellResponse": {
      "type": "object",
      "properties": {
        "Success": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "restDeleteDataSourceResponse": {
      "type": "object",
      "properties": {
        "Success": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "restDeleteNodesRequest": {
      "type": "object",
      "properties": {
        "Nodes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/treeNode"
          }
        },
        "Recursive": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "restDeleteNodesResponse": {
      "type": "object",
      "properties": {
        "DeleteJobs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/restBackgroundJobResult"
          }
        }
      }
    },
    "restError": {
      "type": "object",
      "properties": {
        "Code": {
          "type": "string",
          "title": "Unique ID of the error"
        },
        "Title": {
          "type": "string",
          "title": "Human-readable, short label"
        },
        "Detail": {
          "type": "string",
          "title": "Human-readable, longer description"
        },
        "Source": {
          "type": "string",
          "title": "Cells service name or other quickly useful info"
        },
        "Meta": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "Additional Metadata"
        }
      },
      "title": "Generic error message"
    },
    "restDeleteResponse": {
      "type": "object",
      "properties": {
        "Success": {
          "type": "boolean",
          "format": "boolean",
          "title": "If delete was successfull"
        },
        "NumRows": {
          "type": "string",
          "format": "int64",
          "title": "Number of records deleted during operation"
        }
      },
      "title": "Generic Message"
    },
    "restDeleteShareLinkResponse": {
      "type": "object",
      "properties": {
        "Success": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "restDeleteUserMetaTagsResponse": {
      "type": "object",
      "properties": {
        "Success": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "restDiscoveryResponse": {
      "type": "object",
      "properties": {
        "PackageType": {
          "type": "string"
        },
        "PackageLabel": {
          "type": "string"
        },
        "Version": {
          "type": "string"
        },
        "BuildStamp": {
          "type": "integer",
          "format": "int32"
        },
        "BuildRevision": {
          "type": "string"
        },
        "Endpoints": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        }
      }
    },
    "restFrontBinaryRequest": {
      "type": "object",
      "properties": {
        "BinaryType": {
          "type": "string",
          "title": "Currently supported values are USER and GLOBAL"
        },
        "Uuid": {
          "type": "string",
          "title": "Id of the binary"
        }
      },
      "title": "Donwload binary"
    },
    "restFrontBinaryResponse": {
      "type": "object",
      "title": "Not used, endpoint returns octet-stream"
    },
    "restFrontBootConfResponse": {
      "type": "object",
      "properties": {
        "JsonData": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        }
      }
    },
    "restFrontEnrollAuthRequest": {
      "type": "object",
      "properties": {
        "EnrollType": {
          "type": "string"
        },
        "EnrollInfo": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        }
      }
    },
    "restFrontEnrollAuthResponse": {
      "type": "object",
      "properties": {
        "Info": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "Any parameters can be returned"
        }
      }
    },
    "restFrontMessagesResponse": {
      "type": "object",
      "properties": {
        "Messages": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        }
      }
    },
    "restFrontPluginsResponse": {
      "type": "object"
    },
    "restFrontSessionRequest": {
      "type": "object",
      "properties": {
        "ClientTime": {
          "type": "integer",
          "format": "int32",
          "title": "Time reference for computing jwt expiry"
        },
        "AuthInfo": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "Data sent back by specific auth steps"
        },
        "Logout": {
          "type": "boolean",
          "format": "boolean",
          "title": "Kill session now"
        }
      }
    },
    "restFrontSessionResponse": {
      "type": "object",
      "properties": {
        "JWT": {
          "type": "string"
        },
        "ExpireTime": {
          "type": "integer",
          "format": "int32"
        },
        "Trigger": {
          "type": "string",
          "title": "Trigger a specific Auth step"
        },
        "TriggerInfo": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "Additional data for the trigger"
        }
      }
    },
    "restFrontStateResponse": {
      "type": "object"
    },
    "restGetBulkMetaRequest": {
      "type": "object",
      "properties": {
        "NodePaths": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "NodeUuids": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "AllMetaProviders": {
          "type": "boolean",
          "format": "boolean"
        },
        "Versions": {
          "type": "boolean",
          "format": "boolean"
        },
        "Offset": {
          "type": "integer",
          "format": "int32"
        },
        "Limit": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "restHeadNodeResponse": {
      "type": "object",
      "properties": {
        "Node": {
          "$ref": "#/definitions/treeNode"
        }
      }
    },
    "restListPeerFoldersRequest": {
      "type": "object",
      "properties": {
        "PeerAddress": {
          "type": "string"
        },
        "Path": {
          "type": "string"
        }
      }
    },
    "restListPeersAddressesResponse": {
      "type": "object",
      "properties": {
        "PeerAddresses": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "restListProcessesRequest": {
      "type": "object",
      "properties": {
        "PeerId": {
          "type": "string"
        },
        "ServiceName": {
          "type": "string"
        }
      }
    },
    "restListProcessesResponse": {
      "type": "object",
      "properties": {
        "Processes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/restProcess"
          }
        }
      }
    },
    "restListSharedResourcesRequest": {
      "type": "object",
      "properties": {
        "ShareType": {
          "$ref": "#/definitions/ListSharedResourcesRequestListShareType",
          "title": "Filter output to a given type"
        },
        "Subject": {
          "type": "string",
          "title": "Will restrict the list to the shares readable by a specific subject.\nIn user-context, current user is used by default. In admin-context, this can\nbe any resource policy subject"
        },
        "OwnedBySubject": {
          "type": "boolean",
          "format": "boolean",
          "title": "If true, will also check filter the output to shares actually owned by subject"
        },
        "Offset": {
          "type": "integer",
          "format": "int32"
        },
        "Limit": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "restListSharedResourcesResponse": {
      "type": "object",
      "properties": {
        "Resources": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ListSharedResourcesResponseSharedResource"
          },
          "title": "Actual results"
        },
        "Offset": {
          "type": "integer",
          "format": "int32",
          "title": "Cursor informations"
        },
        "Limit": {
          "type": "integer",
          "format": "int32"
        },
        "Total": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "restListTemplatesResponse": {
      "type": "object",
      "properties": {
        "Templates": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/restTemplate"
          }
        }
      }
    },
    "restListUserMetaTagsResponse": {
      "type": "object",
      "properties": {
        "Tags": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "restLogMessageCollection": {
      "type": "object",
      "properties": {
        "Logs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/logLogMessage"
          }
        }
      },
      "title": "Collection of serialized log messages"
    },
    "restMetaCollection": {
      "type": "object",
      "properties": {
        "NodePath": {
          "type": "string"
        },
        "Metadatas": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/restMetadata"
          }
        }
      }
    },
    "restMetaNamespaceRequest": {
      "type": "object",
      "properties": {
        "NodePath": {
          "type": "string"
        },
        "Namespace": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "restMetadata": {
      "type": "object",
      "properties": {
        "Namespace": {
          "type": "string"
        },
        "JsonMeta": {
          "type": "string"
        }
      }
    },
    "restNodesCollection": {
      "type": "object",
      "properties": {
        "Parent": {
          "$ref": "#/definitions/treeNode"
        },
        "Children": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/treeNode"
          }
        }
      }
    },
    "restOpenApiResponse": {
      "type": "object"
    },
    "restPagination": {
      "type": "object",
      "properties": {
        "Limit": {
          "type": "integer",
          "format": "int32",
          "title": "Current Limit parameter, either passed by request or default value"
        },
        "CurrentOffset": {
          "type": "integer",
          "format": "int32",
          "title": "Current Offset value"
        },
        "Total": {
          "type": "integer",
          "format": "int32",
          "title": "Total number of records"
        },
        "CurrentPage": {
          "type": "integer",
          "format": "int32",
          "title": "Current number of Page"
        },
        "TotalPages": {
          "type": "integer",
          "format": "int32",
          "title": "Number of detected pages"
        },
        "NextOffset": {
          "type": "integer",
          "format": "int32",
          "title": "Offset value for next page if there is one"
        },
        "PrevOffset": {
          "type": "integer",
          "format": "int32",
          "title": "Offset value for previous page, if there is one"
        }
      },
      "title": "Generic container for responses sending pagination information"
    },
    "restProcess": {
      "type": "object",
      "properties": {
        "ID": {
          "type": "string"
        },
        "ParentID": {
          "type": "string"
        },
        "MetricsPort": {
          "type": "integer",
          "format": "int32"
        },
        "PeerId": {
          "type": "string"
        },
        "PeerAddress": {
          "type": "string"
        },
        "StartTag": {
          "type": "string"
        },
        "Services": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "restPutCellRequest": {
      "type": "object",
      "properties": {
        "Room": {
          "$ref": "#/definitions/restCell"
        },
        "CreateEmptyRoot": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "restPutShareLinkRequest": {
      "type": "object",
      "properties": {
        "ShareLink": {
          "$ref": "#/definitions/restShareLink"
        },
        "PasswordEnabled": {
          "type": "boolean",
          "format": "boolean"
        },
        "CreatePassword": {
          "type": "string"
        },
        "UpdatePassword": {
          "type": "string"
        },
        "UpdateCustomHash": {
          "type": "string"
        }
      }
    },
    "restPutUserMetaTagRequest": {
      "type": "object",
      "properties": {
        "Namespace": {
          "type": "string"
        },
        "Tag": {
          "type": "string"
        }
      }
    },
    "restPutUserMetaTagResponse": {
      "type": "object",
      "properties": {
        "Success": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "restRelationResponse": {
      "type": "object",
      "properties": {
        "SharedCells": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idmWorkspace"
          }
        },
        "BelongsToTeams": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idmRole"
          }
        }
      }
    },
    "restResetPasswordRequest": {
      "type": "object",
      "properties": {
        "ResetPasswordToken": {
          "type": "string"
        },
        "UserLogin": {
          "type": "string"
        },
        "NewPassword": {
          "type": "string"
        }
      }
    },
    "restResetPasswordResponse": {
      "type": "object",
      "properties": {
        "Success": {
          "type": "boolean",
          "format": "boolean"
        },
        "Message": {
          "type": "string"
        }
      }
    },
    "restResetPasswordTokenResponse": {
      "type": "object",
      "properties": {
        "Success": {
          "type": "boolean",
          "format": "boolean"
        },
        "Message": {
          "type": "string"
        }
      }
    },
    "restResourcePolicyQuery": {
      "type": "object",
      "properties": {
        "Type": {
          "$ref": "#/definitions/ResourcePolicyQueryQueryType"
        },
        "UserId": {
          "type": "string"
        }
      },
      "title": "Generic Query for limiting results based on resource permissions"
    },
    "restRestoreNodesRequest": {
      "type": "object",
      "properties": {
        "Nodes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/treeNode"
          }
        }
      }
    },
    "restRestoreNodesResponse": {
      "type": "object",
      "properties": {
        "RestoreJobs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/restBackgroundJobResult"
          }
        }
      }
    },
    "restRevokeRequest": {
      "type": "object",
      "properties": {
        "TokenId": {
          "type": "string"
        }
      },
      "title": "Rest request for revocation. Token is not mandatory, if not set\nrequest will use current JWT token"
    },
    "restRevokeResponse": {
      "type": "object",
      "properties": {
        "Success": {
          "type": "boolean",
          "format": "boolean"
        },
        "Message": {
          "type": "string"
        }
      },
      "title": "Rest response"
    },
    "restRolesCollection": {
      "type": "object",
      "properties": {
        "Roles": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idmRole"
          }
        }
      },
      "title": "Roles Collection"
    },
    "restSearchACLRequest": {
      "type": "object",
      "properties": {
        "Queries": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idmACLSingleQuery"
          }
        },
        "Offset": {
          "type": "string",
          "format": "int64"
        },
        "Limit": {
          "type": "string",
          "format": "int64"
        },
        "GroupBy": {
          "type": "integer",
          "format": "int32"
        },
        "CountOnly": {
          "type": "boolean",
          "format": "boolean"
        },
        "Operation": {
          "$ref": "#/definitions/serviceOperationType"
        }
      },
      "title": "Rest request for ACL's"
    },
    "restSearchResults": {
      "type": "object",
      "properties": {
        "Results": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/treeNode"
          }
        },
        "Total": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "restSearchRoleRequest": {
      "type": "object",
      "properties": {
        "Queries": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idmRoleSingleQuery"
          }
        },
        "ResourcePolicyQuery": {
          "$ref": "#/definitions/restResourcePolicyQuery"
        },
        "Offset": {
          "type": "string",
          "format": "int64"
        },
        "Limit": {
          "type": "string",
          "format": "int64"
        },
        "GroupBy": {
          "type": "integer",
          "format": "int32"
        },
        "CountOnly": {
          "type": "boolean",
          "format": "boolean"
        },
        "Operation": {
          "$ref": "#/definitions/serviceOperationType"
        }
      },
      "title": "Roles Search"
    },
    "restSearchUserRequest": {
      "type": "object",
      "properties": {
        "Queries": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idmUserSingleQuery"
          }
        },
        "ResourcePolicyQuery": {
          "$ref": "#/definitions/restResourcePolicyQuery"
        },
        "Offset": {
          "type": "string",
          "format": "int64"
        },
        "Limit": {
          "type": "string",
          "format": "int64"
        },
        "GroupBy": {
          "type": "integer",
          "format": "int32"
        },
        "CountOnly": {
          "type": "boolean",
          "format": "boolean"
        },
        "Operation": {
          "$ref": "#/definitions/serviceOperationType"
        }
      },
      "title": "Users Search"
    },
    "restSearchWorkspaceRequest": {
      "type": "object",
      "properties": {
        "Queries": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idmWorkspaceSingleQuery"
          }
        },
        "ResourcePolicyQuery": {
          "$ref": "#/definitions/restResourcePolicyQuery"
        },
        "Offset": {
          "type": "string",
          "format": "int64"
        },
        "Limit": {
          "type": "string",
          "format": "int64"
        },
        "GroupBy": {
          "type": "integer",
          "format": "int32"
        },
        "CountOnly": {
          "type": "boolean",
          "format": "boolean"
        },
        "Operation": {
          "$ref": "#/definitions/serviceOperationType"
        }
      },
      "title": "Rest request for searching workspaces"
    },
    "restServiceCollection": {
      "type": "object",
      "properties": {
        "Services": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ctlService"
          }
        },
        "Total": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "restSettingsEntry": {
      "type": "object",
      "properties": {
        "Key": {
          "type": "string"
        },
        "Label": {
          "type": "string"
        },
        "Description": {
          "type": "string"
        },
        "Manager": {
          "type": "string"
        },
        "Alias": {
          "type": "string"
        },
        "Metadata": {
          "$ref": "#/definitions/restSettingsEntryMeta"
        }
      }
    },
    "restSettingsEntryMeta": {
      "type": "object",
      "properties": {
        "IconClass": {
          "type": "string"
        },
        "Component": {
          "type": "string"
        },
        "Props": {
          "type": "string"
        },
        "Advanced": {
          "type": "boolean",
          "format": "boolean"
        },
        "Indexed": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "restSettingsMenuResponse": {
      "type": "object",
      "properties": {
        "RootMetadata": {
          "$ref": "#/definitions/restSettingsEntryMeta"
        },
        "Sections": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/restSettingsSection"
          }
        }
      }
    },
    "restSettingsSection": {
      "type": "object",
      "properties": {
        "Key": {
          "type": "string"
        },
        "Label": {
          "type": "string"
        },
        "Description": {
          "type": "string"
        },
        "Children": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/restSettingsEntry"
          }
        }
      }
    },
    "restShareLink": {
      "type": "object",
      "properties": {
        "Uuid": {
          "type": "string"
        },
        "LinkHash": {
          "type": "string"
        },
        "LinkUrl": {
          "type": "string"
        },
        "Label": {
          "type": "string"
        },
        "Description": {
          "type": "string"
        },
        "UserUuid": {
          "type": "string"
        },
        "UserLogin": {
          "type": "string"
        },
        "PasswordRequired": {
          "type": "boolean",
          "format": "boolean"
        },
        "AccessStart": {
          "type": "string",
          "format": "int64"
        },
        "AccessEnd": {
          "type": "string",
          "format": "int64"
        },
        "MaxDownloads": {
          "type": "string",
          "format": "int64"
        },
        "CurrentDownloads": {
          "type": "string",
          "format": "int64"
        },
        "ViewTemplateName": {
          "type": "string"
        },
        "TargetUsers": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/restShareLinkTargetUser"
          }
        },
        "RestrictToTargetUsers": {
          "type": "boolean",
          "format": "boolean"
        },
        "RootNodes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/treeNode"
          }
        },
        "Permissions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/restShareLinkAccessType"
          }
        },
        "Policies": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/serviceResourcePolicy"
          }
        },
        "PoliciesContextEditable": {
          "type": "boolean",
          "format": "boolean"
        }
      },
      "title": "Model for representing a public link"
    },
    "restShareLinkAccessType": {
      "type": "string",
      "enum": [
        "NoAccess",
        "Preview",
        "Download",
        "Upload"
      ],
      "default": "NoAccess",
      "title": "Known values for link permissions"
    },
    "restShareLinkTargetUser": {
      "type": "object",
      "properties": {
        "Display": {
          "type": "string"
        },
        "DownloadCount": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "restSubscriptionsCollection": {
      "type": "object",
      "properties": {
        "subscriptions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/activitySubscription"
          }
        }
      }
    },
    "restTemplate": {
      "type": "object",
      "properties": {
        "UUID": {
          "type": "string",
          "title": "Unique identifier for this template"
        },
        "Label": {
          "type": "string",
          "title": "Human friendly label"
        },
        "Node": {
          "$ref": "#/definitions/restTemplateNode",
          "title": "Template node, can be a file or a tree of folders"
        },
        "Policies": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/serviceResourcePolicy"
          },
          "title": "Associated policies"
        }
      },
      "title": "A template can be used to create files or tree from scratch"
    },
    "restTemplateNode": {
      "type": "object",
      "properties": {
        "IsFile": {
          "type": "boolean",
          "format": "boolean",
          "title": "Whether it's a file or a folder"
        },
        "BinaryUUUID": {
          "type": "string",
          "title": "If stored in binary store, the binary Uuid"
        },
        "EmbedPath": {
          "type": "string",
          "title": "If it's embedded in binary"
        },
        "Children": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/restTemplateNode"
          },
          "title": "One or more children"
        }
      },
      "title": "A template node is representing a file or a folder"
    },
    "restUpdateSharePoliciesRequest": {
      "type": "object",
      "properties": {
        "Uuid": {
          "type": "string"
        },
        "Policies": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/serviceResourcePolicy"
          }
        }
      }
    },
    "restUpdateSharePoliciesResponse": {
      "type": "object",
      "properties": {
        "Success": {
          "type": "boolean",
          "format": "boolean"
        },
        "Policies": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/serviceResourcePolicy"
          }
        },
        "PoliciesContextEditable": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "restUserBookmarksRequest": {
      "type": "object"
    },
    "restUserJobRequest": {
      "type": "object",
      "properties": {
        "JobName": {
          "type": "string"
        },
        "JsonParameters": {
          "type": "string"
        }
      }
    },
    "restUserJobResponse": {
      "type": "object",
      "properties": {
        "JobUuid": {
          "type": "string"
        }
      }
    },
    "restUserJobsCollection": {
      "type": "object",
      "properties": {
        "Jobs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/jobsJob"
          }
        }
      }
    },
    "restUserMetaCollection": {
      "type": "object",
      "properties": {
        "Metadatas": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idmUserMeta"
          }
        }
      },
      "title": "Collection of UserMeta"
    },
    "restUserMetaNamespaceCollection": {
      "type": "object",
      "properties": {
        "Namespaces": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idmUserMetaNamespace"
          }
        }
      },
      "title": "Collection of Meta Namespaces"
    },
    "restUserStateResponse": {
      "type": "object",
      "properties": {
        "Workspaces": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idmWorkspace"
          }
        },
        "WorkspacesAccesses": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        }
      }
    },
    "restUsersCollection": {
      "type": "object",
      "properties": {
        "Groups": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idmUser"
          }
        },
        "Users": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idmUser"
          }
        },
        "Total": {
          "type": "integer",
          "format": "int32"
        }
      },
      "title": "Users Collection"
    },
    "restVersioningPolicyCollection": {
      "type": "object",
      "properties": {
        "Policies": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/treeVersioningPolicy"
          }
        }
      }
    },
    "restWorkspaceCollection": {
      "type": "object",
      "properties": {
        "Workspaces": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idmWorkspace"
          }
        },
        "Total": {
          "type": "integer",
          "format": "int32"
        }
      },
      "title": "Rest response for workspace search"
    },
    "serviceOperationType": {
      "type": "string",
      "enum": [
        "OR",
        "AND"
      ],
      "default": "OR"
    },
    "serviceQuery": {
      "type": "object",
      "properties": {
        "SubQueries": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/protobufAny"
          }
        },
        "Operation": {
          "$ref": "#/definitions/serviceOperationType"
        },
        "ResourcePolicyQuery": {
          "$ref": "#/definitions/serviceResourcePolicyQuery"
        },
        "Offset": {
          "type": "string",
          "format": "int64"
        },
        "Limit": {
          "type": "string",
          "format": "int64"
        },
        "groupBy": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "serviceResourcePolicy": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "int64"
        },
        "Resource": {
          "type": "string"
        },
        "Action": {
          "$ref": "#/definitions/serviceResourcePolicyAction"
        },
        "Subject": {
          "type": "string"
        },
        "Effect": {
          "$ref": "#/definitions/serviceResourcePolicyPolicyEffect"
        },
        "JsonConditions": {
          "type": "string"
        }
      }
    },
    "serviceResourcePolicyAction": {
      "type": "string",
      "enum": [
        "ANY",
        "OWNER",
        "READ",
        "WRITE",
        "EDIT_RULES"
      ],
      "default": "ANY"
    },
    "serviceResourcePolicyPolicyEffect": {
      "type": "string",
      "enum": [
        "deny",
        "allow"
      ],
      "default": "deny"
    },
    "serviceResourcePolicyQuery": {
      "type": "object",
      "properties": {
        "Subjects": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "Empty": {
          "type": "boolean",
          "format": "boolean"
        },
        "Any": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "treeChangeLog": {
      "type": "object",
      "properties": {
        "Uuid": {
          "type": "string",
          "title": "Unique commit ID"
        },
        "Description": {
          "type": "string",
          "title": "Human-readable description of what happened"
        },
        "MTime": {
          "type": "string",
          "format": "int64",
          "title": "Unix Timestamp"
        },
        "Size": {
          "type": "string",
          "format": "int64",
          "title": "Content Size at that moment"
        },
        "Data": {
          "type": "string",
          "format": "byte",
          "title": "Arbitrary additional data"
        },
        "OwnerUuid": {
          "type": "string",
          "title": "Who performed this action"
        },
        "Event": {
          "$ref": "#/definitions/treeNodeChangeEvent",
          "title": "Event that triggered this change"
        }
      }
    },
    "treeGeoPoint": {
      "type": "object",
      "properties": {
        "Lat": {
          "type": "number",
          "format": "double"
        },
        "Lon": {
          "type": "number",
          "format": "double"
        }
      }
    },
    "treeGeoQuery": {
      "type": "object",
      "properties": {
        "Center": {
          "$ref": "#/definitions/treeGeoPoint",
          "title": "Either use a center point and a distance"
        },
        "Distance": {
          "type": "string",
          "description": "Example formats supported:\n\"5in\" \"5inch\" \"7yd\" \"7yards\" \"9ft\" \"9feet\" \"11km\" \"11kilometers\"\n\"3nm\" \"3nauticalmiles\" \"13mm\" \"13millimeters\" \"15cm\" \"15centimeters\"\n\"17mi\" \"17miles\" \"19m\" \"19meters\"\nIf the unit cannot be determined, the entire string is parsed and the\nunit of meters is assumed."
        },
        "TopLeft": {
          "$ref": "#/definitions/treeGeoPoint",
          "title": "Or use a bounding box with TopLeft and BottomRight points"
        },
        "BottomRight": {
          "$ref": "#/definitions/treeGeoPoint"
        }
      }
    },
    "treeListNodesRequest": {
      "type": "object",
      "properties": {
        "Node": {
          "$ref": "#/definitions/treeNode"
        },
        "Recursive": {
          "type": "boolean",
          "format": "boolean"
        },
        "Ancestors": {
          "type": "boolean",
          "format": "boolean"
        },
        "WithVersions": {
          "type": "boolean",
          "format": "boolean"
        },
        "WithCommits": {
          "type": "boolean",
          "format": "boolean"
        },
        "Limit": {
          "type": "string",
          "format": "int64"
        },
        "Offset": {
          "type": "string",
          "format": "int64"
        },
        "FilterType": {
          "$ref": "#/definitions/treeNodeType"
        }
      }
    },
    "treeNode": {
      "type": "object",
      "properties": {
        "Uuid": {
          "type": "string",
          "title": "------------------------------------\nCore identification of the node\n------------------------------------"
        },
        "Path": {
          "type": "string"
        },
        "Type": {
          "$ref": "#/definitions/treeNodeType"
        },
        "Size": {
          "type": "string",
          "format": "int64",
          "title": "Size of the file, or cumulated size of folder"
        },
        "MTime": {
          "type": "string",
          "format": "int64",
          "title": "Last modification Timestamp"
        },
        "Mode": {
          "type": "integer",
          "format": "int32",
          "title": "Permission mode, like 0777"
        },
        "Etag": {
          "type": "string",
          "title": "Hash of the content if node is a LEAF, Uuid or"
        },
        "Commits": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/treeChangeLog"
          },
          "title": "List of successive commits"
        },
        "MetaStore": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "------------------------------------\nThen a free K =\u003e V representation of any kind of metadata\n------------------------------------"
        },
        "AppearsIn": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/treeWorkspaceRelativePath"
          },
          "title": "Can be used for output when node is appearing in multiple workspaces"
        }
      }
    },
    "treeNodeChangeEvent": {
      "type": "object",
      "properties": {
        "Type": {
          "$ref": "#/definitions/NodeChangeEventEventType"
        },
        "Source": {
          "$ref": "#/definitions/treeNode"
        },
        "Target": {
          "$ref": "#/definitions/treeNode"
        },
        "Metadata": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        },
        "Silent": {
          "type": "boolean",
          "format": "boolean"
        },
        "Optimistic": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "treeNodeType": {
      "type": "string",
      "enum": [
        "UNKNOWN",
        "LEAF",
        "COLLECTION"
      ],
      "default": "UNKNOWN",
      "title": "==========================================================\n* Standard Messages\n=========================================================="
    },
    "treeQuery": {
      "type": "object",
      "properties": {
        "PathPrefix": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "Limit to a given subtree"
        },
        "MinSize": {
          "type": "string",
          "format": "int64",
          "title": "Range for size"
        },
        "MaxSize": {
          "type": "string",
          "format": "int64"
        },
        "MinDate": {
          "type": "string",
          "format": "int64",
          "title": "Range for date"
        },
        "MaxDate": {
          "type": "string",
          "format": "int64"
        },
        "Type": {
          "$ref": "#/definitions/treeNodeType",
          "title": "Limit to a given node type"
        },
        "FileName": {
          "type": "string",
          "title": "Search in filename"
        },
        "Content": {
          "type": "string",
          "title": "Search in content"
        },
        "FreeString": {
          "type": "string",
          "title": "Free Query String (for metadata)"
        },
        "Extension": {
          "type": "string",
          "title": "Search files by extension"
        },
        "GeoQuery": {
          "$ref": "#/definitions/treeGeoQuery",
          "title": "Search geographically"
        }
      },
      "title": "Search Queries"
    },
    "treeReadNodeRequest": {
      "type": "object",
      "properties": {
        "Node": {
          "$ref": "#/definitions/treeNode",
          "title": "Input node"
        },
        "WithCommits": {
          "type": "boolean",
          "format": "boolean",
          "title": "Gather commit information"
        },
        "WithExtendedStats": {
          "type": "boolean",
          "format": "boolean",
          "title": "Get extended stats - For folders, computes ChildrenCount"
        },
        "ObjectStats": {
          "type": "boolean",
          "format": "boolean",
          "title": "Used internally for the router ReadNode request, stat the datasource instead of index"
        }
      },
      "title": "Request / Responses Messages"
    },
    "treeReadNodeResponse": {
      "type": "object",
      "properties": {
        "Success": {
          "type": "boolean",
          "format": "boolean"
        },
        "Node": {
          "$ref": "#/definitions/treeNode"
        }
      }
    },
    "treeSearchRequest": {
      "type": "object",
      "properties": {
        "Query": {
          "$ref": "#/definitions/treeQuery",
          "title": "The query object"
        },
        "Size": {
          "type": "integer",
          "format": "int32",
          "title": "Limit the number of results"
        },
        "From": {
          "type": "integer",
          "format": "int32",
          "title": "Start at given position"
        },
        "Details": {
          "type": "boolean",
          "format": "boolean",
          "title": "Load node details"
        },
        "Facet": {
          "type": "string",
          "title": "Facet search"
        }
      }
    },
    "treeSyncChange": {
      "type": "object",
      "properties": {
        "seq": {
          "type": "string",
          "format": "uint64"
        },
        "nodeId": {
          "type": "string"
        },
        "type": {
          "$ref": "#/definitions/treeSyncChangeType"
        },
        "source": {
          "type": "string"
        },
        "target": {
          "type": "string"
        },
        "node": {
          "$ref": "#/definitions/treeSyncChangeNode"
        }
      }
    },
    "treeSyncChangeNode": {
      "type": "object",
      "properties": {
        "bytesize": {
          "type": "string",
          "format": "int64"
        },
        "md5": {
          "type": "string"
        },
        "mtime": {
          "type": "string",
          "format": "int64"
        },
        "nodePath": {
          "type": "string"
        },
        "repositoryIdentifier": {
          "type": "string"
        }
      }
    },
    "treeSyncChangeType": {
      "type": "string",
      "enum": [
        "unknown",
        "create",
        "delete",
        "path",
        "content"
      ],
      "default": "unknown"
    },
    "treeVersioningKeepPeriod": {
      "type": "object",
      "properties": {
        "IntervalStart": {
          "type": "string"
        },
        "MaxNumber": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "treeVersioningPolicy": {
      "type": "object",
      "properties": {
        "Uuid": {
          "type": "string"
        },
        "Name": {
          "type": "string"
        },
        "Description": {
          "type": "string"
        },
        "VersionsDataSourceName": {
          "type": "string"
        },
        "VersionsDataSourceBucket": {
          "type": "string"
        },
        "MaxTotalSize": {
          "type": "string",
          "format": "int64"
        },
        "MaxSizePerFile": {
          "type": "string",
          "format": "int64"
        },
        "IgnoreFilesGreaterThan": {
          "type": "string",
          "format": "int64"
        },
        "KeepPeriods": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/treeVersioningKeepPeriod"
          }
        }
      }
    },
    "treeWorkspaceRelativePath": {
      "type": "object",
      "properties": {
        "WsUuid": {
          "type": "string",
          "title": "Workspace Id"
        },
        "WsLabel": {
          "type": "string",
          "title": "Workspace Label"
        },
        "Path": {
          "type": "string",
          "title": "Relative Path inside workspace"
        }
      },
      "title": "Used in AppearsIn to signal a node is\nappearing in multiple workspaces in the current context"
    },
    "updateApplyUpdateRequest": {
      "type": "object",
      "properties": {
        "TargetVersion": {
          "type": "string",
          "title": "Version of the target binary"
        },
        "PackageName": {
          "type": "string",
          "title": "Name of the package if it's not the same as the current binary"
        }
      }
    },
    "updateApplyUpdateResponse": {
      "type": "object",
      "properties": {
        "Success": {
          "type": "boolean",
          "format": "boolean"
        },
        "Message": {
          "type": "string"
        }
      }
    },
    "updatePackage": {
      "type": "object",
      "properties": {
        "PackageName": {
          "type": "string",
          "title": "Name of the application"
        },
        "Version": {
          "type": "string",
          "title": "Version of this new binary"
        },
        "ReleaseDate": {
          "type": "integer",
          "format": "int32",
          "title": "Release date of the binary"
        },
        "Label": {
          "type": "string",
          "title": "Short human-readable description"
        },
        "Description": {
          "type": "string",
          "title": "Long human-readable description (markdown)"
        },
        "ChangeLog": {
          "type": "string",
          "title": "List or public URL of change logs"
        },
        "License": {
          "type": "string",
          "title": "License of this package"
        },
        "BinaryURL": {
          "type": "string",
          "title": "Https URL where to download the binary"
        },
        "BinaryChecksum": {
          "type": "string",
          "title": "Checksum of the binary to verify its integrity"
        },
        "BinarySignature": {
          "type": "string",
          "title": "Signature of the binary"
        },
        "BinaryHashType": {
          "type": "string",
          "title": "Hash type used for the signature"
        },
        "BinarySize": {
          "type": "string",
          "format": "int64",
          "title": "Size of the binary to download"
        },
        "BinaryOS": {
          "type": "string",
          "title": "GOOS value used at build time"
        },
        "BinaryArch": {
          "type": "string",
          "title": "GOARCH value used at build time"
        },
        "IsPatch": {
          "type": "boolean",
          "format": "boolean",
          "title": "Not used : if binary is a patch"
        },
        "PatchAlgorithm": {
          "type": "string",
          "title": "Not used : if a patch, how to patch (bsdiff support)"
        },
        "ServiceName": {
          "type": "string",
          "title": "Not used : at a point we may deliver services only updates"
        },
        "Status": {
          "$ref": "#/definitions/PackagePackageStatus"
        }
      }
    },
    "updateUpdateRequest": {
      "type": "object",
      "properties": {
        "Channel": {
          "type": "string",
          "title": "Channel name"
        },
        "PackageName": {
          "type": "string",
          "title": "Name of the currently running application"
        },
        "CurrentVersion": {
          "type": "string",
          "title": "Current version of the application"
        },
        "GOOS": {
          "type": "string",
          "title": "Current GOOS"
        },
        "GOARCH": {
          "type": "string",
          "title": "Current GOARCH"
        },
        "ServiceName": {
          "type": "string",
          "title": "Not Used : specific service to get updates for"
        },
        "LicenseInfo": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "For enterprise version, info about the current license"
        }
      }
    },
    "updateUpdateResponse": {
      "type": "object",
      "properties": {
        "Channel": {
          "type": "string"
        },
        "AvailableBinaries": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/updatePackage"
          },
          "title": "List of available binaries"
        }
      }
    }
  },
  "externalDocs": {
    "description": "More about Pydio Cells Apis",
    "url": "https://pydio.com"
  }
}
`