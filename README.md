# pocket4d-server

This is pocket4d testing server only for POC 0 usa.

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

Request: 

POST **/api/v1/bundled/{name}**

form-data
key: bundled     
value: path of the file

curl example 'curl -F bundled=@mini1.json localhost:8080/api/v1/bundled/mini'

Response:

Sample

```json
{
    "app_id": "352a66d5-5cb4-46ac-83b9-7bf904d72707",
    "name": "mini1"
}
```

### List all mini program bundled

Request: 

GET **/api/v1/bundled**

Response:

```json
[
    {
        "AppId": "ae885b94-abee-4542-91b6-53ca658382d7",
        "Name": "mini1"
    },
    {
        "AppId": "807540ae-2653-45f3-91f9-d9745a298150",
        "Name": "mini2"
    }
]
```

### Download bundled by appid and name

Request: 

GET **/api/v1/bundled/{appId}/{name}**
