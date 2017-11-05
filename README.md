# gogen

[![Go Report Card](https://goreportcard.com/badge/github.com/lwhile/gogen)](https://goreportcard.com/report/github.com/lwhile/gogen)

gogen is a tool used to generate JSON structure definition automatically.

You can write json files to generate Go's struct definitions

## Status 

Developing 

## EXample

```json
{
    "key1":1,
    "key2":[1,2,3],
    "key3":"Hello",
    "key4":{
        "key4A":1,
        "key4B":[1,2,3],
        "key4C":"Hello",
        "key4D":{
            "key4DA":"123"
        }
    },
    "key5":["A","B","C"],
    "key6":[
        {"key6A":123},
        {"key6A":123}
    ]
}
    
```

use gogen you can get this go file:

```go
package main

type T struct{
    key3  string
    key4  struct{
        key4A  float64
        key4B  []float64
        key4C  string
        key4D  struct{
            key4DA  string
        }
    }
    key5  []string
    key6  []struct{
        key6A  float64
    }
    key1  float64
    key2  []float64
}

```

## usage

```
go get -v -u github.com/lwhile/gogen

cd $GOPATH/src/github.com/lwhile/gogen

go build -o gogen cmd/cmd.go
```

```
./gogen -input /usr/local/Go/src/github.com/lwhile/gogen/main.json -output ./A.go
```

Or start gogen as a http service:

```
./gogen -web 
```

Open [http://127.0.0.1:4928/static/index.html](http://127.0.0.1:4928/static/index.html)


## rode map
- [x] Use fixed order

- [ ] Add json tag 

- [ ] Converts static file into managable Go source code

- [ ] Import web front end page 