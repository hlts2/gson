package main

import (
	"fmt"
	"log"

	"github.com/hlts2/gson"
)

var jsonString = `
{"friends": [{"id": "1111", "name": "hlts2", "like": ["apple", "strawberry", "pineapple"]}, {"id": "2121", "name": "hiroto", "like": ["watermelon"]}]}
`

func main() {
	g, err := gson.NewGosonFromString(jsonString)
	if err != nil {
		log.Fatal(err)
	}

	result, _ := g.Path("/friends/1/like/0")
	str, _ := result.String()
	fmt.Println(str)

	result, _ = g.Path("/friends")
	fmt.Println(result.Indent("", " "))
}
