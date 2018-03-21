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
Get searches json for specified path. A path is in "/" syntax such as "/friends".
And for arrays, access by numbers such as "/friends/0".

```go
var jsonString = `
{
    "friends": [
        {
            "id": "1111",
            "name": "hlts2",
            "like": [
                "apple",
                "strawberry",
                "pineapple"
            ]
        },
        {
            "id": "2121",
            "name": "hiroto",
            "like": [
                "watermelon"
            ]
        }
    ]
}
`
g, _ := gson.NewGosonFromString(jsonString)

path := "/friends/1/like/0"
if g.HasWithPath(path) {
    result, _ := g.Path(path)
    str, _ := result.String()
    fmt.Println(str)
}
```

### Get Value（By Search）

### Indent String
