package main

import (
	"fmt"
	"log"

	"github.com/hlts2/gson"
)

var json = `
{"friends": [{"id": "1111", "name": "hlts2", "like": ["apple", "strawberry", "pineapple"]}, {"id": "2121", "name": "hiroto", "like": ["watermelon"]}]}
`

func main() {
	g, err := gson.CreateWithBytes([]byte(json))
	if err != nil {
		log.Fatal(err)
	}

	result, err := g.GetByKeys("friends")
	if err != nil {
		log.Fatal(err)
	}

	s, err := result.SliceE()
	if err != nil {
		log.Fatal(err)
	}

	for _, value := range s {
		fmt.Printf("value: %v\n", value.Interface())
	}
}
