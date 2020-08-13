# pocket4d-server

This is pocket4d testing server only for POC 0 usage.

## Build binary

```shell script
go build -o pocket4d-server main/main.go
```

## Setup storage

Currently, we store all mini program file into a directly called `bundled` which is in the same level of the executable 

binary (we can make a parameter for this in the future).

## Run server

```shell script
./pocket4d-server
```

## API endpoint

### Upload bundled

**/api/v1/bundled/{name}**

*Upload a mini program bundled, *

Request: form-data

key: bundled     
value: path of the file

Response:

Sample

```json
{
    "app_id": "352a66d5-5cb4-46ac-83b9-7bf904d72707",
    "name": "mini1"
}
```

