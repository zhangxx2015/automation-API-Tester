
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
	"baseUrl":"https://app.ipdatacloud.com",
	"uuid":"uuid",
	"unixTimestamp":"unixTimestamp",
	"Nanosecond":"Nanosecond"
}
```

Write API request tasks in JSON format and save them as files

0001.json
``` json
{
    "url": "/v2/free_query",
    "querys": {
        "ip": "8.8.8.8"
    },
    "headers": {
    },
    "method": "GET",
    "body": {
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
2024/04/26 14:54:07 0001.json
2024/04/26 14:54:07     >>>>
2024/04/26 14:54:07      GET /v2/free_query
2024/04/26 14:54:07      {"ip":"8.8.8.8"}
2024/04/26 14:54:07      {"key1":"val1","key2":"val2"}
2024/04/26 14:54:08 response text:{"code":200,"data":{"city":"Mountain View","country":"US","country_english":"","ip":"8.8.8.8","isp":"谷歌公司","province":"California"},"msg":"success"}
2024/04/26 14:54:08     <<<<
2024/04/26 14:54:08      {
                          "code": 200,
                          "data": {
                            "city": "Mountain View",
                            "country": "US",        
                            "country_english": "",  
                            "ip": "8.8.8.8",        
                            "isp": "谷歌公司",      
                            "province": "California"
                          },
                          "msg": "success"
                        }
2024/04/26 14:54:08     assert
2024/04/26 14:54:08      null

2024/04/26 14:54:08 testing is passed
```

### Call result assertion

Edit the 'vars.json' file, which is used to define global variables

vars.json
``` json
{
    "baseUrl":"https://app.ipdatacloud.com",
    "uuid":"uuid",
    "unixTimestamp":"unixTimestamp",
    "Nanosecond":"Nanosecond"
}
```

Write API request tasks in JSON format and save them as files

0001.json
``` json
{
    "url": "/v2/free_query",
    "querys": {
        "ip": "8.8.8.8"
    },
    "headers": {
    },
    "method": "GET",
    "body": {
        "key1": "val1",
        "key2": "val2"
    }
}
```

0002.json
``` json
{
    "url": "/v2/free_query",
    "querys": {
        "ip": "114.114.114.114"
    },
    "headers": {
    },
    "method": "GET",
    "body": {
        "key1": "val1",
        "key2": "{{.unixTimestamp}}"
    },
    "assert":{
        "{{.response.code}}":"200"
    }
}
```

0003.json
``` json
{
    "url": "/v2/free_query",
    "querys": {
        "ip": "180.76.76.76"
    },
    "headers": {
    },
    "method": "GET",
    "body": {
        "key1": "val1",
        "key2": "{{.unixTimestamp}}"
    },
	"assert":{
		"{{.response.code}}":"200"
	}
}

```

Edit the tasks.json file, which is used to define all request tasks

tasks.json

``` json
[
	"0001.json",
    "0002.json",
    "0003.json"
]
```

Run test
``` bash
go run main.go
```


outputs

``` bash
2024/04/26 14:56:35 0001.json
2024/04/26 14:56:35     >>>>
2024/04/26 14:56:35      GET /v2/free_query
2024/04/26 14:56:35      {"ip":"8.8.8.8"}
2024/04/26 14:56:35      {"key1":"val1","key2":"val2"}
2024/04/26 14:56:35 response text:{"code":200,"data":{"city":"Mountain View","country":"US","country_english":"","ip":"8.8.8.8","isp":"谷歌公司","province":"California"},"msg":"success"}
2024/04/26 14:56:35     <<<<
2024/04/26 14:56:35      {
                          "code": 200,
                          "data": {
                            "city": "Mountain View",        
                            "country": "US",
                            "country_english": "",
                            "ip": "8.8.8.8",
                            "isp": "谷歌公司",
                            "province": "California"        
                          },
                          "msg": "success"
                        }
2024/04/26 14:56:35     assert
2024/04/26 14:56:35      null

2024/04/26 14:56:35 0002.json
2024/04/26 14:56:35     >>>>
2024/04/26 14:56:35      GET /v2/free_query                 
2024/04/26 14:56:35      {"ip":"114.114.114.114"}           
2024/04/26 14:56:35      {"key1":"val1","key2":"1714114595"}
2024/04/26 14:56:35 response text:{"code":200,"data":{"city":"","country":"CN","country_english":"China","ip":"114.114.114.114","isp":"移动","province":"江苏"},"msg":"success"}
2024/04/26 14:56:35     <<<<                                                                                                                                                    
2024/04/26 14:56:35      {                                                                                                                                                      
                          "code": 200,                                                                                                                                          
                          "data": {                                                                                                                                             
                            "city": "",                                                                                                                                         
                            "country": "CN",                                                                                                                                    
                            "country_english": "China",                                                                                                                         
                            "ip": "114.114.114.114",                                                                                                                            
                            "isp": "移动",                                                                                                                                      
                            "province": "江苏"                                                                                                                                  
                          },                                                                                                                                                    
                          "msg": "success"
                        }
2024/04/26 14:56:35     assert
2024/04/26 14:56:35      {"{{.response.code}}":"200"}
2024/04/26 14:56:35              express: {{.response.code}} key: 200 val: 200 pass: true

2024/04/26 14:56:35 0003.json
2024/04/26 14:56:35     >>>>
2024/04/26 14:56:35      GET /v2/free_query
2024/04/26 14:56:35      {"ip":"180.76.76.76"}
2024/04/26 14:56:35      {"key1":"val1","key2":"1714114595"}
2024/04/26 14:56:35 response text:{"code":200,"data":{"city":"北京","country":"CN","country_english":"China","ip":"180.76.76.76","isp":"百度","province":"北京"},"msg":"success"}
2024/04/26 14:56:35     <<<<
2024/04/26 14:56:35      {
                          "code": 200,
                          "data": {
                            "city": "北京",
                            "country": "CN",
                            "country_english": "China",
                            "ip": "180.76.76.76",
                            "isp": "百度",
                            "province": "北京"
                          },
                          "msg": "success"
                        }
2024/04/26 14:56:35     assert
2024/04/26 14:56:35      {"{{.response.code}}":"200"}
2024/04/26 14:56:35              express: {{.response.code}} key: 200 val: 200 pass: true

2024/04/26 14:56:35 testing is passed
```

### Request parameters from previous request

Edit the 'vars.json' file, which is used to define global variables

vars.json
``` json
{
	"baseUrl":"https://api.oioweb.cn",
	"uuid":"uuid",
	"unixTimestamp":"unixTimestamp",
	"Nanosecond":"Nanosecond"
}
```

Write API request tasks in JSON format and save them as files

0001.json
``` json
{
    "url": "/api/common/teladress",
    "querys": {
        "mobile": "13512345678"
    },
    "headers": {
    },
    "method": "GET",
    "body": {
    },
	"assert":{
		"{{.response.code}}":"200"
	}
}
```
0002.json
``` json
{
    "url": "/api/weather/weather",
    "querys": {
        "city_name": "{{.response.result.city}}",
		"pageNo": 1,
		"pageSize": 20
    },
    "headers": {
    },
    "method": "GET",
    "body": {
    },
	"assert":{
		"{{.response.code}}":"200"
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
2024/04/26 14:58:00 0001.json
2024/04/26 14:58:00     >>>>
2024/04/26 14:58:00      GET /api/common/teladress
2024/04/26 14:58:00      {"mobile":"13512345678"} 
2024/04/26 14:58:00      {}
2024/04/26 14:58:00 response text:{"code":200,"result":{"areaCode":"023","provCode":"500000","city":"重庆","cityCode":"500100","num":1351234,"name":"移动全球通卡","postCode":"400000","type":1,"prov":"重庆"},"msg":"success"}
2024/04/26 14:58:00     <<<<
2024/04/26 14:58:00      {
                          "code": 200,
                          "msg": "success",
                          "result": {
                            "areaCode": "023",
                            "city": "重庆",
                            "cityCode": "500100",
                            "name": "移动全球通卡",
                            "num": 1351234,
                            "postCode": "400000",
                            "prov": "重庆",
                            "provCode": "500000",
                            "type": 1
                          }
                        }
2024/04/26 14:58:00     assert
2024/04/26 14:58:00      {"{{.response.code}}":"200"}
2024/04/26 14:58:00              express: {{.response.code}} key: 200 val: 200 pass: true
                                                                                         
2024/04/26 14:58:00 0002.json                                                            
2024/04/26 14:58:00     >>>>                                                             
2024/04/26 14:58:00      GET /api/weather/weather                                        
2024/04/26 14:58:00      {"city_name":"重庆","pageNo":1,"pageSize":20}                   
2024/04/26 14:58:00      {}                                                              
2024/04/26 14:58:00 response text:{"code":200,"result":{"alert":null,"aqi":33,"city_name":"重庆","current_condition":"晴","current_temperature":26,"current_time":1714114003,"dat_condition":"多云","dat_high_temperature":30,"dat_low_temperature":19,"dat_weather_icon_id":"1","day_condition":"晴","download_icon":
1,"forecast_list":[{"condition":"多云转阴","date":"2024-04-25","high_temperature":"24","low_temperature":"18","weather_icon_id":"1","wind_direction":"南风","wind_level":"2"},{"condition":"晴转阴","date":"2024-04-26","high_temperature":"27","low_temperature":"18","weather_icon_id":"0","wind_direction":"东北风"
,"wind_level":"1"},{"condition":"小雨转多云","date":"2024-04-27","high_temperature":"27","low_temperature":"18","weather_icon_id":"7","wind_direction":"东南风","wind_level":"1"},{"condition":"多云转中雨","date":"2024-04-28","high_temperature":"30","low_temperature":"19","weather_icon_id":"1","wind_direction":
"东北风","wind_level":"1"},{"condition":"小雨","date":"2024-04-29","high_temperature":"22","low_temperature":"19","weather_icon_id":"7","wind_direction":"北风","wind_level":"1"},{"condition":"小雨","date":"2024-04-30","high_temperature":"20","low_temperature":"17","weather_icon_id":"7","wind_direction":"西北 
风","wind_level":"1"},{"condition":"小雨","date":"2024-05-01","high_temperature":"20","low_temperature":"17","weather_icon_id":"7","wind_direction":"西南风","wind_level":"1"},{"condition":"小雨","date":"2024-05-02","high_temperature":"19","low_temperature":"17","weather_icon_id":"7","wind_direction":"东风","w
ind_level":"1"},{"condition":"阴转小雨","date":"2024-05-03","high_temperature":"21","low_temperature":"16","weather_icon_id":"2","wind_direction":"北风","wind_level":"1"},{"condition":"小雨","date":"2024-05-04","high_temperature":"21","low_temperature":"17","weather_icon_id":"7","wind_direction":"北风","wind_
level":"1"},{"condition":"多云转阴","date":"2024-05-05","high_temperature":"25","low_temperature":"18","weather_icon_id":"1","wind_direction":"北风","wind_level":"1"},{"condition":"晴转阴","date":"2024-05-06","high_temperature":"29","low_temperature":"20","weather_icon_id":"0","wind_direction":"东北风","wind_
level":"1"},{"condition":"多云转小雨","date":"2024-05-07","high_temperature":"35","low_temperature":"23","weather_icon_id":"1","wind_direction":"东北风","wind_level":"1"},{"condition":"小雨","date":"2024-05-08","high_temperature":"30","low_temperature":"22","weather_icon_id":"7","wind_direction":"东北风","win
d_level":"2"},{"condition":"小雨转阴","date":"2024-05-09","high_temperature":"28","low_temperature":"21","weather_icon_id":"7","wind_direction":"西南风","wind_level":"1"},{"condition":"阴转小雨","date":"2024-05-10","high_temperature":"34","low_temperature":"23","weather_icon_id":"2","wind_direction":"东南风",
"wind_level":"2"}],"high_temperature":27,"hourly_forecast":[{"condition":"晴","hour":"14","temperature":"24","weather_icon_id":"0","wind_direction":"E","wind_level":"6.52"},{"condition":"晴","hour":"15","temperature":"25","weather_icon_id":"0","wind_direction":"NE","wind_level":"7.42"},{"condition":"晴","hour
":"16","temperature":"26","weather_icon_id":"0","wind_direction":"NE","wind_level":"7.42"},{"condition":"晴","hour":"17","temperature":"27","weather_icon_id":"0","wind_direction":"NE","wind_level":"9.29"},{"condition":"晴","hour":"18","temperature":"27","weather_icon_id":"0","wind_direction":"NE","wind_level"
:"9.29"},{"condition":"晴","hour":"19","temperature":"26","weather_icon_id":"0","wind_direction":"N","wind_level":"9.29"},{"condition":"晴","hour":"20","temperature":"25","weather_icon_id":"30","wind_direction":"N","wind_level":"9.29"},{"condition":"多云","hour":"21","temperature":"24","weather_icon_id":"31",
"wind_direction":"N","wind_level":"9.29"},{"condition":"多云","hour":"22","temperature":"23","weather_icon_id":"31","wind_direction":"N","wind_level":"7.42"},{"condition":"阴","hour":"23","temperature":"21","weather_icon_id":"2","wind_direction":"N","wind_level":"7.42"},{"condition":"阴","hour":"0","temperatu
re":"20","weather_icon_id":"2","wind_direction":"N","wind_level":"7.42"},{"condition":"阴","hour":"1","temperature":"19","weather_icon_id":"2","wind_direction":"E","wind_level":"5.62"},{"condition":"阴","hour":"2","temperature":"19","weather_icon_id":"2","wind_direction":"E","wind_level":"5.62"},{"condition":
"阴","hour":"3","temperature":"18","weather_icon_id":"2","wind_direction":"E","wind_level":"5.62"},{"condition":"阴","hour":"4","temperature":"18","weather_icon_id":"2","wind_direction":"E","wind_level":"3.71"},{"condition":"阴","hour":"5","temperature":"18","weather_icon_id":"2","wind_direction":"E","wind_le
vel":"3.71"},{"condition":"阴","hour":"6","temperature":"19","weather_icon_id":"2","wind_direction":"SE","wind_level":"5.62"},{"condition":"阴","hour":"7","temperature":"19","weather_icon_id":"2","wind_direction":"S","wind_level":"5.62"},{"condition":"阴","hour":"8","temperature":"18","weather_icon_id":"2","w
ind_direction":"S","wind_level":"7.42"},{"condition":"小雨","hour":"9","temperature":"18","weather_icon_id":"7","wind_direction":"S","wind_level":"7.42"},{"condition":"小雨","hour":"10","temperature":"19","weather_icon_id":"7","wind_direction":"S","wind_level":"7.42"},{"condition":"小雨","hour":"11","temperat
ure":"19","weather_icon_id":"7","wind_direction":"S","wind_level":"7.42"},{"condition":"小雨","hour":"12","temperature":"22","weather_icon_id":"7","wind_direction":"S","wind_level":"9.29"},{"condition":"小雨","hour":"13","temperature":"22","weather_icon_id":"7","wind_direction":"S","wind_level":"9.29"},{"cond
ition":"小雨","hour":"14","temperature":"24","weather_icon_id":"7","wind_direction":"E","wind_level":"3.20"}],"low_temperature":18,"moji_city_id":52,"night_condition":"阴","origin_data":[],"quality_level":"优","tips":"略微偏热，注意衣物变化。","tomorrow_aqi":95,"tomorrow_condition":"小雨转多云","tomorrow_high
_temperature":27,"tomorrow_low_temperature":18,"tomorrow_quality_level":"良","tomorrow_weather_icon_id":"7","update_time":"2024-04-26 14:16:03","weather_icon_id":"0","wind_direction":"东风","wind_level":2},"msg":"success"}
2024/04/26 14:58:00     <<<<
2024/04/26 14:58:00      {
                          "code": 200,
                          "msg": "success",
                          "result": {
                            "alert": null,
                            "aqi": 33,
                            "city_name": "重庆",
                            "current_condition": "晴",
                            "current_temperature": 26,
                            "current_time": 1714114003,
                            "dat_condition": "多云",
                            "dat_high_temperature": 30,
                            "dat_low_temperature": 19,
                            "dat_weather_icon_id": "1",
                            "day_condition": "晴",
                            "download_icon": 1,
                            "forecast_list": [
                              {
                                "condition": "多云转阴",
                                "date": "2024-04-25",
                                "high_temperature": "24",
                                "low_temperature": "18",
                                "weather_icon_id": "1",
                                "wind_direction": "南风",
                                "wind_level": "2"
                              },
                              {
                                "condition": "晴转阴",
                                "date": "2024-04-26",
                                "high_temperature": "27",
                                "low_temperature": "18",
                                "weather_icon_id": "0",
                                "wind_direction": "东北风",
                                "wind_level": "1"
                              },
                              {
                                "condition": "小雨转多云",
                                "date": "2024-04-27",
                                "high_temperature": "27",
                                "low_temperature": "18",
                                "weather_icon_id": "7",
                                "wind_direction": "东南风",
                                "wind_level": "1"
                              },
                              {
                                "condition": "多云转中雨",
                                "date": "2024-04-28",
                                "high_temperature": "30",
                                "low_temperature": "19",
                                "weather_icon_id": "1",
                                "wind_direction": "东北风",
                                "wind_level": "1"
                              },
                              {
                                "condition": "小雨",
                                "date": "2024-04-29",
                                "high_temperature": "22",
                                "low_temperature": "19",
                                "weather_icon_id": "7",
                                "wind_direction": "北风",
                                "wind_level": "1"
                              },
                              {
                                "condition": "小雨",
                                "date": "2024-04-30",
                                "high_temperature": "20",
                                "low_temperature": "17",
                                "weather_icon_id": "7",
                                "wind_direction": "西北风",
                                "wind_level": "1"
                              },
                              {
                                "condition": "小雨",
                                "date": "2024-05-01",
                                "high_temperature": "20",
                                "low_temperature": "17",
                                "weather_icon_id": "7",
                                "wind_direction": "西南风",
                                "wind_level": "1"
                              },
                              {
                                "condition": "小雨",
                                "date": "2024-05-02",
                                "high_temperature": "19",
                                "low_temperature": "17",
                                "weather_icon_id": "7",
                                "wind_direction": "东风",
                                "wind_level": "1"
                              },
                              {
                                "condition": "阴转小雨",
                                "date": "2024-05-03",
                                "high_temperature": "21",
                                "low_temperature": "16",
                                "weather_icon_id": "2",
                                "wind_direction": "北风",
                                "wind_level": "1"
                              },
                              {
                                "condition": "小雨",
                                "date": "2024-05-04",
                                "high_temperature": "21",
                                "low_temperature": "17",
                                "weather_icon_id": "7",
                                "wind_direction": "北风",
                                "wind_level": "1"
                              },
                              {
                                "condition": "多云转阴",
                                "date": "2024-05-05",
                                "high_temperature": "25",
                                "low_temperature": "18",
                                "weather_icon_id": "1",
                                "wind_direction": "北风",
                                "wind_level": "1"
                              },
                              {
                                "condition": "晴转阴",
                                "date": "2024-05-06",
                                "high_temperature": "29",
                                "low_temperature": "20",
                                "weather_icon_id": "0",
                                "wind_direction": "东北风",
                                "wind_level": "1"
                              },
                              {
                                "condition": "多云转小雨",
                                "date": "2024-05-07",
                                "high_temperature": "35",
                                "low_temperature": "23",
                                "weather_icon_id": "1",
                                "wind_direction": "东北风",
                                "wind_level": "1"
                              },
                              {
                                "condition": "小雨",
                                "date": "2024-05-08",
                                "high_temperature": "30",
                                "low_temperature": "22",
                                "weather_icon_id": "7",
                                "wind_direction": "东北风",
                                "wind_level": "2"
                              },
                              {
                                "condition": "小雨转阴",
                                "date": "2024-05-09",
                                "high_temperature": "28",
                                "low_temperature": "21",
                                "weather_icon_id": "7",
                                "wind_direction": "西南风",
                                "wind_level": "1"
                              },
                              {
                                "condition": "阴转小雨",
                                "date": "2024-05-10",
                                "high_temperature": "34",
                                "low_temperature": "23",
                                "weather_icon_id": "2",
                                "wind_direction": "东南风",
                                "wind_level": "2"
                              }
                            ],
                            "high_temperature": 27,
                            "hourly_forecast": [
                              {
                                "condition": "晴",
                                "hour": "14",
                                "temperature": "24",
                                "weather_icon_id": "0",
                                "wind_direction": "E",
                                "wind_level": "6.52"
                              },
                              {
                                "condition": "晴",
                                "hour": "15",
                                "temperature": "25",
                                "weather_icon_id": "0",
                                "wind_direction": "NE",
                                "wind_level": "7.42"
                              },
                              {
                                "condition": "晴",
                                "hour": "16",
                                "temperature": "26",
                                "weather_icon_id": "0",
                                "wind_direction": "NE",
                                "wind_level": "7.42"
                              },
                              {
                                "condition": "晴",
                                "hour": "17",
                                "temperature": "27",
                                "weather_icon_id": "0",
                                "wind_direction": "NE",
                                "wind_level": "9.29"
                              },
                              {
                                "condition": "晴",
                                "hour": "18",
                                "temperature": "27",
                                "weather_icon_id": "0",
                                "wind_direction": "NE",
                                "wind_level": "9.29"
                              },
                              {
                                "condition": "晴",
                                "hour": "19",
                                "temperature": "26",
                                "weather_icon_id": "0",
                                "wind_direction": "N",
                                "wind_level": "9.29"
                              },
                              {
                                "condition": "晴",
                                "hour": "20",
                                "temperature": "25",
                                "weather_icon_id": "30",
                                "wind_direction": "N",
                                "wind_level": "9.29"
                              },
                              {
                                "condition": "多云",
                                "hour": "21",
                                "temperature": "24",
                                "weather_icon_id": "31",
                                "wind_direction": "N",
                                "wind_level": "9.29"
                              },
                              {
                                "condition": "多云",
                                "hour": "22",
                                "temperature": "23",
                                "weather_icon_id": "31",
                                "wind_direction": "N",
                                "wind_level": "7.42"
                              },
                              {
                                "condition": "阴",
                                "hour": "23",
                                "temperature": "21",
                                "weather_icon_id": "2",
                                "wind_direction": "N",
                                "wind_level": "7.42"
                              },
                              {
                                "condition": "阴",
                                "hour": "0",
                                "temperature": "20",
                                "weather_icon_id": "2",
                                "wind_direction": "N",
                                "wind_level": "7.42"
                              },
                              {
                                "condition": "阴",
                                "hour": "1",
                                "temperature": "19",
                                "weather_icon_id": "2",
                                "wind_direction": "E",
                                "wind_level": "5.62"
                              },
                              {
                                "condition": "阴",
                                "hour": "2",
                                "temperature": "19",
                                "weather_icon_id": "2",
                                "wind_direction": "E",
                                "wind_level": "5.62"
                              },
                              {
                                "condition": "阴",
                                "hour": "3",
                                "temperature": "18",
                                "weather_icon_id": "2",
                                "wind_direction": "E",
                                "wind_level": "5.62"
                              },
                              {
                                "condition": "阴",
                                "hour": "4",
                                "temperature": "18",
                                "weather_icon_id": "2",
                                "wind_direction": "E",
                                "wind_level": "3.71"
                              },
                              {
                                "condition": "阴",
                                "hour": "5",
                                "temperature": "18",
                                "weather_icon_id": "2",
                                "wind_direction": "E",
                                "wind_level": "3.71"
                              },
                              {
                                "condition": "阴",
                                "hour": "6",
                                "temperature": "19",
                                "weather_icon_id": "2",
                                "wind_direction": "SE",
                                "wind_level": "5.62"
                              },
                              {
                                "condition": "阴",
                                "hour": "7",
                                "temperature": "19",
                                "weather_icon_id": "2",
                                "wind_direction": "S",
                                "wind_level": "5.62"
                              },
                              {
                                "condition": "阴",
                                "hour": "8",
                                "temperature": "18",
                                "weather_icon_id": "2",
                                "wind_direction": "S",
                                "wind_level": "7.42"
                              },
                              {
                                "condition": "小雨",
                                "hour": "9",
                                "temperature": "18",
                                "weather_icon_id": "7",
                                "wind_direction": "S",
                                "wind_level": "7.42"
                              },
                              {
                                "condition": "小雨",
                                "hour": "10",
                                "temperature": "19",
                                "weather_icon_id": "7",
                                "wind_direction": "S",
                                "wind_level": "7.42"
                              },
                              {
                                "condition": "小雨",
                                "hour": "11",
                                "temperature": "19",
                                "weather_icon_id": "7",
                                "wind_direction": "S",
                                "wind_level": "7.42"
                              },
                              {
                                "condition": "小雨",
                                "hour": "12",
                                "temperature": "22",
                                "weather_icon_id": "7",
                                "wind_direction": "S",
                                "wind_level": "9.29"
                              },
                              {
                                "condition": "小雨",
                                "hour": "13",
                                "temperature": "22",
                                "weather_icon_id": "7",
                                "wind_direction": "S",
                                "wind_level": "9.29"
                              },
                              {
                                "condition": "小雨",
                                "hour": "14",
                                "temperature": "24",
                                "weather_icon_id": "7",
                                "wind_direction": "E",
                                "wind_level": "3.20"
                              }
                            ],
                            "low_temperature": 18,
                            "moji_city_id": 52,
                            "night_condition": "阴",
                            "origin_data": [],
                            "quality_level": "优",
                            "tips": "略微偏热，注意衣物变化。",
                            "tomorrow_aqi": 95,
                            "tomorrow_condition": "小雨转多云",
                            "tomorrow_high_temperature": 27,
                            "tomorrow_low_temperature": 18,
                            "tomorrow_quality_level": "良",
                            "tomorrow_weather_icon_id": "7",
                            "update_time": "2024-04-26 14:16:03",
                            "weather_icon_id": "0",
                            "wind_direction": "东风",
                            "wind_level": 2
                          }
                        }
2024/04/26 14:58:00     assert
2024/04/26 14:58:00      {"{{.response.code}}":"200"}
2024/04/26 14:58:00              express: {{.response.code}} key: 200 val: 200 pass: true

2024/04/26 14:58:00 testing is passed


```


#