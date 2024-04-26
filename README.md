
This is a API testing tool. 🚀

## Features

* RESTful api calls
* Request orchestration
* Global variable rendering
* Method Chaining
* Call result assertion
* golang template

## Get started

### basic usage

Edit the 'vars.json' file, which is used to define global variables

vars.json
``` json
{
	"baseUrl":"https://api.ipify.org",
	"uuid":"uuid",
	"unixTimestamp":"unixTimestamp",
	"Nanosecond":"Nanosecond"
}
```

Write API request tasks in JSON format and save them as files

0001.json
``` json
{
	"url":"/",
	"querys":{
        "format":"json"
    },
	"headers":{},
	"method":"GET",
	"body":{
	  "key1": "val1",
      "key2": "val2"
	}
}
```

Edit the tasks.json file, which is used to define all request tasks

tasks.json

``` json
[
	"0001.json"
]
```

Run test
``` bash
go run main.go
```

outputs

``` bash
API server listening at: 127.0.0.1:14406
2024/04/26 11:32:54 0001.json
2024/04/26 11:32:54     >>>>
2024/04/26 11:32:54      GET /
2024/04/26 11:32:54      {"format":"json"}
2024/04/26 11:32:54      {"key1":"val1","key2":"val2"}
2024/04/26 11:33:09 {"ip":"117.14.58.46"}
2024/04/26 11:33:13     <<<<
2024/04/26 11:33:13      {
                          "ip": "117.14.58.46"
                        }
2024/04/26 11:33:13     !!!!
2024/04/26 11:33:13      null

2024/04/26 11:33:13 testing is passed
```

### Call result assertion

Edit the 'vars.json' file, which is used to define global variables

vars.json
``` json
{
	"baseUrl":"https://api.ipify.org",
	"uuid":"uuid",
	"unixTimestamp":"unixTimestamp",
	"Nanosecond":"Nanosecond"
}
```

Write API request tasks in JSON format and save them as files

0001.json
``` json
{
	"url":"/",
	"querys":{
        "format":"json"
    },
	"headers":{},
	"method":"GET",
	"body":{
	  "key1": "val1",
      "key2": "val2"
	}
}
```

0002.json
``` json
{
	"url":"/",
	"querys":{
        "format":"json"
    },
	"headers":{},
	"method":"GET",
	"body":{
	  "key1": "val1",
      "key2": "{{.unixTimestamp}}"
	}
}
```

Edit the tasks.json file, which is used to define all request tasks

tasks.json

``` json
[
	"0001.json",
    "0002.json"
]
```

Run test
``` bash
go run main.go
```


outputs

``` bash
API server listening at: 127.0.0.1:14406
2024/04/26 11:32:54 0001.json
2024/04/26 11:32:54     >>>>
2024/04/26 11:32:54      GET /
2024/04/26 11:32:54      {"format":"json"}
2024/04/26 11:32:54      {"key1":"val1","key2":"val2"}
2024/04/26 11:33:09 {"ip":"117.14.58.46"}
2024/04/26 11:33:13     <<<<
2024/04/26 11:33:13      {
                          "ip": "117.14.58.46"
                        }
2024/04/26 11:33:13     !!!!
2024/04/26 11:33:13      null

2024/04/26 11:33:15 0002.json
2024/04/26 11:33:15     >>>>
2024/04/26 11:33:15      GET /
2024/04/26 11:33:16      {"format":"json"}
2024/04/26 11:33:16      {"key1":"val1","key2":"val2"}
2024/04/26 11:33:17 {"ip":"117.14.58.46"}
2024/04/26 11:33:17     <<<<
2024/04/26 11:33:17      {
                          "ip": "117.14.58.46"
                        }
2024/04/26 11:33:18     !!!!
2024/04/26 11:33:18      null


2024/04/26 11:33:13 testing is passed
```