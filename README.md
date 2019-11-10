Gock HTTP Server
================
A controllable mock HTTP server for functional testing for services that do not 
expose an appropriate test environment.

## Usage

Just start the binary.

The following environment variables must be set:

* `BASE_PORT` - the base port to listen on, the channels are relative to this.

## Using mocks

Mocks are accessible within the `/mock` sub-path namespace. You should therefore configure your testing URI as
`http://127.0.0.1:$BASE_PORT$/mock/example-path`.

Mocks are matched based on rules, which is covered below.

## Configuring mocks
The `POST /api/mock/{mock}` API endpoint should be used to register mocks with the server.

Example payload:

```json
{
  "matchRule": {
    "type": "allOf",
    "rules": [
      {
        "type": "pathEquals", 
        "value": "foo/bar" 
      },
      {
        "type": "anyOf",
        "rules": [
          {
            "type": "methodEquals",
            "value": "POST"
          },
          {
            "type": "methodEquals",
            "value": "PUT"
          }
        ]       
      }
    ]
  },
  "response": {
    "statusCode": 201,
    "headers": {
      "Content-Type": "application/json; charset=utf-8"
    },
    "content": "{\"example\": \"response\"}"
  }  
}
```

A request of `http://127.0.0.1:$BASE_PORT$/mock/foo/bar` would then respond with:

```http request
HTTP/1.1 201 Created
content-length: 23
content-type: application/json; charset=utf-8
connection: keep-alive
keep-alive: timeout=15
date: Thu, 23 May 2019 16:42:28 GMT

{"example": "response"}
```

Full documentation TODO, these are likely to expand rather quickly.

## API reference

### `POST /api/reset`
Resets all mocks back to an empty state for the channel it is called on.

### `POST /api/mock/{mock}`
Creates a new or replaces an existing (identified by the `{mock}` fragment in the URI path) mock.
