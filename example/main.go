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
	g, err := gson.NewGsonFromByte([]byte(jsonString))
	if err != nil {
		log.Fatal(err)
	}

	result, err := g.GetByKeys("friends")
	if err != nil {
		log.Fatal(err)
	}

	slice, err := result.Slice()
	if err != nil {
		log.Fatal(err)
	}

	for _, value := range slice {
		fmt.Printf("value: %v\n", value.Interface())
	}

	m, err := slice[0].Map()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(m["name"].Interface())
}
