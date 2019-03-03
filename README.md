# gson  [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT) [![Go Report Card](https://goreportcard.com/badge/github.com/hlts2/gson)](https://goreportcard.com/report/github.com/hlts2/gson) [![GoDoc](http://godoc.org/github.com/hlts2/gson?status.svg)](http://godoc.org/github.com/hlts2/gson) [![Join the chat at https://gitter.im/hlts2/gson](https://badges.gitter.im/hlts2/gson.svg)](https://gitter.im/hlts2/gson?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
Get JSON values quickly

## Requirement
- go (>= 1.11)

## Example

### Create from byte

Create gson object from `[]byte`. Returns an error if the bytes are not valid json.

```go
g, err := gson.CreateWithBytes(data)
```

### Create from io.Reader

Create gson object from a `io.Reader`. Returns an error if the resp.Body are not valid json.

```go
g, err := gson.CreateWithReader(resp.Body)
```

### Get value by path

`GetByPath` gets json value for specified path. The path is syntax such as `created_at.date`. And if you want to get the elements of json array, please put number in keys such as `likes.0`.

```go

var json = `
{
    "id": "1111",
    "name": "hlts2",
    "likes": [
        "apple",
        "strawberry",
        "pineapple"
    ],
    "created_at": {
        "date": "2017-05-10 12:54:18",
        "timezone": "UTC"
    }
}
`

g, _ := gson.CreateWithBytes([]byte(json))

result, _ := g.GetByPath("likes.1")

str := result.String() // str, err := result.StringE()

fmt.Println(str) //strawberry

```

#### Path and Syntax


```json
{  
    "name":{"first":"little", "last":"tiny"},
    "age":24,
    "children":["hiroto", "haruki"],
    "friends":[  
        {  
            "id":0,
            "name":"Beck Hansen"
        },
        {  
            "id":1,
            "name":"Cleveland Gomez"
        },
        {  
            "id":2,
            "name":"Norton Duncan"
        }
    ]
}
```
```
"name.last"          >> "tiny"
"age"                >> 24
"children"           >> ["hiroto", "haruki"]
"children.1"         >> "haruki"
"friends.1"          >> {"id":1,"name":"Cleveland Gomez"}
"friends.1.id"       >> 1
"friends.#.id"       >> [0, 1, 2]
```

### Get value by keys

`GetByKeys` gets json value for specified keys. keys are given as string slice such as `[]string{"created_at", "date"}`. And if you want to get the elements of json array, please put number in keys such as `[]string{"likes", "0"}`.

```go
var json = `
{
    "id": "1111",
    "name": "hlts2",
    "likes": [
        "apple",
        "strawberry",
        "pineapple"
    ],
    "created_at": {
        "date": "2017-05-10 12:54:18",
        "timezone": "UTC"
    }
}
`

g, _ := gson.CreateWithBytes([]byte(json))

result, _ := g.GetByKeys("likes", "1")

str := result.String() // str, err := result.StringE()

fmt.Println(str) //strawberry

```

### Iterating objects

```go
json := `{"created_at": {"date": "2017-05-10 12:54:18"}}`

g, _ := gson.CreateWithBytes([]byte(json))

result, _ :=  g.GetByKeys("created_at")

m := result.Map() // m, err := result.MapE()

for key, value := range m {
    fmt.Printf("key: %s, value: %v", key, value.Interface{}) //key: date, value: 2017-05-10 12:54:18
}

```

### Iterating slice

```go

json := `{"Likes": ["pen"]}`

g, _ := gson.CreateWithBytes([]byte(json))

result, _ :=  g.GetByKeys("Likes")

s := result.Slice() // s, err := result.SliceE()

for _, value := range s {
    fmt.Printf("value: %v", value.Interface()) //value: pen
}

```

### Indent String

`Indent` returns the formatted json string

```go

json := `{"Accounts": [{"ID": "1111"}, {"ID": "2222"}]}`

g, _ := gson.CreateWithBytes([]byte(json))

var buf bytes.Buffer
g.Indent(&buf, "", "  ")

fmt.Println(buf.String())
/*
{
    "IDs": [
        {
            "ID": "1111"
        },
        {
            "ID": "2222"
        }
    ]
}
*/
```
