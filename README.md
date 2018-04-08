# gson
Simple Json Parse Library for Go

## Requirement
Go 1.8

## Installation
```shell
go get github.com/hlts2/gson
```
## Example

### Get Value（By Path）
Get searches json for specified path. A path is in "." syntax such as "created_at.date".

And for arrays, access by numbers such as "likes.0".

```go
var jsonString = `
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
g, _ := gson.NewGosonFromString(jsonString)

path := "friends.1.like.0"
if g.HasWithPath(path) {
    result, _ := g.Path(path)
    str, _ := result.String()
    fmt.Println(str)
}
```

### Get Value（By Search）

### Indent String
