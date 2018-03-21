package main

import (
	"fmt"
	"log"

	"github.com/hlts2/gson"
)

var jsonString = `
{"name": "hlts2"}
`

func main() {
	g, err := gson.NewGosonFromString(jsonString)
	if err != nil {
		log.Fatal(err)
	}

	result, err := g.Path("/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
