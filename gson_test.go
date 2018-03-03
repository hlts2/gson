package gson

import (
	"strings"
	"testing"
)

type NewTest struct {
	json    string
	isGoson bool
	isErr   bool
}

type NewTests []NewTest

func (n NewTest) CheckError(err error) bool {
	if n.isErr && err != nil {
		return true
	}

	if !n.isErr && err == nil {
		return true
	}

	return false
}

func (n NewTest) CheckGoson(g *Goson) bool {
	if n.isGoson && g != nil {
		return true
	}

	if !n.isGoson && g == nil {
		return true
	}
	return false
}

var newTests = NewTests{
	NewTest{
		json:    `1`,
		isGoson: true,
		isErr:   false,
	},
	NewTest{
		json:    `"2"`,
		isGoson: true,
		isErr:   false,
	},
	NewTest{
		json:    `{"picture": "http://hogehoge"}`,
		isGoson: true,
		isErr:   false,
	},
	NewTest{
		json:    `{afsf: adfaasf`,
		isGoson: false,
		isErr:   true,
	},
	NewTest{
		json: `
		{"friends": [
      		{
        		"id": 0,
				"name": "hiro"
			},
			{
				"id": 1,
				"name": "hiroto"
			},
			{
				"id": 2,
				"name": "hlts2"
			}
		]}
	  `,
		isGoson: true,
		isErr:   false,
	},
	NewTest{
		json:    `[{"name": "little"}, {"name": "tiny"}]`,
		isGoson: true,
		isErr:   false,
	},
	NewTest{
		json:    `[{"name": "litt]`,
		isGoson: false,
		isErr:   true,
	},
}

func TestNewGosonFromString(t *testing.T) {
	for _, test := range newTests {
		g, err := NewGosonFromString(test.json)

		if !test.CheckError(err) {
			t.Error("")
		}

		if !test.CheckGoson(g) {
			t.Error("")
		}
	}
}

func TestNewGosonFromByte(t *testing.T) {
	for _, test := range newTests {
		g, err := NewGosonFromByte([]byte(test.json))

		if !test.CheckError(err) {
			t.Error("")
		}

		if !test.CheckGoson(g) {
			t.Error("")
		}
	}
}

func TestNewGosonFromReader(t *testing.T) {
	for _, test := range newTests {
		g, err := NewGosonFromReader(strings.NewReader(test.json))

		if !test.CheckError(err) {
			t.Error("")
		}

		if !test.CheckGoson(g) {
			t.Error("")
		}
	}
}
