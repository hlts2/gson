package gson

import (
	"reflect"
	"strings"
	"testing"
)

type TestData struct {
	json    string
	isGoson bool
	path    string
	keys    []string
	result  *Result
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

func (n TestData) CheckGosonInstance(g *Goson) bool {
	if n.isGoson && g != nil {
		return true
	}

	if !n.isGoson && g == nil {
		return true
	}
	return false
}

func (n TestData) CheckResultObject(result *Result) bool {
	if result == nil && n.result == nil {
		return true
	}

	if reflect.DeepEqual(n.result.object, result.object) {
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

		if !test.CheckGosonInstance(g) {
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

		if !test.CheckGosonInstance(g) {
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

		if !test.CheckGosonInstance(g) {
			t.Error("")
		}
	}
}

var searchTests = TestDatas{
	TestData{
		json:   `{"name": "hlts2"}`,
		keys:   []string{"name"},
		path:   "/name",
		result: &Result{"hlts2"},
		isErr:  false,
	},
	TestData{
		json:   `[{"name": "hlts2"}]`,
		keys:   []string{"0"},
		path:   "/0",
		result: &Result{map[string]interface{}{"name": "hlts2"}},
		isErr:  false,
	},
	TestData{
		json:   `[{"name": "hlts2"}]`,
		keys:   []string{"10"},
		path:   "/10",
		result: &Result{nil},
		isErr:  true,
	},
	TestData{
		json: `
		{"friends": [
      		{
				"id": "0",
				"name": "hiro"
			},
      		{
				"id": "1",
				"name": "hlts2"
			}
		]}
		`,
		keys:   []string{"friends"},
		path:   "/friends",
		result: &Result{[]interface{}{map[string]interface{}{"id": "0", "name": "hiro"}, map[string]interface{}{"id": "1", "name": "hlts2"}}},
		isErr:  false,
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
		keys:   []string{"friends", "100", "name"},
		path:   "/friends/100/name",
		result: &Result{nil},
		isErr:  true,
	},
}

func TestSearch(t *testing.T) {
	for _, test := range searchTests {
		g, _ := NewGosonFromString(test.json)

		result, err := g.Search(test.keys...)

		if !test.CheckError(err) {
			t.Error("")
		}

		if !test.CheckResultObject(result) {
			t.Error("")
		}
	}
}

func TestPath(t *testing.T) {
	for _, test := range searchTests {
		g, _ := NewGosonFromString(test.json)

		result, err := g.Path(test.path)

		if !test.CheckError(err) {
			t.Error("")
		}

		if !test.CheckResultObject(result) {
			t.Error("")
		}
	}
}

var resultTests = TestDatas{
	TestData{},
}

func TestResult(t *testing.T) {
}
