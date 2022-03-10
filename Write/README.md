
# 	InfluxDB - Activity
This activity provides your Flogo app the ability to insert data on InfluxDB v2.0 Cloud

This work is a derivative of that initial work at github.com/shaliniGovindaNayak

## Installation

```bash
flogo install github.com/leonobrien/InfluxDB-Flogo/Write
```
Link for flogo web:
```
github.com/leonobrien/InfluxDB-Flogo/Write
```

## Schema
Inputs and Outputs:

```json
"inputs":[
{
      "name": "Host",
      "type": "string",
      "required": true
    },
    {
      "name":"Organisation",
      "type":"string",
      "required":true
    },
    {
      "name":"Bucket",
      "type":"string",
      "required":true
    },
    {
       "name":"Token",
       "type":"any",
       "required":true
    },
    {
       "name":"Value",
       "type":"string"
    }
  ]
```
## Inputs
| Input                          | Description    |
|:-------------------------------|:---------------|
| Host                           | Host name along with port where influxDB is running.           |
| Organisation                   | InfluxDB 2 Cloud Organisation you are connecting to   |
| Bucket                         | Name of the bucket to use        |
| Token                          | API Security token to use         |
| Value                          | Data that needs to be inserted |


## Ouputs
| Output       | Description                                            |
|:-------------|:-------------------------------------------------------|
| Output       | The sucess message indicating the insertion |
