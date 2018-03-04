package gson

import (
	"strings"
	"testing"
)

type TestData struct {
	json    string
	isGoson bool
	data    interface{}
	isErr   bool
}

type TestDatas []TestData

func (n TestData) CheckError(err error) bool {
	if n.isErr && err != nil {
		return true
	}

	if !n.isErr && err == nil {
		return true
	}

	return false
}

func (n TestData) CheckGoson(g *Goson) bool {
	if n.isGoson && g != nil {
		return true
	}

	if !n.isGoson && g == nil {
		return true
	}
	return false
}

var newTests = TestDatas{
	TestData{
		json:    `1`,
		isGoson: true,
		isErr:   false,
	},
	TestData{
		json:    `"2"`,
		isGoson: true,
		isErr:   false,
	},
	TestData{
		json:    `{"picture": "http://hogehoge"}`,
		isGoson: true,
		isErr:   false,
	},
	TestData{
		json:    `{afsf: adfaasf`,
		isGoson: false,
		isErr:   true,
	},
	TestData{
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
	TestData{
		json:    `[{"name": "little"}, {"name": "tiny"}]`,
		isGoson: true,
		isErr:   false,
	},
	TestData{
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
