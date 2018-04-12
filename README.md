# gson
gson is simple json parse library for go.
gson is inspired by [jason](https://github.com/antonholmquist/jason).

## Requirement
Go 1.8

## Installation
```shell
go get github.com/hlts2/gson
```
## Example

### Create from byte

Create gson object from `[]byte`. Returns an error if the bytes are not valid json.

```go
g, err := gson.NewGsonFromByte(data)
```

### Create from io.Reader

Create gson object from a `io.Reader`. Returns an error if the resp.Body are not valid json.

```go
g, err := gson.NewGsonFromReader(resp.Body)
```

### Get value by path

`GetByPath` gets json value for specified path. The path is in "." syntax such as "created_at.date". And if you want to get the elements of json array, please put number in keys such as "likes.0".

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

g, _ := gson.NewGosonFromByte([]byte(json))

result, _ := g.GetByPath("likes.1")

str, _ := result.String()

fmt.Println(str) //strawberry

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

g, _ := gson.NewGosonFromByte([]byte(json))

result, _ := g.GetByKeys("likes", "1")

str, _ := result.String()

fmt.Println(str) //strawberry

```

### Indent String

`Indent` returns the formatted json string

```go
/*
json := `{"Accounts": [{"ID": "1111"}, {"ID": "2222"}]}`
*/

str, _  := g.Indent("", "  ")

fmt.Println(str)
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
